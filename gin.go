package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/page", http.Dir("./pages"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/test", func(c *gin.Context) {
		name := c.Query("name")
		result := name + "ありがとうございます"
		c.Writer.WriteString(result)
	})
	router.Run()
}
