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
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	err := database.Init() //データベースを初期化
	if err != nil {
		log.Println("DB ERROR:", err)
	} else {
		api.InitTop(router)
		router.Run()
	}
}
