package httputils

import "net/http"

var codeStatusMap = map[int]string{
	http.StatusContinue:           "Continue",
	http.StatusSwitchingProtocols: "Switching Protocols",
	http.StatusProcessing:         "Processing",
	http.StatusEarlyHints:         "Early Hints",

	http.StatusOK:                   "请求成功",
	http.StatusCreated:              "Created",
	http.StatusAccepted:             "Accepted",
	http.StatusNonAuthoritativeInfo: "Non-Authoritative Information",
	http.StatusNoContent:            "No Content",
	http.StatusResetContent:         "Reset Content",
	http.StatusPartialContent:       "Partial Content",
	http.StatusMultiStatus:          "Multi-Status",
	http.StatusAlreadyReported:      "Already Reported",
	http.StatusIMUsed:               "IM Used",

	http.StatusMultipleChoices:   "Multiple Choices",
	http.StatusMovedPermanently:  "Moved Permanently",
	http.StatusFound:             "Found",
	http.StatusSeeOther:          "See Other",
	http.StatusNotModified:       "Not Modified",
	http.StatusUseProxy:          "Use Proxy",
	http.StatusTemporaryRedirect: "Temporary Redirect",
	http.StatusPermanentRedirect: "Permanent Redirect",

	http.StatusBadRequest:                   "参数错误",
	http.StatusUnauthorized:                 "认证失败，请重新登录",
	http.StatusPaymentRequired:              "Payment Required",
	http.StatusForbidden:                    "权限不足，拒绝访问",
	http.StatusNotFound:                     "资源不存在",
	http.StatusMethodNotAllowed:             "Method Not Allowed",
	http.StatusNotAcceptable:                "Not Acceptable",
	http.StatusProxyAuthRequired:            "Proxy Authentication Required",
	http.StatusRequestTimeout:               "Request Timeout",
	http.StatusConflict:                     "Conflict",
	http.StatusGone:                         "Gone",
	http.StatusLengthRequired:               "Length Required",
	http.StatusPreconditionFailed:           "Precondition Failed",
	http.StatusRequestEntityTooLarge:        "Request Entity Too Large",
	http.StatusRequestURITooLong:            "Request URI Too Long",
	http.StatusUnsupportedMediaType:         "Unsupported Media Type",
	http.StatusRequestedRangeNotSatisfiable: "Requested Range Not Satisfiable",
	http.StatusExpectationFailed:            "Expectation Failed",
	http.StatusTeapot:                       "I'm a teapot",
	http.StatusMisdirectedRequest:           "Misdirected Request",
	http.StatusUnprocessableEntity:          "Unprocessable Entity",
	http.StatusLocked:                       "Locked",
	http.StatusFailedDependency:             "Failed Dependency",
	http.StatusTooEarly:                     "Too Early",
	http.StatusUpgradeRequired:              "Upgrade Required",
	http.StatusPreconditionRequired:         "Precondition Required",
	http.StatusTooManyRequests:              "Too Many Requests",
	http.StatusRequestHeaderFieldsTooLarge:  "Request Header Fields Too Large",
	http.StatusUnavailableForLegalReasons:   "Unavailable For Legal Reasons",

	http.StatusInternalServerError:           "服务内部处理错误",
	http.StatusNotImplemented:                "Not Implemented",
	http.StatusBadGateway:                    "Bad Gateway",
	http.StatusServiceUnavailable:            "Service Unavailable",
	http.StatusGatewayTimeout:                "Gateway Timeout",
	http.StatusHTTPVersionNotSupported:       "HTTP Version Not Supported",
	http.StatusVariantAlsoNegotiates:         "Variant Also Negotiates",
	http.StatusInsufficientStorage:           "Insufficient Storage",
	http.StatusLoopDetected:                  "Loop Detected",
	http.StatusNotExtended:                   "Not Extended",
	http.StatusNetworkAuthenticationRequired: "Network Authentication Required",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func statusText(code int) string {
	return codeStatusMap[code]
}

func OK(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusOK, statusText(http.StatusOK), msg)
}

func BadRequest(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusBadRequest, statusText(http.StatusBadRequest), msg)
}

func ResourceNotFound(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusNotFound, statusText(http.StatusNotFound), msg)
}

func Forbidden(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusForbidden, statusText(http.StatusForbidden), msg)
}

func Unauthorized(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusUnauthorized, statusText(http.StatusUnauthorized), msg)
}

func InternalServerError(w http.ResponseWriter, r *http.Request, msg interface{}) {
	Response(w, r, http.StatusInternalServerError, statusText(http.StatusInternalServerError), msg)
}
