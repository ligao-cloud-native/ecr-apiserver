package middleware

import (
	"github.com/gorilla/mux"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/auth"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/log"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware/user"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
	"net/http"
)

type Middleware interface {
	Middleware(next http.Handler) http.Handler
}

type Middlewares struct {
	r           *mux.Router
	middlewares []Middleware
}

func InitMiddlewares(r *mux.Router) *Middlewares {
	mws := Middlewares{r: r}
	//log
	mws.middlewares = append(mws.middlewares, log.NewLogMiddleware())

	// user mode
	umw := user.NewUserMiddleware(config.Conf.UserMode)
	umw.User.Init()
	mws.middlewares = append(mws.middlewares, umw)

	// enable auth
	if config.Conf.AuthEnable {
		mws.middlewares = append(mws.middlewares, auth.NewAuthMiddleware())
	}

	// access: get/create/update/delete operator

	return &mws
}

func (mws *Middlewares) Use() {
	for _, mw := range mws.middlewares {
		mws.r.Use(mw.Middleware)
	}
}
