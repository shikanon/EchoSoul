package models

import (
	"gorm.io/gorm"
)

// Podcast 播客
type Podcast struct {
	gorm.Model
	ID                uint      `gorm:"primaryKey"` // 播客唯一标识符
	Title             string    `gorm:"not null"`   // 播客名称
	Description       string    // 播客简介
	Tags              []string  `gorm:"type:text[]"` // 播客的标签列表
	ImageURL          string    // 播客的封面图片链接
	Display           int       // 播客的播放次数
	Focus             int       // 播客的订阅数
	Score             float64   // 播客的评分
	EpisodeCount      int       // 播客包含的节目数量
	Subscribed        bool      // 是否已订阅
	CreatedBy         uint      `gorm:"not null"` // 上传者的用户ID或官方ID
	UploaderType      string    `gorm:"not null"` // "user" 或 "official"
	CoverImage        string    // 保存播客的封面图片链接
	SubscriptionCount int       `gorm:"default:0"` // 订阅数
	Episodes          []Episode `gorm:"foreignKey:PodcastID"`
}

func (Podcast) TableName() string {
	return "podcasts"
}
