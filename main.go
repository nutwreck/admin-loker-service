package main

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-loker-service/models"
	"github.com/nutwreck/admin-loker-service/pkg"
	"github.com/nutwreck/admin-loker-service/routes"

	_ "github.com/nutwreck/admin-loker-service/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Admin Loker API
//	@version		1.0
//	@description	Dokumentasi untuk Service API Admin Loker

//  @Schemes http https

//	@securityDefinitions.basic	BasicAuth

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				User JWT Bearer - Format Input Value : Bearer[ space ][ JWTToken ]

func main() {

	/**
	* ========================
	*  Setup Application
	* ========================
	 */

	db := setupDatabase()
	app := setupApp()

	/**
	* ========================
	* Initialize All Route
	* ========================
	 */

	routes.NewRouteUser(db, app)
	routes.NewRoutePendidikan(db, app)
	routes.NewRouteLevelPekerjaan(db, app)
	routes.NewRouteTipePekerjaan(db, app)
	routes.NewRouteTahunPengalaman(db, app)
	routes.NewRouteKategoriPekerjaan(db, app)
	routes.NewRouteKeahlian(db, app)
	routes.NewRouteJenisPerusahaan(db, app)
	routes.NewRouteConstant(app)
	routes.NewRouteWilayah(db, app)

	/**
	* ========================
	*  Listening Server Port
	* ========================
	 */

	err := app.Run(":" + pkg.GodotEnv("PORT"))

	if err != nil {
		defer logrus.Error("Server is not running")
		logrus.Fatal(err)
	}
}

/**
* ========================
* Database Setup
* ========================
 */

func setupDatabase() *gorm.DB {
	var dsn string
	if pkg.GodotEnv("GO_ENV") == "release" {
		dsn = "host=" + pkg.GodotEnv("POSTGRES_HOST_PROD") + " user=" + pkg.GodotEnv("POSTGRES_USER_PROD") + " password=" + pkg.GodotEnv("POSTGRES_PASSWORD_PROD") + " dbname=" + pkg.GodotEnv("POSTGRES_DB_PROD") + " port=" + pkg.GodotEnv("POSTGRES_PORT_PROD") + " sslmode=" + pkg.GodotEnv("POSTGRES_SSL_PROD")
	} else {
		dsn = "host=" + pkg.GodotEnv("POSTGRES_HOST") + " user=" + pkg.GodotEnv("POSTGRES_USER") + " password=" + pkg.GodotEnv("POSTGRES_PASSWORD") + " dbname=" + pkg.GodotEnv("POSTGRES_DB") + " port=" + pkg.GodotEnv("POSTGRES_PORT") + " sslmode=" + pkg.GodotEnv("POSTGRES_SSL")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})

	if err != nil {
		defer logrus.Info("Database connection failed")
		logrus.Fatal(err)
		return nil
	}

	//  Initialize all model for auto migration here
	err = db.AutoMigrate(
		&models.ModelUser{},
		&models.ModelJenisPerusahaan{},
		&models.ModelKategoriPekerjaan{},
		&models.ModelKeahlian{},
		&models.ModelLevelPekerjaan{},
		&models.ModelPendidikan{},
		&models.ModelTahunPengalaman{},
		&models.ModelTipePekerjaan{},
		&models.ModelKelurahan{},
		&models.ModelKecamatan{},
		&models.ModelKabupaten{},
		&models.ModelProvinsi{},
		&models.ModelNegara{},
	)

	if err != nil {
		defer logrus.Info("Database migration failed")
		logrus.Fatal(err)
		return nil
	}

	return db
}

/**
* ========================
* Application Setup
* ========================
 */

func setupApp() *gin.Engine {

	app := gin.Default()

	if pkg.GodotEnv("GO_ENV") != "development" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize all middleware here
	app.Use(helmet.Default())
	app.Use(gzip.Gzip(gzip.BestCompression))
	app.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "DELETE", "PATCH", "PUT", "OPTIONS"},
		AllowHeaders:    []string{"Content-Type", "Authorization", "Accept-Encoding"},
	}))

	//Docs Swagger Without Model Section
	app.GET("/d29ya2Vyc3VjaA/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1))) //index.html

	return app
}
