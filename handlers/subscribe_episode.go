package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shikanon/EchoSoul/models"
	"gorm.io/gorm"
)

// SubscribeEpisodeRequest 描述了收藏单曲的请求JSON结构
type SubscribeEpisodeRequest struct {
	EpisodeID uint `json:"episodeID" binding:"required"`
}

// SubscribeEpisodeResponse 描述了收藏单曲的响应JSON结构
type SubscribeEpisodeResponse struct {
	StatusCode int                   `json:"statusCode"`
	Data       *SubscribeEpisodeData `json:"data,omitempty"`
}

// SubscribeEpisodeData 描述了响应中的数据部分
type SubscribeEpisodeData struct {
	Result  string `json:"result"`
	Message string `json:"message,omitempty"`
}

// @Summary 收藏播客单曲
// @Description 用于收藏播客中某个节目
// @Accept  json
// @Produce  json
// @Param   episodeID  body    int  true   "要收藏的播客单曲 ID"
// @Header  200 {string} Authorization "Bearer <token>"
// @Success 200 {object} map[string]interface{}
// @Router /api/episode/subscribe [post]
func SubscribeEpisodeHandler(c *gin.Context) {
	var request SubscribeEpisodeRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, SubscribeEpisodeResponse{
			StatusCode: http.StatusBadRequest,
			Data: &SubscribeEpisodeData{
				Result:  "error",
				Message: "收藏失败: Invalid request body",
			},
		})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, SubscribeEpisodeResponse{
			StatusCode: http.StatusInternalServerError,
			Data: &SubscribeEpisodeData{
				Result:  "error",
				Message: "数据库连接失败",
			},
		})
		return
	}
	database := db.(*gorm.DB)

	// 检查节目是否已收藏
	var existingFavorite models.UserFavorite
	userID := c.GetUint("userID") // 假设用户ID从中间件或其他地方获取
	if err := database.Where("user_id = ? AND episode_id = ?", userID, request.EpisodeID).First(&existingFavorite).Error; err == nil {
		c.JSON(http.StatusOK, SubscribeEpisodeResponse{
			StatusCode: http.StatusOK,
			Data: &SubscribeEpisodeData{
				Result:  "success",
				Message: "该节目已收藏",
			},
		})
		return
	}

	// 添加收藏
	newFavorite := models.UserFavorite{
		UserID:     userID,
		EpisodeID:  request.EpisodeID,
		FavoriteAt: time.Now(),
	}
	if err := database.Create(&newFavorite).Error; err != nil {
		c.JSON(http.StatusInternalServerError, SubscribeEpisodeResponse{
			StatusCode: http.StatusInternalServerError,
			Data: &SubscribeEpisodeData{
				Result:  "error",
				Message: "收藏失败",
			},
		})
		return
	}

	c.JSON(http.StatusOK, SubscribeEpisodeResponse{
		StatusCode: http.StatusOK,
		Data: &SubscribeEpisodeData{
			Result:  "success",
			Message: "该节目已成功收藏",
		},
	})
}
