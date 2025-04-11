package models

type Client struct {
    BaseModel
    InstagramUsername string `gorm:"uniqueIndex" json:"instagram_username"`
    AccessToken       string `json:"access_token"`
}
