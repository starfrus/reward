package repo

import (
	"context"
	"fmt"
	"rewardapp/internal/core/domain"
	core_error "rewardapp/internal/core/errors"
)

func (r *UserRepo) Create(ctx context.Context, user domain.User) (domain.User, error) {

	query := `INSERT INTO users (login, password_hash)
	VALUES($1, $2)
	RETURNING id, version, login, password_hash, completed_targets_count, created_at;
	`

	row := r.db.QueryRow(ctx, query, user.Login, user.Password)

	var userModel domain.User

	err := row.Scan(
		&userModel.ID,
		&userModel.Version,
		&userModel.Login,
		&userModel.Password,
		&userModel.CompletedTargetsCount,
		&userModel.CreatedAt,
	)
	if err != nil {
		if isUniqueViolation(err) {
			return domain.User{}, core_error.ErrUserAlreadyExists
		}
		return domain.User{}, fmt.Errorf("create user: %w", err)
	}

	userDomain := domain.NewUser(
		userModel.ID,
		userModel.Version,
		userModel.Login,
		userModel.Password,
		userModel.CompletedTargetsCount,
		userModel.CreatedAt,
	)

	return userDomain, nil
}
