package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swagFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/task_iman/pkg/logger"
	v1 "github.com/task_iman/api/handlers"
)

type RoutetOptions struct {
	Log logger.Logger
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func New(option RoutetOptions) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "App is running...",
		})
	})

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log: option.Log,
	})

	api := router.Group("/v1")

	api.GET("/days", handlerV1.TokenMiddleWare, handlerV1.Days)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("swagger/*any", ginSwagger.WrapHandler(swagFile.Handler, url))

	return router
}
