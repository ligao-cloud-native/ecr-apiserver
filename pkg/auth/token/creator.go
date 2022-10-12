package token

import (
	"fmt"
	regtoken "github.com/docker/distribution/registry/auth/token"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/auth/user"
	"github.com/ligao-cloud-native/ecr-apiserver/utils"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
)

// baseCreator a base token creator
type baseCreator struct {
	filters map[FilterType]accessFilter
}

func (c *baseCreator) CreateToken(r *http.Request) (*ECRToken, error) {
	// action access
	actions := getResourceActions(r.URL.Query()["scope"])
	if err := c.accessFilter(r, actions); err != nil {
		return nil, err
	}

	//make token
	return MakeToken(actions)
}

func (c *baseCreator) accessFilter(r *http.Request, actions []*regtoken.ResourceActions) error {
	if len(actions) == 0 {
		username := r.Context().Value("username").(string)
		password := r.Context().Value("password").(string)
		// 验证平台登录用户
		ok, u := user.VerifyUser(username, password)
		if u != nil { // 用户存在
			if !ok { // 密码不正确
				//启用云平台验证系统
				verified := user.VerifyCloudAccount(username, password)
				if !verified {
					klog.Errorf("cloud user %s verified failed", username)
					return fmt.Errorf("Unauthorized")
				}
			}
		} else {
			// 用户不存在，验证service account账户(即门户创建的访问凭证账号)
			verified := user.VerifyServiceAccount(username, password)
			if !verified {
				klog.Errorf("harbor user %s verified failed", username)
				return fmt.Errorf("Unauthorized")
			}
		}

		return nil
	}

	for _, a := range actions {
		klog.Infof("%v", *a)
		f, ok := c.filters[FilterType(a.Type)]
		if !ok {
			continue
		}
		// 下面filter修改了actions
		sourceActions := a.Actions

		if err := f.filter(r, a); err != nil {
			klog.Error(err.Error())
			return err
		}

		// push镜像， 需要校验配额
		if a.Type == string(RepositoryFilter) && utils.StrInArray("push", sourceActions) {

		}
	}

	return nil
}

// registryCreator a registry service token creator
type registryCreator struct {
	baseCreator
}

func newRegistryCreator() *registryCreator {
	filters := map[FilterType]accessFilter{
		RepositoryFilter: &repositoryFilter{
			parser: &baseParser{},
		},
		RegistryFilter: &registryFilter{},
	}

	return &registryCreator{
		baseCreator: baseCreator{filters: filters},
	}
}

// notaryCreator a notary service token creator
type notaryCreator struct {
	baseCreator
}

func newNotaryCreator() *notaryCreator {
	filters := map[FilterType]accessFilter{
		RepositoryFilter: &repositoryFilter{
			parser: &endpointParser{},
		},
	}

	return &notaryCreator{
		baseCreator: baseCreator{filters: filters},
	}
}

func getResourceActions(scopeParam []string) (actions []*regtoken.ResourceActions) {
	//?scope=#{resource.Type}:#{resource.Name}:#{action1},#{action1} 2&&scope=a b
	// repository:library/nginx:v1:push
	var scopes []string
	for _, sp := range scopeParam {
		scopes = append(scopes, strings.Split(sp, " ")...)
	}
	klog.Infof("actions: %v", scopes)

	for _, scope := range scopes {
		if scope == "" {
			continue
		}

		resource := strings.Split(scope, ":")
		if len(resource) == 0 {
			continue
		}

		resourceActions := regtoken.ResourceActions{}
		switch len(resource) {
		case 1:
			resourceActions.Type = resource[0]
		case 2:
			resourceActions.Type = resource[0]
			resourceActions.Name = resource[1]
		default:
			resourceActions.Type = resource[0]
			resourceActions.Name = strings.Join(resource[1:len(resource)-1], ":")
			if len(resource[len(resource)-1]) > 0 {
				resourceActions.Actions = strings.Split(resource[len(resource)-1], ",")
			}
		}
		actions = append(actions, &resourceActions)
	}

	return
}
