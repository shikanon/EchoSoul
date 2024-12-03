package models

import (
	"time"
)

// Episode 表示播客的一集节目
type Episode struct {
	ID          uint   `gorm:"primaryKey"`
	PodcastID   uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	FileURL     string `gorm:"not null"`
	CoverImage  string  // 保存节目的封面图片链接
	FavoriteCount int    `gorm:"default:0"` // 收藏数
	Duration    int // 节目的时长（以秒为单位）
	CreatedAt   time.Time
}

func (Episode) TableName() string {
	return "episodes"
}
