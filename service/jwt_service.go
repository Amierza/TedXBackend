package service

import (
	"fmt"
	"os"
	"time"

	"github.com/Amierza/TedXBackend/dto"
	"github.com/golang-jwt/jwt/v5"
)

type (
	IJWTService interface {
		GenerateToken(userID string, role string) (string, error)
		ValidateToken(token string) (*jwt.Token, error)
		GetUserIDByToken(tokenString string) (string, error)
		GetRoleIDByToken(tokenString string) (string, error)
	}

	jwtCustomClaim struct {
		UserID   string `json:"user_id"`
		RoleName string `json:"role_name"`
		jwt.RegisteredClaims
	}

	JWTService struct {
		secretKey string
		issuer    string
	}
)

func NewJWTService() *JWTService {
	return &JWTService{
		secretKey: getSecretKey(),
		issuer:    "Template",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}

	return secretKey
}

func (j *JWTService) GenerateToken(userID string, roleName string) (string, error) {
	claims := jwtCustomClaim{
		userID,
		roleName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 3600)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		return "", dto.ErrGenerateToken
	}

	return tokenString, nil
}

func (j *JWTService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, dto.ErrUnexpectedSigningMethod
	}

	return []byte(j.secretKey), nil
}

func (j *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, j.parseToken)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (j *JWTService) GetUserIDByToken(tokenString string) (string, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", dto.ErrValidateToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", dto.ErrTokenInvalid
	}

	userID := fmt.Sprintf("%v", claims["user_id"])

	return userID, nil
}

func (j *JWTService) GetRoleIDByToken(tokenString string) (string, error) {
	token, err := j.ValidateToken(tokenString)
	if err != nil {
		return "", dto.ErrValidateToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", dto.ErrTokenInvalid
	}

	roleID := fmt.Sprintf("%v", claims["role_id"])

	return roleID, nil
}
