package service

import (
	"context"
	"fmt"
	"rewardapp/internal/core/domain"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Create(ctx context.Context, user domain.User) (domain.User, error) {
	// 1. Проверка логина
	if err := user.ValidateLogin(); err != nil {
		return domain.User{}, err
	}

	if err := user.ValidatePassword(); err != nil {
		return domain.User{}, err
	}

	// 2. Хэшируем
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return domain.User{}, fmt.Errorf("hash password: %w", err)
	}

	//3. Подменяем пароль
	user.Password = string(hash)

	// 3. Передаем в репозиторий. Он сохранит и вернет полного юзера из БД.
	user, err = s.repo.Create(ctx, user)
	if err != nil {
		return domain.User{}, err
	}
	// 4. Возвращаем результат
	return user, nil
}
