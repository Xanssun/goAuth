package config

import (
	"goauth/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "host=postgres_db user=app dbname=auth_db password=123qwe port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logging.Log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	logging.Log.Info("Успешное подключение к базе данных")
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		logging.Log.Errorf("Ошибка при получении sqlDB: %v", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		logging.Log.Errorf("Ошибка при закрытии базы данных: %v", err)
	} else {
		logging.Log.Info("Соединение с базой данных закрыто")
	}
}
