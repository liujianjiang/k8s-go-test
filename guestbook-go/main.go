// geektutu.com
// main.go
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
)

var Rdb *redis.Client

// 初始化连接
func InitClient() (err error) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err = Rdb.Ping().Result()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	r := gin.Default()

	fmt.Println("初始化redis……")
	err := InitClient()
	if err != nil {
		fmt.Println(err)
	}

	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Geektutu 3333")
	})
	r.Run(":3000") // listen and serve on 0.0.0.0:3000
}
