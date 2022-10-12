package auth

import "net/http"

type openapiProvider struct {
	name string
}

func (p *openapiProvider) Name() string {
	return p.name
}

func (p *openapiProvider) Auth(r *http.Request) error {
	return nil
}
