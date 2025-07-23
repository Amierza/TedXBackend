package routes

import (
	"github.com/Amierza/TedXBackend/handler"
	"github.com/Amierza/TedXBackend/middleware"
	"github.com/Amierza/TedXBackend/service"
	"github.com/gin-gonic/gin"
)

func User(route *gin.Engine, userHandler handler.IUserHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/user")
	{
		// Authentication
		routes.POST("/login", userHandler.Login)

		// Ticket
		routes.GET("/get-all-ticket", userHandler.GetAllTicket)

		// Sponsorship
		routes.GET("/get-all-sponsorship", userHandler.GetAllSponsorship)

		// Speaker
		routes.GET("/get-all-speaker", userHandler.GetAllSpeaker)

		// Merch
		routes.GET("/get-all-merch", userHandler.GetAllMerch)

		// Bundle
		// routes.GET("/get-all-bundle", userHandler.GetAllBundle)

		routes.Use(middleware.Authentication(jwtService))
		{
			// User
			routes.GET("/get-detail-user", userHandler.GetDetailUser)
		}
	}
}
