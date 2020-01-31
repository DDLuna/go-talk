package main

import "github.com/gin-gonic/gin"

type returnStruct struct {
	Message string `json:"msg"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		a := returnStruct{"pong"}
		c.JSON(200, a)
	})
	r.Run()
}
