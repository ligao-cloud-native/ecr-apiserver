package token

import (
	"fmt"
	regtoken "github.com/docker/distribution/registry/auth/token"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/auth/user"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"k8s.io/klog/v2"
	"net/http"
	"regexp"
	"strings"
)

type FilterType string

const (
	RegistryFilter   FilterType = "registry"
	RepositoryFilter FilterType = "repository"
)

type accessFilter interface {
	filter(r *http.Request, a *regtoken.ResourceActions) error
}

type registryFilter struct{}

func (reg *registryFilter) filter(r *http.Request, a *regtoken.ResourceActions) error {
	return nil
}

type repositoryFilter struct {
	parser parser
}

// docker pull/push auth
func (repo *repositoryFilter) filter(r *http.Request, a *regtoken.ResourceActions) error {
	klog.Info("exec repository filter")
	img, err := repo.parser.parse(a.Name)
	if err != nil {
		klog.Errorf("filter parse image error: %s", err.Error())
		return err
	}

	ns, err := models.Namespace{Name: img.namespace}.Get()
	if err != nil || ns.Name == "" {
		klog.Warningf("namespace %s not exist, no permission", img.namespace)
		return AccessDeny(a)
	}

	if !checkRepoFomart(img.name) {
		klog.Warningf("repository name %s invalid format, no permission", img.name)
		return AccessDeny(a)
	}

	region := r.Header.Get("region")
	username := r.Context().Value("username").(string)
	password := r.Context().Value("password").(string)
	repoInfo, err := models.Repository{Name: img.name, Namespace: img.namespace, Region: region}.Get()
	if err != nil {
		klog.Warningf("repository %s not exist , no permission", img.name)
		return AccessDeny(a)
	}

	//没有login的用户，只允许pull免密的仓库
	if username == "" {
		if repoInfo.Name != "" && repoInfo.AvoidPassword != nil && *repoInfo.AvoidPassword {
			klog.Info("avoidPassword allow read!")
			return AccessOnlyPull(a)
		}
		klog.Warningf("no login, and repo %s is private", img.name)
		return AccessDeny(a)
	}

	// 非镜像商城处理逻辑
	ok, userInfo := user.VerifyUser(username, password)
	if userInfo != nil {
		if ok {
			if userInfo.SystemAdmin == 1 {
				return AccessOnlyPull(a)
			}
			if userInfo.SystemAdmin == 2 {
				return AccessPullPush(a)
			}

			//私有仓库，创建人可以pull/push
			if repoInfo.Name != "" && !repoInfo.Public {
				if strings.EqualFold(repoInfo.Username, username) {
					return AccessPullPush(a)
				}
				//免密可pull
				if repoInfo.AvoidPassword != nil && *repoInfo.AvoidPassword {
					return AccessOnlyPull(a)
				}

				return AccessDeny(a)
			}

		} else {
			//启用云平台验证系统
			verified := user.VerifyCloudAccount(username, password)
			if !verified {
				if repoInfo.AvoidPassword != nil && *repoInfo.AvoidPassword {
					return AccessOnlyPull(a)
				}
				return fmt.Errorf("unauthorized")
			}

		}

	} else {
		// 用户不存在，验证service account账户(即门户创建的访问凭证账号)
		verified := user.VerifyServiceAccount(username, password)
		if verified {
			if repoInfo.Name != "" && !repoInfo.Public {
				if repoInfo.AvoidPassword != nil && *repoInfo.AvoidPassword {
					return AccessOnlyPull(a)
				}

				return AccessDeny(a)
			}
		} else {
			if repoInfo.AvoidPassword != nil && *repoInfo.AvoidPassword {
				return AccessOnlyPull(a)
			}
			return fmt.Errorf("unauthorized")
		}

	}

	// 镜像商城处理逻辑
	if region == "master" {
	}

	return nil

}

func checkRepoFomart(name string) bool {
	pat := "^([a-z0-9]+(?:[._-][a-z0-9]+)*){2,128}$"
	reg := regexp.MustCompile(pat)
	return reg.MatchString(name)
}
