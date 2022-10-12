package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
)

type User struct {
	gorm.Model
	Password string `gorm:"size:255" json:"password"`
	Username string `gorm:"size:255" json:"username"`
	//1: only pull; 2: push and pull; 0: common
	SystemAdmin     int    `gorm:"size:255" json:"system_admin"`
	RoleName        string `gorm:"size:255" json:"role_name"`
	ResourceGroupID string `gorm:"size:255" json:"resource_group_id"`
}

func init() {
	db.Register("user", &User{})
}

func (u User) Get() (*User, error) {
	uRaw := &u
	err := db.DB.Where("username = ?", u.Username).First(uRaw).Error
	return uRaw, err
}
