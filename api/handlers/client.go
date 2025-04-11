package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/ali93rahmani/instagram-bot/db/models"
    "gorm.io/gorm"
)

// @Summary Create a new client
// @Description Register a new Instagram client with username and access token
// @Tags Clients
// @Accept json
// @Produce json
// @Param client body models.Client true "Client data"
// @Success 201 {object} models.Client
// @Failure 400 {object} map[string]string
// @Router /clients [post]
func CreateClient(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var client models.Client
        if err := c.ShouldBindJSON(&client); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        if err := db.Create(&client).Error; err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create client: " + err.Error()})
            return
        }

        c.JSON(http.StatusCreated, client)
    }
}
