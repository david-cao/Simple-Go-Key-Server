package database

import (
	"flag"
	"github.com/garyburd/redigo/redis"
	"time"
)

var (
	redisAddress   = flag.String("http://localhost", ":6999", "Address to the Redis server")
	redisPassword  = flag.String("", "", "")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
	redisPool      = newRedisPool(*redisAddress, *maxConnections)
)

// Taken from redis docs
func newRedisPool(server string, maxConnections int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		MaxActive:   maxConnections,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			// if _, err := c.Do("AUTH", password); err != nil {
			// 	fmt.Println("Failed to auth: ", err)
			// 	c.Close()
			// 	return nil, err
			// }
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func Get() redis.Conn {
	return redisPool.Get()
}
