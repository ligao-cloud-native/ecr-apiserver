package main

import (
	"flag"
	"github.com/ligao-cloud-native/ecr-apiserver/api"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db"
	"k8s.io/component-base/logs"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "f", "", "service config file.")
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	flag.Parse()
	config.LoadConfig(configFile)

	db.InitDB()

	api.StartHTTPServer()
}
