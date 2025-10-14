package main

import (
	"fmt"
	"os"

	_ "github.com/HMTCITS/hmtc-backend-2025/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/HMTCITS/hmtc-backend-2025/config"
	"github.com/HMTCITS/hmtc-backend-2025/controller"
	"github.com/HMTCITS/hmtc-backend-2025/middleware"
	"github.com/HMTCITS/hmtc-backend-2025/migration"
	"github.com/HMTCITS/hmtc-backend-2025/repository"
	"github.com/HMTCITS/hmtc-backend-2025/router"
	"github.com/HMTCITS/hmtc-backend-2025/service"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.ConnectDatabase()

	userRepo        repository.UserRepository        = repository.NewUserRepository(db)
	shortLinkRepo   repository.ShortLinkRepository   = repository.NewShortLinkRepository(db)
	fileTARepo      repository.FileTARepository      = repository.NewFileTARepository(db)
	userFileReqRepo repository.UserFileReqRepository = repository.NewUserFileRepository(db)

	userService        service.UserService        = service.NewUserService(userRepo)
	shortLinkService   service.ShortLinkService   = service.NewShortLinkService(shortLinkRepo)
	fileTAService      service.FileTAService      = service.NewFileTAService(fileTARepo, userFileReqRepo)
	userFileReqService service.UserFileReqService = service.NewUserFileService(userFileReqRepo)

	userController        controller.UserController        = controller.NewUserController(userService)
	healthController      controller.HealthController      = controller.NewHealthController()
	shortLinkController   controller.ShortLinkController   = controller.NewShortLinkController(shortLinkService)
	fileTAController      controller.FileTAController      = controller.NewFileTAController(fileTAService)
	userFileReqController controller.UserFileReqController = controller.NewUserFileReqController(userFileReqService)
)

// @title	hmtc documentation
// @version 1.0
// @description API hmtc link shortener dan user
// @host localhost:3000
// @BasePath /api
func main() {
	fmt.Println("Backend HMTC 2025")

	defer config.CloseDatabase(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	router.User(server, userController)
	router.ShortLink(server, shortLinkController)
	// add swagger
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Health(server, healthController)
	router.FileTA(server, fileTAController)
	router.UserFileReq(server, userFileReqController)
	server.MaxMultipartMemory = 10 << 20

	for _, arg := range os.Args {
		if arg == "--migrate" {
			if err := migration.Migrate(db); err != nil {
				panic("Failed to migrate database")
			}
			return
		}
	}

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
