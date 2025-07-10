package middleware

import (
	"net/http"
	"strings"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/Amierza/TedXBackend/service"
	"github.com/Amierza/TedXBackend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RouteAccessControl(jwtService service.IJWTService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_FOUND, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if !strings.Contains(authHeader, "Bearer") {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		authHeader = strings.Replace(authHeader, "Bearer ", "", -1)
		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_NOT_VALID, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		if !token.Valid {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_TOKEN_DENIED_ACCESS, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_GET_CUSTOM_CLAIMS, nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, res)
			return
		}

		roleName, ok := claims["role_name"]
		if !ok {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_GET_ROLE_USER, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		if roleName != "admin" {
			res := utils.BuildResponseFailed(dto.MESSAGE_FAILED_PROSES_REQUEST, dto.MESSAGE_FAILED_ACCESS_DENIED, nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, res)
			return
		}

		ctx.Next()
	}
}
