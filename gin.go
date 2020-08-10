package main

import (
	"college/api"
	"college/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.StaticFS("/page", http.Dir("./pages"))
	router.StaticFS("/image", http.Dir("./resources"))
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	err := database.Init() //データベースを初期化
	if err != nil {
		log.Println("DB ERROR:", err)
	} else {
		api.InitTop(router)
		router.GET("/test", func(c *gin.Context) {
			name := c.Query("name")
			result := name + "ありがとうございます"
			c.Writer.WriteString(result)
		})
		router.Run()
	}
}
