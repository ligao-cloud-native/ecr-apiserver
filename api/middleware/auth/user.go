package auth

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

type userProvider struct {
	name string
}

func (p *userProvider) Name() string {
	return p.name
}

func (p *userProvider) Auth(r *http.Request) error {
	user := r.Context().Value("user")
	if user == nil {
		return fmt.Errorf("get user error from request context")
	}

	var ua UserAccount
	if err := jsoniter.Unmarshal(user.([]byte), &ua); err != nil {
		return fmt.Errorf("user info invailied: %v", user)
	}

	return nil
}

type UserAccount struct {
	LoginName string `json:"login_name"`
}
