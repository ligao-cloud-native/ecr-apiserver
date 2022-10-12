package models

import (
	"github.com/jinzhu/gorm"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
)

//Namespace is harbor project
type Namespace struct {
	gorm.Model
	//项目名称
	Name string `gorm:"size:255" json:"name"`
	//项目中支持的镜像仓库和chart仓库的总数量，-1表示不设限
	CountLimit int `gorm:"default:-1" json:"count_limit"`
	// 分配给项目的最大存储空间，-1表示不设限
	StorageLimit int `gorm:"default:-1" json:"storage_limit"`
	// 项目的访问级别，公开或私有
	Public bool `gorm:"default:true" json:"public"`
	//创建项目的用户
	Username string `gorm:"size:255" json:"username"`
	//创建项目的用户的中文名
	CnUsername string `gorm:"size:255" json:"cn_username"`
	// 用户所在的资源组
	ResourceGroupUuid string        `gorm:"size:255" json:"-"`
	ResourceGroup     ResourceGroup `gorm:"foreignkey:ResourceGroupUuid;association_foreignkey;ResourceGroupUuid"`
}

func init() {
	db.Register("namespace", &Namespace{})
}

func (ns Namespace) Get() (*Namespace, error) {
	nsRaw := &ns
	err := db.DB.Where("name = ?", ns.Name).First(nsRaw).Error
	return nsRaw, err
}

func (ns Namespace) Create() error {
	return db.DB.Create(&ns).Error
}

func (ns Namespace) Update() {}

func (ns Namespace) Delete() {}

func (ns Namespace) IsExist() bool {
	if db.DB.Where("name = ?", ns.Name).First(&Namespace{}).RecordNotFound() {
		return false
	}

	return true
}
