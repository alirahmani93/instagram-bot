package instagram

import (
    "encoding/json"
    "net/http"
    "io/ioutil"
    "github.com/alirahmani93/instagram-bot/db/models"
    "gorm.io/gorm"
)

type Comment struct {
    ID        string `json:"id"`
    Text      string `json:"text"`
    Timestamp string `json:"timestamp"`
    Username  string `json:"username"`
    LikeCount int    `json:"like_count"`
}

type CommentsResponse struct {
    Data   []Comment `json:"data"`
    Paging struct {
        Next string `json:"next"`
    } `json:"paging"`
}

func GetComments(postID, accessToken string) ([]Comment, error) {
    var allComments []Comment
    url := "https://graph.instagram.com/v13.0/" + postID + "/comments?fields=id,text,timestamp,username,like_count&access_token=" + accessToken

    for {
        resp, err := http.Get(url)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        var commentsResp CommentsResponse
        if err := json.Unmarshal(body, &commentsResp); err != nil {
            return nil, err
        }

        allComments = append(allComments, commentsResp.Data...)

        if commentsResp.Paging.Next == "" {
            break
        }
        url = commentsResp.Paging.Next
    }

    return allComments, nil
}

func SendDM(recipientUsername, message, accessToken string) error {
    url := "https://graph.instagram.com/v13.0/me/messages?access_token=" + accessToken
    payload := map[string]interface{}{
        "recipient": map[string]string{"username": recipientUsername},
        "message":   map[string]string{"text": message},
    }
    payloadBytes, _ := json.Marshal(payload)

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
    if err != nil {
        return err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return fmt.Errorf("failed to send DM, status: %d", resp.StatusCode)
    }
    return nil
}

func MonitorComments(db *gorm.DB) {
    var clients []models.Client
    if err := db.Find(&clients).Error; err != nil {
        log.Printf("Error fetching clients: %v", err)
        return
    }

    for _, client := range clients {
        var posts []models.PostToMonitor
        if err := db.Where("client_id = ?", client.ID).Find(&posts).Error; err != nil {
            log.Printf("Error fetching posts for client %d: %v", client.ID, err)
            continue
        }

        for _, post := range posts {
            comments, err := GetComments(post.PostID, client.AccessToken)
            if err != nil {
                log.Printf("Error fetching comments for post %s: %v", post.PostID, err)
                continue
            }

            for _, comment := range comments {
                var sentDM models.SentDM
                if db.Where("comment_id = ?", comment.ID).First(&sentDM).Error == nil {
                    continue // Send predifined Message
                }

                if strings.Contains(comment.Text, post.Keyword) {
                    if err := SendDM(comment.Username, post.PredefinedMessage, client.AccessToken); err != nil {
                        log.Printf("Error sending DM to %s: %v", comment.Username, err)
                        continue
                    }

                    sentDM := models.SentDM{
                        CommentID: comment.ID,
                        ClientID:  client.ID,
                        SentAt:    time.Now(),
                    }
                    if err := db.Create(&sentDM).Error; err != nil {
                        log.Printf("Error saving sent DM: %v", err)
                    }
                }
            }
        }
    }
}
