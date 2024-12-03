package models

import "time"

// UserSubscription 表示用户订阅关系
type UserSubscription struct {
	UserID       uint `gorm:"primaryKey"`
	PodcastID    uint `gorm:"primaryKey"`
	SubscribedAt time.Time
}

func (UserSubscription) TableName() string {
	return "user_subscriptions"
}
