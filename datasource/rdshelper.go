// Package datasource 数据源创建
package datasource

import (
	"LuckyGo/conf"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"sync"
	"time"
)

var rdsLock sync.Mutex
var cacheInstance *RedisConn

type RedisConn struct {
	pool      *redis.Pool
	showDebug bool
}

func (rds *RedisConn) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	coon := rds.pool.Get()
	defer coon.Close()

	t1 := time.Now().UnixNano()
	reply, err = coon.Do(commandName, args...)
	if err != nil {
		e := coon.Err()
		if e != nil {
			log.Fatal("rdshelper.Do", err, e)
		}
	}
	t2 := time.Now().UnixNano()
	if rds.showDebug {
		fmt.Printf("[redis] [info] [%dus]=%s,err=%s,args=%v,reply=%s\n",
			(t2-t1)/1000, commandName, err, args, reply)
	}
	return reply, err
}
func (rds *RedisConn) ShowDebug(b bool) {
	rds.showDebug = b
}

func InstanceCache() *RedisConn {
	if cacheInstance != nil {
		return cacheInstance
	}
	rdsLock.Lock()
	defer rdsLock.Unlock()

	if cacheInstance != nil {
		return cacheInstance
	}
	return NewCache()
}

// 实现连接池
func NewCache() *RedisConn {
	pool := redis.Pool{Dial: func() (redis.Conn, error) {
		c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", conf.RdsMaster.Host, conf.RdsMaster.Port))
		if err != nil {
			log.Fatal("rdshelper.NewCache Dial error=", err)
			return nil, err
		}
		return c, nil
	},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         1000,
		MaxActive:       10000,
		IdleTimeout:     0,
		Wait:            false,
		MaxConnLifetime: 0,
	}
	instance := &RedisConn{
		pool: &pool,
	}
	cacheInstance = instance
	instance.ShowDebug(true)
	return instance
}
