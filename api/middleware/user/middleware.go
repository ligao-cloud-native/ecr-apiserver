package user

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/user/account"
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"net/http"
)

type Interface interface {
	Init()
	GetUserAccountInfo(r *http.Request) *account.UserAccountInfo
	GetRoleGroup(r *http.Request) error
}

type UserMiddleware struct {
	User Interface
}

func NewUserMiddleware(userMode string) *UserMiddleware {
	um := UserMiddleware{}
	switch {
	case userMode == "" || userMode == "stub":
		um.User = &StubUser{}
	case userMode == "cloud":
		um.User = &CloudApi{}
	}

	return &um
}

func (a *UserMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user := a.User.GetUserAccountInfo(r); user != nil {
			userJson, _ := jsoniter.Marshal(user)

			r = r.WithContext(context.WithValue(r.Context(), "user", userJson))
		} else {
			httputils.Unauthorized(w, r, "authorized failed")
			return
		}

		next.ServeHTTP(w, r)
	})

}
