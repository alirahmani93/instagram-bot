package api

import (
    "github.com/gin-gonic/gin"
    "github.com/alirahmani93/instagram-bot/api/handlers"
    "github.com/alirahmani93/instagram-bot/api/middleware"
    "github.com/swaggo/files"
    "github.com/swaggo/gin-swagger"
    "gorm.io/gorm"
    _ "github.com/alirahmani93/instagram-bot/docs"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()
    router.Use(middleware.ErrorHandler())

    // روت‌ها
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Instagram Bot is running"})
    })

    router.POST("/clients", handlers.CreateClient(db))
    router.POST("/posts", handlers.CreatePost(db))

    // Swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    return router
}