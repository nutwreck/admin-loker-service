package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/configs"
	"github.com/nutwreck/admin-loker-service/handlers"
	"github.com/nutwreck/admin-loker-service/middlewares"
)

func NewRouteConstant(router *gin.Engine) {
	route := router.Group("/api/v1/constant")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(configs.RoleConfig))

	route.GET("/jenis-kelamin", handlers.HandlerJenisKelamin)
	route.GET("/status-pernikahan", handlers.HandlerStatusPernikahan)
}
