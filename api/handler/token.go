package handler

import (
	"context"
	"fmt"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/auth/token"
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"net/http"
	"strings"
)

// https://goharbor.io/docs/2.3.0/install-config/customize-token-service/
// Get /service/token?service=registry&scope=repository:library/nginx:v1:push
// GET /service/token?service=registry.docker.io&scope=repository:samalba/my-app:pull,push
func GetToken(w http.ResponseWriter, r *http.Request) {
	service := r.FormValue("service")
	region := r.FormValue("region")
	// -H "Authorization: Basic QWxhZGRpbjpvcGVuIHNlc2FtZQ=="
	username, pwd, _ := r.BasicAuth()

	username = strings.ToLower(username)
	r = r.WithContext(context.WithValue(r.Context(), "username", username))
	r = r.WithContext(context.WithValue(r.Context(), "password", pwd))
	r.Header.Set("region", region)

	creatorFactory := token.NewCreatorFactory(service)
	creatorObj := creatorFactory.Create()
	if creatorObj == nil {
		httputils.BadRequest(w, r, fmt.Sprintf("cannot handle %s service token", service))
		return
	}

	tk, err := creatorObj.CreateToken(r)
	if err != nil {
		httputils.InternalServerError(w, r, err.Error())
		return
	}

	httputils.OK(w, r, tk)
}
