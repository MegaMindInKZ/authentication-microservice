package main

import (
	"github.com/MegaMindInKZ/authentication-microservice.git/service-1/controller"
	"github.com/MegaMindInKZ/authentication-microservice.git/service-1/middlewares"
	models "github.com/MegaMindInKZ/authentication-microservice.git/service-1/model"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	apiRoutingGroup := engine.Group("/api")

	apiRoutingGroup.POST("/register", controller.Register)
	apiRoutingGroup.POST("/login", controller.Login)

	adminRoutingGroup := engine.Group("/api/admin")
	adminRoutingGroup.Use(middlewares.JwtAuthMiddleware())
	engine.Run(":3000")
}

func init() {
	models.ConnectDataBase()
	models.DB.AutoMigrate(&models.User{})
}
