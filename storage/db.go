package storage

import (
	"fmt"

	models "github.com/ntwaliheritier/go_blog_app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}

func NewConnection(config Config) error {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", 
		config.Host, config.Username, config.Password, config.DBName, config.Port, config.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Post{})
	return nil
}