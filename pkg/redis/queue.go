package redis

import (
	goredis "github.com/gomodule/redigo/redis"
	"github.com/ligao-cloud-native/ecr-apiserver/pkg/config"
)

func GetRedisPool() *goredis.Pool {
	return &goredis.Pool{
		MaxIdle: 6,
		Wait:    true,
		Dial: func() (goredis.Conn, error) {
			conn, err := goredis.Dial("tcp", config.Conf.Redis.Url)
			if err != nil {
				return nil, err
			}

			if _, err := conn.Do("AUTH", config.Conf.Redis.Password); err != nil {
				_ = conn.Close()
				return nil, err
			}

			if _, err := conn.Do("SELECT", config.Conf.Redis.DbNum); err != nil {
				_ = conn.Close()
				return nil, err
			}

			return conn, nil
		},
	}
}
