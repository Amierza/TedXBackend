package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Amierza/TedXBackend/cmd"
	"github.com/Amierza/TedXBackend/config/database"
	"github.com/Amierza/TedXBackend/config/midtrans"
	"github.com/Amierza/TedXBackend/handler"
	"github.com/Amierza/TedXBackend/middleware"
	"github.com/Amierza/TedXBackend/repository"
	"github.com/Amierza/TedXBackend/routes"
	"github.com/Amierza/TedXBackend/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.SetUpPostgreSQLConnection()
	defer database.ClosePostgreSQLConnection(db)

	fmt.Println("MIDTRANS_SERVER_KEY:", os.Getenv("MIDTRANS_SERVER_KEY"))
	fmt.Println("APP_ENV:", os.Getenv("APP_ENV"))
	midtrans.InitMidtransClient(os.Getenv("MIDTRANS_SERVER_KEY"), os.Getenv("APP_ENV") == "production")

	if len(os.Args) > 1 {
		cmd.Command(db)
		return
	}

	var (
		jwtService = service.NewJWTService()

		userRepo    = repository.NewUserRepository(db)
		userService = service.NewUserService(userRepo, jwtService)
		userHandler = handler.NewUserHandler(userService)

		adminRepo    = repository.NewAdminRepository(db)
		adminService = service.NewAdminService(adminRepo, jwtService)
		adminHandler = handler.NewAdminHandler(adminService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.User(server, userHandler, jwtService)
	routes.Admin(server, adminHandler, jwtService)

	server.Static("/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
