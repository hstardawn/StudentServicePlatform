package main

import (
	"StudentServicePlatform/internal/middleware"
	"StudentServicePlatform/internal/pkg/database"
	"StudentServicePlatform/internal/router"
	"StudentServicePlatform/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//
	db := database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	r.Use(middleware.ErrHandler())
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
