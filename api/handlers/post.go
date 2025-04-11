package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/alirahmani93/instagram-bot/db/models"
    "gorm.io/gorm"
)

// @Summary Create a post to monitor
// @Description Add a post to monitor for a specific keyword and send a predefined DM
// @Tags Posts
// @Accept json
// @Produce json
// @Param post body models.PostToMonitor true "Post data"
// @Success 201 {object} models.PostToMonitor
// @Failure 400 {object} map[string]string
// @Router /posts [post]
func CreatePost(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var post models.PostToMonitor
        if err := c.ShouldBindJSON(&post); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Create(&post).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create post: " + err.Error()})
            return
        }

        c.JSON(http.StatusCreated, post)
    }
}
