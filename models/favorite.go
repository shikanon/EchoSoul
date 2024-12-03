package models

import "time"

// UserFavorite 表示用户收藏的节目
type UserFavorite struct {
	UserID     uint `gorm:"primaryKey"`
	EpisodeID  uint `gorm:"primaryKey"`
	FavoriteAt time.Time
}

func (UserFavorite) TableName() string {
	return "user_favorites"
}
