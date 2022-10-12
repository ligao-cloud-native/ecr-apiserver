package bus

import (
	"github.com/asaskevich/EventBus"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/webhook/notification"
)

type NotificationBus struct {
	Bus      EventBus.Bus
	Callback func(payload *notification.Payload)
}

func (bus *NotificationBus) DeliverToBus(topic string, payload *notification.Payload) {
	bus.Bus.Publish(topic, payload)
	bus.Bus.WaitAsync()
}

func (bus *NotificationBus) BusSubscribe(topic string) error {
	return bus.Bus.SubscribeAsync(topic, bus.Callback, false)
}
