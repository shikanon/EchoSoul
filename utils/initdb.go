package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/shikanon/EchoSoul/models"
)

func InitDB() *gorm.DB {
	// 从环境变量中读取数据库连接信息
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 进行自动迁移
	db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{}, &models.UserSubscription{})

	return db
}

func MockDB() *gorm.DB {
	dbUser := "podcast_user"
	dbPassword := "your_password123"
	dbHost := "localhost"
	dbPort := "13306"
	dbName := "echosouldb"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)
	fmt.Printf("dsn: %s\n", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 重置数据库表
	err = db.Migrator().DropTable(&models.User{}, &models.Podcast{}, &models.Episode{}, &models.UserSubscription{})
	if err != nil {
		panic("failed to drop table")
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&models.User{}, &models.Podcast{}, &models.Episode{}, &models.UserSubscription{})
	if err != nil {
		panic("failed to migrate database")
	}

	return db
}
