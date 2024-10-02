package main

import (
	"StudentServicePlatform/internal/middleware"
	"StudentServicePlatform/internal/pkg/database"
	"StudentServicePlatform/internal/router"
	"StudentServicePlatform/internal/service"
	"StudentServicePlatform/pkg/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger()
	db := database.Init()
	service.ServiceInit(db)
	r := gin.Default()
	r.Use(middleware.ErrHandler())
	r.NoMethod(middleware.HandleNotFond)
	r.NoRoute(middleware.HandleNotFond)
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal(err)
	}
}
