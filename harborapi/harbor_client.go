package harborapi

import (
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/rest"
	"k8s.io/klog/v2"
	"sync"
)

var once sync.Once
var HarborClients harborClient

type harborClient struct {
	clients map[string]*ClientSet
}

func InitHarborClients() {
	once.Do(func() {
		harbors, err := models.Harbor{}.List()
		if err != nil {
			klog.Error("get harbor server error: %s", err.Error())
			return
		}

		HarborClients = harborClient{clients: make(map[string]*ClientSet)}
		for _, harbor := range harbors {
			klog.Infof("init harbor client: [%s] %s", harbor.Region, harbor.ApiServer)
			restConfig := rest.Config{
				Host:     harbor.ApiServer,
				APIPath:  "/api/v2.0",
				Username: harbor.Username,
				Password: harbor.Password,
			}
			c, err := NewClient(&restConfig)
			if err != nil {
				klog.Error("init harbor client error: %s", err.Error())
				return
			}
			HarborClients.clients[harbor.Region] = c
		}

	})

}

func (c *harborClient) List() map[string]*ClientSet {
	return c.clients
}

func (c *harborClient) Get(region string) *ClientSet {
	return c.clients[region]
}
