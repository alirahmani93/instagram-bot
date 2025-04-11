package main

import (
    "log"
    "github.com/ali93rahmani/instagram-bot/api"
    "github.com/ali93rahmani/instagram-bot/db"
    "github.com/joho/godotenv"
    "github.com/go-co-op/gocron"
    "time"
    "github.com/ali93rahmani/instagram-bot/internal/instagram"
)

// @title Instagram Bot API
// @version 1.0
// @description A RESTful API for an Instagram bot that monitors comments and sends predefined DMs.
// @host localhost:8080
// @BasePath /
func main() {
    // بارگذاری متغیرهای محیطی
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // اتصال به دیتابیس
    database, err := db.NewDatabase()
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // تنظیم زمان‌بندی برای مانیتورینگ کامنت‌ها
    scheduler := gocron.NewScheduler(time.UTC)
    scheduler.Every(5).Minutes().Do(func() {
        instagram.MonitorComments(database)
    })
    scheduler.StartAsync()

    // راه‌اندازی سرور Gin
    router := api.SetupRouter(database)
    if err := router.Run(":8080"); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}
