package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swag 文档文件
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware

	"github.com/shikanon/EchoSoul/handlers"
	_ "github.com/shikanon/EchoSoul/swagger_docs"
	"github.com/shikanon/EchoSoul/utils"
)

// @title EchoSoul API文档
// @version 1.0
// @description EchoSoul 播客项目的 API 文档
// @host localhost:8080
// @BasePath /api
func main() {
	db := utils.InitDB()

	r := gin.Default()

	// 将数据库实例传递给路由或中间件
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// API 路由
	api := r.Group("/api")
	{
		// 用户登录路由
		api.POST("/login", handlers.LoginHandler)
		// 用户收藏播客节目
		api.POST("/episode/subscribe", handlers.SubscribeEpisodeHandler)
		// 取消收藏播客路由
		api.POST("/episode/unsubscribe", handlers.UnsubscribeEpisodeHandler)
		// 查询播客列表路由
		api.POST("/podcast/query", handlers.QueryPodcastHandler)
		// 创建播客路由
		api.POST("/podcast/create", handlers.CreatePodcastHandler)
	}

	// Swagger路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 启动服务
	err := r.Run()
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
}
