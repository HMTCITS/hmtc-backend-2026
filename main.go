package main

import (
	"fmt"
	"os"

	"github.com/HMTCITS/hmtc-backend-2025/controller"
	_ "github.com/HMTCITS/hmtc-backend-2025/docs"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/HMTCITS/hmtc-backend-2025/migration"
	"github.com/HMTCITS/hmtc-backend-2025/router"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

// @title HMTC API Documentation
// @version 1.0
// @description API Documentation
// @host localhost:5000
// @BasePath /api
func main() {
	fmt.Println("Backend HMTC 2025")

	config.LoadConfig()

	var db *gorm.DB = config.ConnectDatabase()

	defer config.CloseDatabase(db)

	var (
		userRepo       repository.UserRepository       = repository.NewUserRepository(db)
		shortLinkRepo  repository.ShortLinkRepository  = repository.NewShortLinkRepository(db)
		oauthTokenRepo repository.OAuthTokenRepository = repository.NewOAuthTokenRepo(db)

		userService      service.UserService      = service.NewUserService(userRepo)
		shortLinkService service.ShortLinkService = service.NewShortLinkService(shortLinkRepo)
		driveService     service.DriveService     = service.NewDriveService(oauthTokenRepo)
		sheetsService    service.SheetsService    = service.NewSheetsService(oauthTokenRepo)
		magangService    service.MagangService    = service.NewMagangService(driveService, sheetsService)

		userController      controller.UserController      = controller.NewUserController(userService)
		healthController    controller.HealthController    = controller.NewHealthController()
		shortLinkController controller.ShortLinkController = controller.NewShortLinkController(shortLinkService)
		magangController    controller.MagangController    = controller.NewMagangController(magangService, oauthTokenRepo)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// start schedule poller (if configured)
	middleware.StartSchedulePoller()
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.User(server, userController)
	router.ShortLink(server, shortLinkController)
	router.Magang(server, magangController)
	router.Health(server, healthController)

	if err := migration.Migrate(db); err != nil {
		panic("Failed to migrate database")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "5000"
	}

	if err := server.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
