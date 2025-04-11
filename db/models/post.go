package models

type PostToMonitor struct {
    BaseModel
    ClientID          uint   `json:"client_id"`
    PostID            string `json:"post_id"`
    Keyword           string `json:"keyword"`
    PredefinedMessage string `json:"predefined_message"`
}
