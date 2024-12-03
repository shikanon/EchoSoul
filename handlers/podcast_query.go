package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/shikanon/EchoSoul/models"
    "gorm.io/gorm"
)

// QueryPodcastRequest 描述了查询播客请求的 JSON 结构
type QueryPodcastRequest struct {
    PageNum       int    `json:"page_num" binding:"required"`
    PageMaxItems  int    `json:"page_max_items" binding:"required"`
    UserID        *uint  `json:"user,omitempty"`
    CatalogID     *uint  `json:"catalogId,omitempty"`
    TagID         *uint  `json:"tagId,omitempty"`
    Content       string `json:"content,omitempty"`
}

// QueryPodcastResponse 描述了查询播客响应的 JSON 结构
type QueryPodcastResponse struct {
    StatusCode int                 `json:"statusCode"`
    Data       []QueryPodcastData `json:"data,omitempty"`
}

// QueryPodcastData 描述了响应中的数据部分
type QueryPodcastData struct {
    PodcastID    uint     `json:"podcastId"`
    PodcastName  string   `json:"podcastName"`
    Description  string   `json:"description"`
    Tags         []string `json:"tags"`
    ImageURL     string   `json:"imageUrl"`
    Display      int      `json:"display"`
    Focus        int      `json:"focus"`
    Score        float64  `json:"score"`
    EpisodeCount int      `json:"episodeCount"`
    Subscribed   bool     `json:"subscribed"`
}

// @Summary 查询播客列表
// @Description 查询播客列表，支持过滤和分页，过滤条件支持查询自己订阅的播客列表
// @Accept  json
// @Produce  json
// @Param	queryPodcastRequest	body	QueryPodcastRequest	true	"查询播客请求"
// @Success 200 {object} QueryPodcastResponse
// @Failure 400 {object} QueryPodcastResponse
// @Failure 401 {object} QueryPodcastResponse
// @Failure 500 {object} QueryPodcastResponse
// @Router /api/podcast/query [post]
func QueryPodcastHandler(c *gin.Context) {
    var request QueryPodcastRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, QueryPodcastResponse{StatusCode: http.StatusBadRequest})
        return
    }

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, QueryPodcastResponse{StatusCode: http.StatusInternalServerError})
        return
    }
    database := db.(*gorm.DB)

    var podcasts []models.Podcast
    query := database.Preload("Episodes")

    if request.UserID != nil {
        var subscriptions []models.UserSubscription
        if err := query.Where("user_id = ?", request.UserID).Find(&subscriptions).Error; err != nil {
            c.JSON(http.StatusInternalServerError, QueryPodcastResponse{StatusCode: http.StatusInternalServerError})
            return
        }
        var podcastIDs []uint
        for _, sub := range subscriptions {
            podcastIDs = append(podcastIDs, sub.PodcastID)
        }
        query = query.Where("id IN ?", podcastIDs)
    }

    if request.CatalogID != nil {
        query = query.Where("catalog_id = ?", request.CatalogID)
    }

    if request.TagID != nil {
        query = query.Where("tags @> ?", request.TagID)
    }

    if request.Content != "" {
        query = query.Where("title LIKE ?", "%"+request.Content+"%")
    }

    if err := query.Offset((request.PageNum - 1) * request.PageMaxItems).Limit(request.PageMaxItems).Find(&podcasts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, QueryPodcastResponse{StatusCode: http.StatusInternalServerError})
        return
    }

    var responsePodcasts []QueryPodcastData
    for _, podcast := range podcasts {
        responsePodcasts = append(responsePodcasts, QueryPodcastData{
            PodcastID:    podcast.ID,
            PodcastName:  podcast.Title,
            Description:  podcast.Description,
            Tags:         podcast.Tags,
            ImageURL:     podcast.ImageURL,
            Display:      podcast.Display,
            Focus:        podcast.Focus,
            Score:        podcast.Score,
            EpisodeCount: podcast.EpisodeCount,
            Subscribed:   podcast.Subscribed,
        })
    }

    c.JSON(http.StatusOK, QueryPodcastResponse{
        StatusCode: http.StatusOK,
        Data:       responsePodcasts,
    })
}