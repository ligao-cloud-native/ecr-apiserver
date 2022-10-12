package user

import (
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
)

func VerifyUser(username, pwd string) (bool, *models.User) {
	user, err := models.User{Username: username}.Get()
	if err != nil {
		return false, nil
	}

	if user.Password != pwd {
		return false, user
	}

	return true, user
}
