package server

import (
	"github.com/ThailanTec/swagger-poc/controller"
	"github.com/ThailanTec/swagger-poc/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/user", controller.CreateUser)
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Title = "Swagger Thailan Teste"
	docs.SwaggerInfo.Description = "This is a sample server Test Serve."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "http://localhost:8001"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router.GET("/", controller.BVindo)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
