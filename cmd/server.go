package cmd

import (
	"goauth/config"
	"goauth/internal/models"
	"goauth/internal/routers"
	"goauth/pkg/logging"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func Run() {
	logging.InitLogger()

	// Подключение к базе данных через GORM
	config.ConnectDB()

	// Автоматическое создание таблицы для модели User
	config.DB.AutoMigrate(&models.User{})

	defer config.CloseDB()

	r := routers.SetupRouter()

	go func() {
		logging.Log.WithFields(logrus.Fields{
			"port": 3000,
		}).Info("Сервер запущен")
		if err := http.ListenAndServe(":3000", r); err != nil {
			logging.Log.Fatalf("Сбой сервера: %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	logging.Log.Info("Завершение работы сервера...")

	config.CloseDB()
	logging.Log.Info("Сервер корректно остановлен")
}
