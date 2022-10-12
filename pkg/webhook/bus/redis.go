package bus

import (
	goredis "github.com/gomodule/redigo/redis"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/redis"
	"k8s.io/klog/v2"
)

//redis 作为消息队列使用

func StartRedisSubscriber() {
	pool := redis.GetRedisPool()
	conn := pool.Get()

	client := goredis.PubSubConn{Conn: conn}
	defer func() {
		_ = client.Close()
	}()

	if err := client.Subscribe("webhook_topic"); err != nil {
		return
	}

	go func() {
		for {
			switch res := client.Receive().(type) {
			case goredis.Message:

			case goredis.Subscription:

			default:
				klog.Errorf("subscribe error: %s", res.(error).Error())
			}

		}
	}()

}
