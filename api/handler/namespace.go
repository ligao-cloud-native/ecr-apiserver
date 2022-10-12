package handler

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/ligao-cloud-native/ecr-apiserver/harborapi"
	"github.com/ligao-cloud-native/ecr-apiserver/harborapi/types/meta"
	typesv2 "github.com/ligao-cloud-native/ecr-apiserver/harborapi/types/v2"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"k8s.io/klog/v2"
	"net/http"
)

func ListNamespace(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	region := query.Get("region")
	if region == "" {
		httputils.BadRequest(w, r, "request query region is nil")
		return
	}

	// 检查用户当前使用的资源组，是否在该用户的资源组列表中
	if query.Get("ResourceGroupUuid") != "" {

	} else {

	}
	klog.Infof("%s", string(r.Context().Value("user").([]byte)))

	opts := meta.ListOptions{}
	if query.Get("name") != "" {
		opts.Name = query.Get("name")
	}
	c := harborapi.HarborClients.Get(region)
	if c == nil {
		httputils.BadRequest(w, r, fmt.Sprintf("not supported region: %s", region))
		return
	}
	ps, err := c.V2().Projects().List(opts)
	if err != nil {
		httputils.InternalServerError(w, r, err.Error())
		return
	}
	httputils.OK(w, r, ps)
}

func CreateNamespace(w http.ResponseWriter, r *http.Request) {
	var ns models.Namespace
	if err := jsoniter.NewDecoder(r.Body).Decode(&ns); err != nil {
		httputils.BadRequest(w, r, err.Error())
		return
	}

	// 判断用户权限

	if ns.IsExist() {
		httputils.BadRequest(w, r, "namespace exists")
		return
	}

	// 在每个harbor上创建project，必须都成功。
	successed := make(map[string]*harborapi.ClientSet)
	project := typesv2.Project{ProjectName: ns.Name}
	for region, c := range harborapi.HarborClients.List() {
		err := c.V2().Projects().Create(&project)
		if err == nil {
			// 是否开启webhook
			if config.Conf.WebHook.Enable {
				// todo create project webhook
			}
			successed[region] = c
		} else {
			//删除已经创建的harbor project
			for _, cli := range successed {
				cli.V2().Projects().Delete()
			}
			klog.Errorf("Region [%s] harbor project [%s] create failed : %s", region, ns.Name, err.Error())
			httputils.InternalServerError(w, r, err.Error())
			return
		}

	}

	// 不同实例下的ns可以同名，但是harbor中的project不能同名。因此名称需要加上随机的后缀。
	// 在调用harbor接口创建项目时，配额设置为不限制，即-1。
	if err := ns.Create(); err != nil {
		klog.Errorf("create namespace %s error: %s", "", err.Error())
		httputils.InternalServerError(w, r, err.Error())
		return
	}

	httputils.OK(w, r, "OK")
}
