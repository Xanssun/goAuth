package handlers

import (
	"encoding/json"
	"goauth/internal/schemas"
	"goauth/internal/services"
	"goauth/pkg/logging"
	"net/http"
)

// AuthHandler обрабатывает запросы аутентификации
type AuthHandler struct {
	Service *services.AuthService
}

// NewAuthHandler создает новый экземпляр AuthHandler
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{Service: service}
}

// Registration обрабатывает запросы на регистрацию пользователя
func (h *AuthHandler) Registration(w http.ResponseWriter, r *http.Request) {
	var req schemas.RegistrationRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		logging.Log.Errorf("Ошибка при декодировании тела запроса: %v", err)
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}

	user, err := h.Service.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		logging.Log.Errorf("Ошибка при создании пользователя: %v", err)
		http.Error(w, "Ошибка при создании пользователя", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
