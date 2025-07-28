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
			// User
			routes.POST("/create-user", adminHandler.CreateUser)
			routes.GET("/get-all-user", adminHandler.GetAllUser)
			routes.GET("/get-detail-user/:id", adminHandler.GetDetailUser)
			routes.PATCH("/update-user/:id", adminHandler.UpdateUser)
			routes.DELETE("/delete-user/:id", adminHandler.DeleteUser)

			// Ticket
			routes.POST("/create-ticket", adminHandler.CreateTicket)
			routes.GET("/get-all-ticket", adminHandler.GetAllTicket)
			routes.GET("/get-detail-ticket/:id", adminHandler.GetDetailTicket)
			routes.PATCH("/update-ticket/:id", adminHandler.UpdateTicket)
			routes.DELETE("/delete-ticket/:id", adminHandler.DeleteTicket)

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

			// Bundle
			routes.POST("/create-bundle", adminHandler.CreateBundle)
			routes.GET("/get-all-bundle", adminHandler.GetAllBundle)
			routes.GET("/get-detail-bundle/:id", adminHandler.GetDetailBundle)
			routes.PATCH("/update-bundle/:id", adminHandler.UpdateBundle)
			routes.DELETE("/delete-bundle/:id", adminHandler.DeleteBundle)

			// Student Ambassador
			routes.POST("/create-student-ambassador", adminHandler.CreateStudentAmbassador)
			routes.GET("/get-all-student-ambassador", adminHandler.GetAllStudentAmbassador)
			routes.GET("/get-detail-student-ambassador/:id", adminHandler.GetDetailStudentAmbassador)
			routes.PATCH("/update-student-ambassador/:id", adminHandler.UpdateStudentAmbassador)
			routes.DELETE("/delete-student-ambassador/:id", adminHandler.DeleteStudentAmbassador)

			// Transaction & Ticket Form
			routes.POST("/create-transaction-ticket", adminHandler.CreateTransactionTicket)
			routes.GET("/get-all-transaction-ticket", adminHandler.GetAllTransactionTicket)
			routes.GET("/get-detail-transaction-ticket/:id", adminHandler.GetDetailTransactionTicket)
		}
	}
}
