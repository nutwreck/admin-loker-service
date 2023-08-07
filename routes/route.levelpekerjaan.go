package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/configs"
	"github.com/nutwreck/admin-loker-service/handlers"
	"github.com/nutwreck/admin-loker-service/middlewares"
	"github.com/nutwreck/admin-loker-service/repositories"
	"github.com/nutwreck/admin-loker-service/services"
	"gorm.io/gorm"
)

func NewRouteLevelPekerjaan(db *gorm.DB, router *gin.Engine) {
	repository := repositories.NewRepositoryLevelPekerjaan(db)
	service := services.NewServiceLevelPekerjaan(repository)
	handler := handlers.NewHandlerLevelPekerjaan(service)

	route := router.Group("/api/v1/level-pekerjaan")
	route.Use(middlewares.AuthToken())
	route.Use(middlewares.AuthRole(configs.RoleConfig))

	router.GET("/api/v1/level-pekerjaan/ping", handler.HandlerPing)
	route.POST("/create", handler.HandlerCreate)
	route.GET("/results", handler.HandlerResults)
	route.GET("/result/:id", handler.HandlerResult)
	route.DELETE("/delete/:id", handler.HandlerDelete)
	route.PUT("/update/:id", handler.HandlerUpdate)
}
