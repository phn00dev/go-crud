package dbconfig

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/phn00dev/go-crud/pkg/config"

)

type DbConfig struct {
	config *config.Config
}

func NewDbConnection(config *config.Config) *DbConfig {
	return &DbConfig{
		config: config,
	}
}

func (dbConfig DbConfig) GetDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbConfig.config.DbConfig.DbHost,
		dbConfig.config.DbConfig.DbUser,
		dbConfig.config.DbConfig.DbPassword,
		dbConfig.config.DbConfig.DbName,
		dbConfig.config.DbConfig.DbPort,
		dbConfig.config.DbConfig.DbSslMode,
		dbConfig.config.DbConfig.DbTimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("database connection error: %v", err)
		return nil, err
	}
	return db, nil
}
