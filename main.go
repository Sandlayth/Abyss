package main

import (
    "github.com/sandlayth/abyss/models"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    models.ConnectDatabase()
    router.Run()
}
