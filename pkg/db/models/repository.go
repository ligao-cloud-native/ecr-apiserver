package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
)

//Repository is a image name
type Repository struct {
	gorm.Model
	Namespace     string `gorm:"size:255" json:"namespace"`
	Name          string `gorm:"size:255" json:"name"`
	Username      string `gorm:"size:255" json:"username"`
	Region        string `gorm:"size:255" json:"region"`
	Description   string `gorm:"size:255" json:"description"`
	Public        bool   `gorm:"default:true" json:"public"`
	AvoidPassword *bool  `gorm:"-" json:"avoid_password"`
}

func init() {
	db.Register("repository", &Repository{})
}

func (repo Repository) Get() (*Repository, error) {
	repoRaw := &repo
	err := db.DB.Where("name = ? AND namespace = ? AND region = ?",
		repo.Name, repo.Namespace, repo.Region).First(repoRaw).Error
	return repoRaw, err
}
