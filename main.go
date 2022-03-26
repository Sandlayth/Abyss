package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sandlayth/abyss/models"
    "github.com/sandlayth/abyss/controllers"
)

func main() {
    router := gin.Default()
    models.ConnectDatabase()
    router.GET("/operations", controllers.FindOperations)
    router.GET("/operations/:id", controllers.FindOperation)
    router.POST("/operations", controllers.CreateOperation)
    router.PATCH("/operations/:id", controllers.UpdateOperation)
    router.DELETE("/operations/:id", controllers.DeleteOperation)
    router.Run()
}
