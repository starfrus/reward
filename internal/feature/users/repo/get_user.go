package repo

import (
	"context"
	"errors"
	"fmt"
	"rewardapp/internal/core/domain"
	core_error "rewardapp/internal/core/errors"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepo) GetUser(
	ctx context.Context,
	id int,
) (domain.UserResponse, error) {
	query := `
	SELECT id, login, completed_targets_count, created_at
	FROM users
	WHERE id=$1;
	`

	row := r.db.QueryRow(ctx, query, id)

	var userModel domain.UserResponse

	err := row.Scan(
		&userModel.ID,
		&userModel.Login,
		&userModel.CompletedTargetsCount,
		&userModel.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.UserResponse{}, fmt.Errorf("user with id=%d: %w", id, core_error.ErrNotFound)
		}
		return domain.UserResponse{}, fmt.Errorf("scan error: %w", err)
	}
	return userModel, nil // userModel уже domain.UserResponse

}
