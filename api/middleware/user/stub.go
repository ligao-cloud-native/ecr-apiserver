package user

import (
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/user/account"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/user/cache"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
	v1 "github.com/ligao-cloud-native/ecr-apiserver/pkg/config/v1"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"net/http"
)

// 从配置文件中获取用户认证信息
type StubUser struct {
	UserAccount v1.UserAccount
}

func (stub *StubUser) Init() {
	stub.UserAccount = config.Conf.UserAccount

	// 将用户信息写入redis缓存
}

func (stub *StubUser) GetUserAccountInfo(r *http.Request) *account.UserAccountInfo {
	name := stub.getUser(r)
	// 从redis缓存中获取用户账号信息
	_ = cache.GetUserAccount(name)

	//test
	return &account.UserAccountInfo{
		LoginName: stub.UserAccount.User.Name,
		CnName:    stub.UserAccount.User.CnName,
		RoleId:    stub.UserAccount.Roles.RoleID,
		RoleName:  stub.UserAccount.Roles.RoleName,
	}

}

func (stub *StubUser) GetResourceGroup(r *http.Request) []models.ResourceGroup {
	return stub.UserAccount.Groups
}

func (stub *StubUser) GetRoleGroup(r *http.Request) error {
	return nil
}

func (stub *StubUser) getUser(r *http.Request) string {
	user := r.Header.Get("user")
	if user != "" {
		return user
	}

	if r.Method == http.MethodGet {
		return r.URL.Query().Get("user")
	}

	return ""
}
