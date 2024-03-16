package main

import "github.com/ThailanTec/swagger-poc/server"

// @title Sua API Swagger Title
// @version 1.0
// @version 2.0
// @description Esta é uma API Swagger para um exemplo de projeto Go com Gin.
// @termsOfService https://example.com/terms/
// @contact.email
// @license.name
// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := server.SetupRouter()
	r.Run(":8001")
}
