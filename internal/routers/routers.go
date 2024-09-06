// internal/routers/routers.go
package routers

import (
	"goauth/internal/handlers"
	"goauth/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func SetupRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Создание сервисов и обработчиков
	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	r.Post("/users/register", authHandler.Registration)

	return r
}
