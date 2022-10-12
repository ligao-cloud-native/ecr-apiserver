package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"k8s.io/klog/v2"
)

var DB *gorm.DB

func InitDB() {
	var err error
	connUrl := "root:lenovo@tcp(192.168.25.3:3306)/ecr?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", connUrl)
	if err != nil {
		klog.Fatalf("init db connection failed: %s", err.Error())
		return
	}

	// 在gorm中，默认的表名都是结构体名称的复数形式
	// 取消表名的复数形式，使得表名和结构体名称一致
	DB.SingularTable(true)
	DB.LogMode(true)

	sqlDB := DB.DB()
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxIdleTime(60)
	sqlDB.SetConnMaxLifetime(7200)

	if err := sqlDB.Ping(); err != nil {
		klog.Fatalf("db connect failed: %s", err.Error())
	}

	for tableName, model := range GetModels() {
		klog.Infof("migrate db table %s", tableName)
		DB.AutoMigrate(model)
	}

}

func Close() error {
	return DB.DB().Close()
}
