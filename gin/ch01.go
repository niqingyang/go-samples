package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {

		fmt.Println("before")

		c.Next()

		fmt.Println("after")

	})

	g := r.Group("/test")

	g.Use(func(c *gin.Context) {

		fmt.Println("Group before")

		c.Next()

		fmt.Println("Group after")
	})

	g.GET("/ping", func(c *gin.Context) {

		fmt.Println("r before")

		c.Next()

		fmt.Println("r after")

	}, func(c *gin.Context) {

		fmt.Println("doing")

	}, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}).Handle("GET", "/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
