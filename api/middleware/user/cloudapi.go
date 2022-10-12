package user

import (
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/user/account"
	"net/http"
)

//调用云平台认证接口进行认证
type CloudApi struct {
}

func (cloud *CloudApi) Init() {

}

func (cloud *CloudApi) GetUserAccountInfo(r *http.Request) *account.UserAccountInfo {
	return nil
}

func (cloud *CloudApi) GetRoleGroup(r *http.Request) error {
	return nil
}
