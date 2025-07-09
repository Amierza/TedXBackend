package routes

import (
	"github.com/Amierza/TedXBackend/handler"
	"github.com/Amierza/TedXBackend/service"
	"github.com/gin-gonic/gin"
)

func Admin(route *gin.Engine, adminHandler handler.IAdminHandler, jwtService service.IJWTService) {
	routes := route.Group("/api/v1/admin")
	{
		// Authentication
		routes.POST("/login", adminHandler.Login)
	}
}
