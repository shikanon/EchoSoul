package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/shikanon/EchoSoul/models"
    "gorm.io/gorm"
)

// CreatePodcastRequest 描述了创建播客请求的 JSON 结构
type CreatePodcastRequest struct {
    Title             string   `json:"title" binding:"required"`
    Description       string   `json:"description"`
    Tags              []string `json:"tags"`
    ImageURL          string   `json:"imageUrl"`
    CreatedBy         uint     `json:"createdBy" binding:"required"`
    UploaderType      string   `json:"uploaderType" binding:"required"`
    CoverImage        string   `json:"coverImage"`
}

// CreatePodcastResponse 描述了创建播客响应的 JSON 结构
type CreatePodcastResponse struct {
    StatusCode int    `json:"statusCode"`
    Message    string `json:"message"`
    PodcastID  uint   `json:"podcastId,omitempty"`
}

// @Summary 创建播客
// @Description 创建一个新的播客
// @Accept  json
// @Produce  json
// @Param createPodcastRequest body CreatePodcastRequest true "创建播客请求"
// @Success 200 {object} CreatePodcastResponse
// @Failure 400 {object} CreatePodcastResponse
// @Failure 401 {object} CreatePodcastResponse
// @Failure 500 {object} CreatePodcastResponse
// @Router /api/podcast/create [post]
func CreatePodcastHandler(c *gin.Context) {
    var request CreatePodcastRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, CreatePodcastResponse{StatusCode: http.StatusBadRequest, Message: "Invalid request body"})
        return
    }

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, CreatePodcastResponse{StatusCode: http.StatusInternalServerError, Message: "Database connection not found"})
        return
    }
    database := db.(*gorm.DB)

    podcast := models.Podcast{
        Title:             request.Title,
        Description:       request.Description,
        Tags:              request.Tags,
        ImageURL:          request.ImageURL,
        CreatedBy:         request.CreatedBy,
        UploaderType:      request.UploaderType,
        CoverImage:        request.CoverImage,
        SubscriptionCount: 0,
        EpisodeCount:      0,
        Focus:             0,
        Display:           0,
        Score:             0.0,
        Subscribed:        false,
    }

    if err := database.Create(&podcast).Error; err != nil {
        c.JSON(http.StatusInternalServerError, CreatePodcastResponse{StatusCode: http.StatusInternalServerError, Message: "Failed to create podcast"})
        return
    }

    c.JSON(http.StatusOK, CreatePodcastResponse{
        StatusCode: http.StatusOK,
        Message:    "Podcast created successfully",
        PodcastID:  podcast.ID,
    })
}