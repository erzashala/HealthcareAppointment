package main

import (
    "healthcare-api/database"
    "healthcare-api/routes"

    _ "healthcare-api/docs"

    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Healthcare Appointment API
// @version 1.0
// @description REST API for managing patients and appointments
// @host localhost:8080
// @BasePath /

func main() {
    database.ConnectDatabase()
    r := routes.SetupRouter()
    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    r.Run(":8080")
}