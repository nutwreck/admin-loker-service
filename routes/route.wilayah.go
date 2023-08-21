package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nutwreck/admin-loker-service/handlers"
	"github.com/nutwreck/admin-loker-service/repositories"
	"github.com/nutwreck/admin-loker-service/services"
	"gorm.io/gorm"
)

func NewRouteWilayah(db *gorm.DB, router *gin.Engine) {
	route := router.Group("/api/v1/wilayah")

	//NEGARA
	repositoryNegara := repositories.NewRepositoryNegara(db)
	serviceNegara := services.NewServiceNegara(repositoryNegara)
	handlerNegara := handlers.NewHandlerNegara(serviceNegara)
	route.POST("/negara/create", handlerNegara.HandlerCreate)
	route.GET("/negara/results", handlerNegara.HandlerResults)
	route.GET("/negara/result/:code_negara", handlerNegara.HandlerResult)
	route.DELETE("/negara/delete/:code_negara", handlerNegara.HandlerDelete)
	route.PUT("/negara/update/:code_negara", handlerNegara.HandlerUpdate)

	//PROVINSI
	repositoryProvinsi := repositories.NewRepositoryProvinsi(db)
	serviceProvinsi := services.NewServiceProvinsi(repositoryProvinsi)
	handlerProvinsi := handlers.NewHandlerProvinsi(serviceProvinsi)
	route.POST("/provinsi/create", handlerProvinsi.HandlerCreate)
	route.GET("/provinsi/results", handlerProvinsi.HandlerResults)
	route.GET("/provinsi/result/:code_provinsi", handlerProvinsi.HandlerResult)
	route.DELETE("/provinsi/delete/:code_provinsi", handlerProvinsi.HandlerDelete)
	route.PUT("/provinsi/update/:code_provinsi", handlerProvinsi.HandlerUpdate)

	//KABUPATEN
	repositoryKabupaten := repositories.NewRepositoryKabupaten(db)
	serviceKabupaten := services.NewServiceKabupaten(repositoryKabupaten)
	handlerKabupaten := handlers.NewHandlerKabupaten(serviceKabupaten)
	route.POST("/kabupaten/create", handlerKabupaten.HandlerCreate)
	route.GET("/kabupaten/results", handlerKabupaten.HandlerResults)
	route.GET("/kabupaten/result/:code_kabupaten", handlerKabupaten.HandlerResult)
	route.DELETE("/kabupaten/delete/:code_kabupaten", handlerKabupaten.HandlerDelete)
	route.PUT("/kabupaten/update/:code_kabupaten", handlerKabupaten.HandlerUpdate)

	//KECAMATAN
	repositoryKecamatan := repositories.NewRepositoryKecamatan(db)
	serviceKecamatan := services.NewServiceKecamatan(repositoryKecamatan)
	handlerKecamatan := handlers.NewHandlerKecamatan(serviceKecamatan)
	route.POST("/kecamatan/create", handlerKecamatan.HandlerCreate)
	route.GET("/kecamatan/results", handlerKecamatan.HandlerResults)
	route.GET("/kecamatan/result/:code_kecamatan", handlerKecamatan.HandlerResult)
	route.DELETE("/kecamatan/delete/:code_kecamatan", handlerKecamatan.HandlerDelete)
	route.PUT("/kecamatan/update/:code_kecamatan", handlerKecamatan.HandlerUpdate)

	//KELURAHAN
	repositoryKelurahan := repositories.NewRepositoryKelurahan(db)
	serviceKelurahan := services.NewServiceKelurahan(repositoryKelurahan)
	handlerKelurahan := handlers.NewHandlerKelurahan(serviceKelurahan)
	route.POST("/kelurahan/create", handlerKelurahan.HandlerCreate)
	route.GET("/kelurahan/results", handlerKelurahan.HandlerResults)
	route.GET("/kelurahan/result/:code_kelurahan", handlerKelurahan.HandlerResult)
	route.DELETE("/kelurahan/delete/:code_kelurahan", handlerKelurahan.HandlerDelete)
	route.PUT("/kelurahan/update/:code_kelurahan", handlerKelurahan.HandlerUpdate)
}
