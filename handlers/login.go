package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/shikanon/EchoSoul/models"
	"gorm.io/gorm"
)

// LoginRequest 描述了用户登录的请求JSON结构
type LoginRequest struct {
	PhoneNum string `json:"phoneNum" binding:"required" example:"12345678901"`
}

// LoginResponse 描述了用户登录的响应JSON结构
type LoginResponse struct {
	StatusCode int        `json:"statusCode"`
	Data       *LoginData `json:"data,omitempty"`
}

// LoginData 描述了响应中的数据部分
type LoginData struct {
	Token string `json:"token"`
}

// 这里定义一个 JWT 密钥，实际应用中请使用更安全的方式管理密钥
var jwtKey = []byte("your_secret_key")

// GenerateJWT 生成JWT
func GenerateJWT(phoneNum string) (string, error) {
	// 设置 JWT 过期时间（例如 24 小时）
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &jwt.RegisteredClaims{
		Subject:   phoneNum,
		ExpiresAt: &jwt.NumericDate{Time: expirationTime},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// @Summary 用户登录
// @Description 用户通过手机号和密码登录
// @Accept  json
// @Produce  json
// @Param   phoneNumber   body    string  true   "Phone Number"
// @Param   password      body    string  true   "Password"
// @Success 200 {object} map[string]interface{}
// @Router /api/login [post]
func LoginHandler(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, LoginResponse{StatusCode: http.StatusBadRequest})
		return
	}

	db, exists := c.Get("db")
	if !exists {
		c.JSON(http.StatusInternalServerError, LoginResponse{StatusCode: http.StatusInternalServerError})
		return
	}
	database := db.(*gorm.DB)

	var user models.User // 使用 models.User
	if err := database.Where("phone_number = ?", request.PhoneNum).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, LoginResponse{StatusCode: http.StatusUnauthorized})
		return
	}

	token, err := GenerateJWT(request.PhoneNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, LoginResponse{StatusCode: http.StatusInternalServerError})
		return
	}

	c.JSON(http.StatusOK, LoginResponse{
		StatusCode: http.StatusOK,
		Data: &LoginData{
			Token: token,
		},
	})
}
