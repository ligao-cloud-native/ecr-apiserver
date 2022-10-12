package config

import (
	v1 "github.com/ligao-cloud-native/ecr-apiserver/pkg/config/v1"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
	"sync"
)

var (
	Conf Config
	once sync.Once
)

type Config struct {
	v1.Configure
}

func LoadConfig(configFile string) {
	once.Do(func() {
		viper.SetConfigName("config")           // 配置文件名称(无扩展名)
		viper.AddConfigPath("/etc/api-server/") // 查找配置文件所在的路径
		viper.AddConfigPath(".")                // 还可以在工作目录中查找配置
		if len(configFile) > 0 {
			viper.SetConfigFile(configFile) // 指定配置文件路径
		} else {
			klog.Infof("use default config file /etc/api-server/config.yaml, ./config.yaml")
		}

		if err := viper.ReadInConfig(); err != nil {
			klog.Fatalf("load service config file failed: %s", err.Error())
		}
		klog.Infof("service config: %s", viper.ConfigFileUsed())

		v := viper.GetViper()
		if err := v.Unmarshal(&Conf.Configure); err != nil {
			klog.Fatalf("parse service config failed: %s", err.Error())
		}

		klog.Infof("config: %v", Conf)

	})

}
