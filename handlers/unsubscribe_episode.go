package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/shikanon/EchoSoul/models"
    "gorm.io/gorm"
)

// UnsubscribeEpisodeRequest 描述了取消收藏播客单曲的请求JSON结构
type UnsubscribeEpisodeRequest struct {
    EpisodeID uint `json:"episodeID" binding:"required" example:"123"`
}

// UnsubscribeEpisodeResponse 描述了取消收藏播客单曲的响应JSON结构
type UnsubscribeEpisodeResponse struct {
    StatusCode int                     `json:"statusCode"`
    Data       *UnsubscribeEpisodeData `json:"data,omitempty"`
}

// UnsubscribeEpisodeData 描述了响应中的数据部分
type UnsubscribeEpisodeData struct {
    Result  string `json:"result"`
    Message string `json:"message"`
}

// @Summary 取消收藏播客单曲
// @Description 用于取消已收藏的节目。如果未收藏，则返回未收藏。
// @Accept  json
// @Produce  json
// @Param   Authorization  header  string  true  "Bearer <token>"
// @Param   UnsubscribeEpisodeRequest  body  UnsubscribeEpisodeRequest  true  "取消收藏请求"
// @Success 200 {object} UnsubscribeEpisodeResponse
// @Failure 400 {object} UnsubscribeEpisodeResponse
// @Failure 401 {object} UnsubscribeEpisodeResponse
// @Failure 500 {object} UnsubscribeEpisodeResponse
// @Router /api/episode/unsubscribe [post]
func UnsubscribeEpisodeHandler(c *gin.Context) {
    var request UnsubscribeEpisodeRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, UnsubscribeEpisodeResponse{
            StatusCode: http.StatusBadRequest,
            Data: &UnsubscribeEpisodeData{
                Result:  "error",
                Message: "Invalid request parameters",
            },
        })
        return
    }

    db, exists := c.Get("db")
    if !exists {
        c.JSON(http.StatusInternalServerError, UnsubscribeEpisodeResponse{
            StatusCode: http.StatusInternalServerError,
            Data: &UnsubscribeEpisodeData{
                Result:  "error",
                Message: "Database connection not found",
            },
        })
        return
    }
    database := db.(*gorm.DB)

    // 从请求头中获取用户ID，假设用户ID是通过JWT解析出来的
    userID := c.GetString("userID")
    if userID == "" {
        c.JSON(http.StatusUnauthorized, UnsubscribeEpisodeResponse{
            StatusCode: http.StatusUnauthorized,
            Data: &UnsubscribeEpisodeData{
                Result:  "error",
                Message: "User not authenticated",
            },
        })
        return
    }

    var userFavorite models.UserFavorite
    if err := database.Where("user_id = ? AND episode_id = ?", userID, request.EpisodeID).First(&userFavorite).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            c.JSON(http.StatusBadRequest, UnsubscribeEpisodeResponse{
                StatusCode: http.StatusBadRequest,
                Data: &UnsubscribeEpisodeData{
                    Result:  "error",
                    Message: "Episode not favorited",
                },
            })
            return
        }
        c.JSON(http.StatusInternalServerError, UnsubscribeEpisodeResponse{
            StatusCode: http.StatusInternalServerError,
            Data: &UnsubscribeEpisodeData{
                Result:  "error",
                Message: "Failed to query database",
            },
        })
        return
    }

    if err := database.Delete(&userFavorite).Error; err != nil {
        c.JSON(http.StatusInternalServerError, UnsubscribeEpisodeResponse{
            StatusCode: http.StatusInternalServerError,
            Data: &UnsubscribeEpisodeData{
                Result:  "error",
                Message: "Failed to unsubscribe episode",
            },
        })
        return
    }

    c.JSON(http.StatusOK, UnsubscribeEpisodeResponse{
        StatusCode: http.StatusOK,
        Data: &UnsubscribeEpisodeData{
            Result:  "success",
            Message: "播客取消成功",
        },
    })
}