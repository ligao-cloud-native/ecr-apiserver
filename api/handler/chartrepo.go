package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"net/http"
)

func CreateChart(w http.ResponseWriter, r *http.Request) {
	instance := mux.Vars(r)["instance"]
	namespace := mux.Vars(r)["namespace"]

	//检查ns
	if !(models.Namespace{Name: namespace}.IsExist()) {
		httputils.BadRequest(w, r, fmt.Sprintf("namespace %s not exist", namespace))
		return
	}
	// 检查配额
	if !(models.Instance{Name: instance}.CheckQuota()) {
		httputils.BadRequest(w, r, fmt.Sprintf("instance %s not enough quota", instance))
		return
	}

	httputils.OK(w, r, "OK")

}
