package transport

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	core_error "rewardapp/internal/core/errors"
)

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userID, err := GetIntPathValue(r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := h.s.GetUser(r.Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, core_error.ErrNotFound):
			http.Error(w, "user not found", http.StatusNotFound)
		default:
			log.Printf("get user: %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
