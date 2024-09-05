package config

import (
	"database/sql"
	"goauth/pkg/logging"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func pingDB(dsn string) bool {
	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		logging.Log.Warnf("Ошибка при создании подключения: %v", err)
		return false
	}
	defer sqlDB.Close()

	err = sqlDB.Ping()
	if err != nil {
		logging.Log.Warnf("База данных не отвечает: %v", err)
		return false
	}
	return true
}

func ConnectDB() {
	dsn := "host=postgres_db user=app dbname=auth_db password=123qwe port=5432 sslmode=disable"

	// Попытка пинга базы данных
	for {
		if pingDB(dsn) {
			logging.Log.Info("База данных готова, подключение...")
			break
		} else {
			logging.Log.Warn("База данных не готова, повторная попытка через 10 секунд...")
			time.Sleep(10 * time.Second)
		}
	}

	// Полное подключение через GORM после успешного пинга
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
