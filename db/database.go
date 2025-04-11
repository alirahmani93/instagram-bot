package db

import (
    "os"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // مهاجرت خودکار مدل‌ها
    err = db.AutoMigrate(&Client{}, &PostToMonitor{}, &SentDM{})
    if err != nil {
        return nil, err
    }

    return db, nil
}
