package handler

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook/notification"
	"k8s.io/klog/v2"
	"net/http"
)

func WebHookHandle(w http.ResponseWriter, r *http.Request) {
	var payload notification.Payload
	if err := jsoniter.NewDecoder(r.Body).Decode(&payload); err != nil {
		klog.Errorf("webhook notification error: %s", err.Error())
		return
	}

	region := r.FormValue("region")
	if region == "" {
		klog.Error("no region from webhook url")
		return
	}

	webhook.Webhook.Payload = &payload
	go webhook.Webhook.Deliver(region)
}
