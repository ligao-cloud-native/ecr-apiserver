package rest

import "time"

type Config struct {
	Host    string
	APIPath string

	// server requires Basic authentication
	Username string
	Password string

	// specify the caller of this request
	UserAgent string

	// time to wait before giving up on a server request.  zero value means no timeout
	Timeout time.Duration
}
