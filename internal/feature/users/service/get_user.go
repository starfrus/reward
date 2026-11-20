package service

import (
	"context"
	"errors"
	"fmt"
	"rewardapp/internal/core/domain"
	core_error "rewardapp/internal/core/errors"
)

func (s *UserService) GetUser(
	ctx context.Context,
	id int,
) (domain.UserResponse, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, core_error.ErrNotFound) {
			return domain.UserResponse{}, err
		}
		return domain.UserResponse{}, fmt.Errorf("get user from repo: %w", err)
	}
	return user, nil
}
