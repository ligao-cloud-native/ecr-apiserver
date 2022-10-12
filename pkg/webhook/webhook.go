package webhook

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/db/models"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook/bus"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook/notification"
	"k8s.io/klog/v2"
	"sync"
)

var Webhook *webhook
var once sync.Once

type webhook struct {
	Payload *notification.Payload

	eventType map[notification.EventType]struct{}
	bus       bus.NotificationBus
}

func InitWebhook() *webhook {
	once.Do(func() {
		hook := webhook{eventType: notification.GetEventType()}
		hook.bus.Bus = EventBus.New()
		hook.bus.Callback = hook.dealHookNotification

		Webhook = &hook
	})

	return Webhook
}

func (hook *webhook) Deliver(region string) {
	for _, event := range hook.Payload.Events {
		topic := fmt.Sprintf("%s:%s:%s", region, event.Project, event.RepoName)
		klog.Infof("topic: %s", topic)
		hook.bus.DeliverToBus(topic, hook.Payload)
	}

}

func (hook *webhook) StartSubscriber() {
	hook.subscribeHookNotification()
	go hook.startSubscriber()
}

// subscribeHookNotification subscribe hook notification from bus
func (hook *webhook) subscribeHookNotification() {
	//hooks, err := models.Webhook{}.List()
	//if err != nil {
	//	klog.Errorf("get all webhooks error: %s", err.Error())
	//	return
	//}
	// test
	hooks := []*models.Webhook{
		{Region: "sz", Namespace: "kmc", Repository: "pwc-apiserver"},
	}

	for _, wk := range hooks {
		topic := fmt.Sprintf("%s:%s:%s", wk.Region, wk.Namespace, wk.Repository)
		if err := hook.bus.BusSubscribe(topic); err != nil {
			klog.Warningf("Webhook [region:%s, namespace:%s, repo:%s] subscribe failed: %s",
				wk.Region, wk.Namespace, wk.Repository, err.Error())
		}
	}
}

// dealHookNotification a callback func to handle webhook notification
func (hook *webhook) dealHookNotification(payload *notification.Payload) {
	klog.Infof("payload: %v", *payload)

}

func (hook *webhook) startSubscriber() {
	bus.StartRedisSubscriber()
}
