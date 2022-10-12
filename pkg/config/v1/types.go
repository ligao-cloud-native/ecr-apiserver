package v1

import "github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"

type Configure struct {
	API         api
	AuthEnable  bool
	UserMode    string
	UserAccount UserAccount
	WebHook     webHook
	Redis       redisConf
}

type api struct {
	Host string
	Port int
}

type UserAccount struct {
	User   user
	Roles  role
	Groups []models.ResourceGroup
}

type user struct {
	Name   string
	CnName string
}

type role struct {
	RoleID   string
	RoleName string
}

type webHook struct {
	Enable bool
	Url    string
}

type redisConf struct {
	Url      string
	Password string
	DbNum    int
}
