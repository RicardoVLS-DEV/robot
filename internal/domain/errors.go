package domain

import "errors"

var (
	ErrInvalidName = errors.New("invalid name")
	ErrInvalidEmail = errors.New("invalid email")
	ErrEmpty = errors.New("field empty")
	ErrInvalidFile = errors.New("invalid file")
	ErrClosingFile = errors.New("error closing the file")
	ErrUnsuccessIntegrante = errors.New("something went wrong by creating a integrante")
)