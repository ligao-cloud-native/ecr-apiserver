package auth

import (
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"net/http"
)

type provider interface {
	Name() string
	Auth(r *http.Request) error
}

type AuthMiddleware struct {
	provider []provider
}

func NewAuthMiddleware() *AuthMiddleware {
	providers := []provider{
		&userProvider{"user"},
		&openapiProvider{"openapi"},
	}

	return &AuthMiddleware{
		provider: providers,
	}
}

func (a *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// openapi和user一个认证成功即可
		for _, p := range a.provider {
			if err := p.Auth(r); err == nil {
				//klog.Infof("%s auth success.", p.Name())
				next.ServeHTTP(w, r)
				return
			} else {
				//klog.Warningf("%s auth failed: %s", p.Name(), err.Error())
			}
		}

		httputils.Unauthorized(w, r, "unauthorized")
	})
}
