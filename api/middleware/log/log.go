package log

import (
	"k8s.io/klog/v2"
	"net/http"
	"time"
)

type LogMiddleware struct {
}

func NewLogMiddleware() *LogMiddleware {
	return &LogMiddleware{}
}

func (m *LogMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		end := time.Now()

		latency := end.Sub(start)
		clientIP := r.RemoteAddr
		urlPath := r.URL.RequestURI()
		method := r.Method

		klog.Infof("|%v |%s |%s %s", latency, clientIP, method, urlPath)
	})
}
