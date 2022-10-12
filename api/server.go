package api

import (
	"github.com/gorilla/mux"
	"github.com/ligao-cloud-native/ecr-apiserver/api/middleware"
	"github.com/ligao-cloud-native/ecr-apiserver/harborapi"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/metrics"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook"
	"k8s.io/klog/v2"
	"net/http"
	"strconv"
	"time"
)

// StartHTTPServer starts the http service
func StartHTTPServer() {
	// registry metrics
	m := metrics.NewMetrics()
	m.HandleMetrics()

	//init harbor clients
	harborapi.InitHarborClients()

	// enable harbor webhook
	if config.Conf.WebHook.Enable {
		webhook.InitWebhook().StartSubscriber()
	}

	router := mux.NewRouter().StrictSlash(true)
	RegisterRootRouters(router)

	v1 := router.PathPrefix("/api/v1").Subrouter()
	middleware.InitMiddlewares(v1).Use()
	RegisterV1Routers(v1)

	server := &http.Server{
		Addr:         config.Conf.API.Host + ":" + strconv.Itoa(config.Conf.API.Port),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	klog.Fatal(server.ListenAndServe())

}
