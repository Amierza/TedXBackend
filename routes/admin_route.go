package routes

import (
	"github.com/Amierza/TedXBackend/handler"
	"github.com/Amierza/TedXBackend/middleware"
	"github.com/Amierza/TedXBackend/service"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, adminHandler handler.IAdminHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/admin")
	{
		// Authentication
		routes.POST("/login", adminHandler.Login)

		routes.Use(middleware.Authentication(jwtService), middleware.RouteAccessControl(jwtService))
		{
			// Sponsorship
			routes.POST("/create-sponsorship", adminHandler.CreateSponsorship)
			routes.GET("/get-all-sponsorship", adminHandler.GetAllSponsorship)
			routes.GET("/get-detail-sponsorship/:id", adminHandler.GetDetailSponsorship)
			routes.PATCH("/update-sponsorship/:id", adminHandler.UpdateSponsorship)
			routes.DELETE("/delete-sponsorship/:id", adminHandler.DeleteSponsorship)

			// Speaker
			routes.POST("/create-speaker", adminHandler.CreateSpeaker)
			routes.GET("/get-all-speaker", adminHandler.GetAllSpeaker)
			routes.GET("/get-detail-speaker/:id", adminHandler.GetDetailSpeaker)
			routes.PATCH("/update-speaker/:id", adminHandler.UpdateSpeaker)
			routes.DELETE("/delete-speaker/:id", adminHandler.DeleteSpeaker)

			// Merch
			routes.POST("/create-merch", adminHandler.CreateMerch)
			routes.GET("/get-all-merch", adminHandler.GetAllMerch)
			routes.GET("/get-detail-merch/:id", adminHandler.GetDetailMerch)
			routes.PATCH("/update-merch/:id", adminHandler.UpdateMerch)
			routes.DELETE("/delete-merch/:id", adminHandler.DeleteMerch)
		}
	}
}
