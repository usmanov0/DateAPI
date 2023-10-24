package handlers

import (
	"fmt"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Tokentime struct {
	Time1 time.Duration
}

var secretKey = []byte("iman")

func CreateToken(tokenParams *Tokentime) (string, error) {
    claims := jwt.StandardClaims{
        ExpiresAt: time.Now().Add(tokenParams.Time1).Unix(),
        IssuedAt:  time.Now().Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token_str, err := token.SignedString(secretKey)
    if err != nil {
        fmt.Println("Error creating token:", err)
        return "", err
    }
    return token_str, nil
}


func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid Signing method")
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("Token isn't Valid")
}

func (h *handlerV1) TokenMiddleWare(ctx *gin.Context) {
	{
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "unauthorized",
				"message": "not allowed",
			})
			return
		}

		claims, err := verifyToken(token)
		if err != nil {
			h.log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "bad token",
				"message": "not allowed",
			})
			return
		}
		fmt.Print(claims)
		ctx.Next()
	}
}
