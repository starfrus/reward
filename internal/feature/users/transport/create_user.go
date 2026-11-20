package transport

import (
	"encoding/json"
	"errors"
	"net/http"
	"rewardapp/internal/core/domain"
	core_error "rewardapp/internal/core/errors"
)

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req CreateRequestDTO

	// 1. Парсим JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid requst JSON", http.StatusBadRequest)
		return
	}

	user, err := h.s.Create(r.Context(), domain.User{
		Login:    req.Login,
		Password: req.Password,
	})

	if err != nil {
		switch {
		case errors.Is(err, core_error.ErrUserAlreadyExists):
			http.Error(w, "user already exists", http.StatusConflict)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	// 4. Возвращаем ответ
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user) // Вернется domain.User, у которого Password скрыт тегом `json:"-"`
}
