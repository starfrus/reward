package transport

import (
	"fmt"
	"net/http"
	core_error "rewardapp/internal/core/errors"
	"rewardapp/internal/feature/users/service"
	"strconv"
)

type UserHandler struct {
	s *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{s: s}
}

// DTO для входящего запроса (без ID, его генерирует БД)
type CreateRequestDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path values: %w",
			key,
			core_error.ErrInvaligArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"path value='%s' by key='%s' not a valid intenger: %v: %w",
			pathValue,
			key,
			err,
			core_error.ErrInvaligArgument,
		)
	}
	return val, nil
}
