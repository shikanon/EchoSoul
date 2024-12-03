package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/shikanon/EchoSoul/models"
	"github.com/shikanon/EchoSoul/utils"
)

func TestLoginHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	db := utils.MockDB()
	router.POST("/api/login", func(c *gin.Context) {
		c.Set("db", db)
		LoginHandler(c)
	})

	// 准备好 mock 数据库和插入一个用户
	db.Create(&models.User{
		PhoneNumber: "13800138000",
		Password:    "password",
	})

	// 创建一个有效的请求体
	body := `{"phoneNum": "13800138000"}`
	req, _ := http.NewRequest("POST", "/api/login", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// 录制响应
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 解析响应
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")

	// 创建一个无效的请求体
	body = `{"phoneNum": "invalid_phone_number"}`
	req, _ = http.NewRequest("POST", "/api/login", strings.NewReader(body))
	req.Header.Add("Content-Type", "application/json")

	// 录制新的响应
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// 解析响应
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
