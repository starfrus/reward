package domain

import (
	"errors"
	"fmt"
	"regexp"
	core_error "rewardapp/internal/core/errors"
	"time"
)

type User struct {
	ID                    int
	Version               int
	Login                 string
	Password              string
	CompletedTargetsCount int
	CreatedAt             time.Time
}

func NewUser(
	id int,
	version int,
	login string,
	password string,
	completedTargetsCount int,
	createdAt time.Time,

) User {
	return User{
		ID:                    id,
		Version:               version,
		Login:                 login,
		Password:              password,
		CompletedTargetsCount: completedTargetsCount,
		CreatedAt:             createdAt,
	}
}

type UserResponse struct {
	ID                    int       `json:"id"`
	Login                 string    `json:"login"`
	CompletedTargetsCount int       `json:"completed_targets_count"`
	CreatedAt             time.Time `json:"created_at"`
}

func (u *User) ValidateLogin() error {
	lengthLogin := len([]rune(u.Login))
	if lengthLogin < 2 && lengthLogin > 30 {
		return fmt.Errorf("invalid `Login` len %d %w",
			lengthLogin,
			core_error.ErrInvaligArgument,
		)
	}

	re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]{2,29}$`)

	if !re.MatchString(u.Login) {
		return fmt.Errorf("invalid `Login` format: %w", core_error.ErrInvaligArgument)
	}
	return nil
}

var (
	reUpper   = regexp.MustCompile(`[A-Z]`)
	reLower   = regexp.MustCompile(`[a-z]`)
	reDigit   = regexp.MustCompile(`[0-9]`)
	reSpecial = regexp.MustCompile(`[!@#$%^&*]`)
)

// ValidatePassword возвращает список ошибок (пустой слайс — пароль валиден)
func (u *User) ValidatePassword() error {
	var errs []error
	p := u.Password
	if n := len([]rune(p)); n < 8 || n > 16 {
		errs = append(errs, fmt.Errorf("длина пароля должна быть от 8 до 16 символов (сейчас %d)", n))
	}
	if !reUpper.MatchString(p) {
		errs = append(errs, errors.New("пароль должен содержать хотя бы одну заглавную латинскую (A-Z) букву"))
	}
	if !reLower.MatchString(p) {
		errs = append(errs, errors.New("пароль должен содержать хотя бы одну строчную латинскую (a-z) букву"))
	}
	if !reDigit.MatchString(p) {
		errs = append(errs, errors.New("пароль должен содержать хотя бы одну цифру (0-9)"))
	}
	if !reSpecial.MatchString(p) {
		errs = append(errs, errors.New("пароль должен содержать хотя бы один спецсимвол: !@#$%^&*"))
	}

	return errors.Join(errs...)
}
