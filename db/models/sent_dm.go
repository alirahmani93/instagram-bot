package models

import "time"

type SentDM struct {
    BaseModel
    CommentID string    `gorm:"uniqueIndex" json:"comment_id"`
    ClientID  uint      `json:"client_id"`
    SentAt    time.Time `json:"sent_at"`
}
