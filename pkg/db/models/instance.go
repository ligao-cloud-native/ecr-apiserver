package models

import "github.com/ligao-cloud-native/ecr-apiserver/pkg/db"

// ecr以实例为粒度进行计费管理
type Instance struct {
	Name string `gorm:"size:255" json:"name"`
	//允许创建的项目数
	NamespaceLimit int `gorm:"default:-1" json:"namespace_limit"`
	//允许创建仓库数，包括镜像仓库和chart仓库，-1表示不设限。（实例下所有项目的）
	RepoCountLimit int `gorm:"default:-1" json:"repo_count_limit"`
	// 分配的最大存储空间，-1表示不设限。（实例下所有项目的）
	StorageLimit int `gorm:"default:-1" json:"storage_limit"`
	//购买时长，单位：月
	BuyDuration int `gorm:"default:-1" json:"buy_duration"`
	//使用的数量
	NamespaceUsed int `gorm:"default:0" json:"namespace_used"`
	RepoCountUsed int `gorm:"default:0" json:"repo_count_used"`
	StorageUsed   int `gorm:"default:0" json:"storage_used"`
	DurationUsed  int `gorm:"default:1" json:"duration_used"`
	// 仓库实例的资源是否全部用完,
	IsEmpty bool `gorm:"default:true" json:"is_empty"`
}

func init() {
	db.Register("instance", &Instance{})
}

// return true: is empty; false: not empty
func (ins Instance) CheckQuota() bool {
	if db.DB.Where("name = ? AND is_empty = 1", ins.Name).First(&Instance{}).RecordNotFound() {
		return false
	}

	return true
}

func (ins Instance) Create() error {
	return db.DB.Create(&ins).Error
}
