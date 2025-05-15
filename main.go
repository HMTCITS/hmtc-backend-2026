package main

import (
	"fmt"
	"os"

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

	userRepo repository.UserRepository = repository.NewUserRepository(db)

	userService service.UserService = service.NewUserService(userRepo)

	userController   controller.UserController   = controller.NewUserController(userService)
	healthController controller.HealthController = controller.NewHealthController()
)

func main() {
	fmt.Println("Backend HMTC 2025")

	defer config.CloseDatabase(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())
	router.User(server, userController)
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
