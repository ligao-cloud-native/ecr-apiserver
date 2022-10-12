package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
	"k8s.io/klog/v2"
)

//Harbor all harbor server
type Harbor struct {
	gorm.Model
	Region    string `gorm:"size:255" json:"region"`
	ApiServer string `gorm:"size:255" json:"api_server"`
	Username  string `gorm:"size:255" json:"username"`
	Password  string `gorm:"size:255" json:"password"`
}

func init() {
	db.Register("harbor", &Harbor{})
}

func (hb Harbor) Get(q string) (*Harbor, error) {
	hbRaw := &Harbor{}
	err := db.DB.Where("region = ?", q).Or("api_server = ?", q).First(hbRaw).Error
	if err != nil {
		klog.Errorf("Model:Harbor:First(%s) error, %s", q, err.Error())
	}
	return hbRaw, err
}

func (hb Harbor) List() (harborList []Harbor, err error) {
	err = db.DB.Find(&harborList).Error
	return
}
