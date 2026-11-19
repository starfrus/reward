package core_error

import "errors"

var (
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrNotFound           = errors.New("not found")
	ErrInvaligArgument    = errors.New("invalid argument")
	ErrConflict           = errors.New("conflict")
	ErrNoRows             = errors.New("no rows")
	ErrViolatesForeignKey = errors.New("violates foreign key")
	ErrUnknown            = errors.New("unknown")
)
