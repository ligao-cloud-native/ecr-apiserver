package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
)

type Webhook struct {
	gorm.Model
	Name       string `gorm:"size:255" json:"name"`
	Region     string `gorm:"size:255" json:"region"`
	Namespace  string `gorm:"size:255" json:"namespace"`
	Repository string `gorm:"size:255" json:"repository"`
}

func init() {
	db.Register("webhook", &User{})
}

func (hook Webhook) Get() (*Webhook, error) {
	hookRaw := &hook
	err := db.DB.Where("name = ?", hook.Name).First(hookRaw).Error
	return hookRaw, err
}

func (hook Webhook) List() ([]*Webhook, error) {
	var hooks []*Webhook
	err := db.DB.Find(&hooks).Error
	return hooks, err
}
