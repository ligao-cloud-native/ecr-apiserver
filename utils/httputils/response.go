package httputils

import (
	jsoniter "github.com/json-iterator/go"
	"k8s.io/klog/v2"
	"net/http"
)

type response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status,omitempty"`
	Message interface{} `json:"message"`
}

func Response(w http.ResponseWriter, r *http.Request, httpCode int, httpStatus string, message interface{}) {
	res := response{
		Code:    httpCode,
		Status:  httpStatus,
		Message: message,
	}

	jsonByte, err := jsoniter.Marshal(res)
	if err != nil {
		klog.Errorf("Marshal [%v] error: %v", res, err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	_, err = r.Cookie("WriteHeader")
	// no cookie WriteHeader
	if err != nil {
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpCode)
		_, _ = w.Write(jsonByte)
	}
}
