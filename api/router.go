package api

import (
	"github.com/gorilla/mux"
	"github.com/ligao-cloud-native/ecr-apiserver/api/handler"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func RegisterRootRouters(r *mux.Router) {
	r.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	r.Handle("/metrics", promhttp.Handler())

	// docker registry auth token
	r.HandleFunc("/service/token", handler.GetToken).Methods("GET")

	// harbor webhooks
	r.HandleFunc("/service/webhooks", handler.WebHookHandle).Methods("POST")
}

func RegisterV1Routers(v1 *mux.Router) {
	// instance: 费用管理的粒度，通过设置实例的项目数，仓库数的配额等来进行计费
	v1.HandleFunc("/instance", handler.CreateInstance).Methods(http.MethodPost)

	// ecr namespace/harbor project
	v1.HandleFunc("/namespace", handler.CreateNamespace).Methods(http.MethodPost)
	v1.HandleFunc("/namespace", handler.ListNamespace).Methods(http.MethodGet)

	// chart repo
	v1.HandleFunc("/chartrepo/{instance}/{namespace}", handler.CreateChart).Methods("POST")
}
