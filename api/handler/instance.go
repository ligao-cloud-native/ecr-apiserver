package handler

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"github.com/ligao-cloud-native/ecr-apiserver/utils/httputils"
	"k8s.io/klog/v2"
	"net/http"
)

func CreateInstance(w http.ResponseWriter, r *http.Request) {
	var ins models.Instance
	if err := jsoniter.NewDecoder(r.Body).Decode(&ins); err != nil {
		httputils.BadRequest(w, r, err.Error())
		return
	}

	if err := ins.Create(); err != nil {
		klog.Errorf("create instance %s error: %s", "", err.Error())
		httputils.InternalServerError(w, r, err.Error())
		return
	}

	httputils.OK(w, r, "OK")

}
