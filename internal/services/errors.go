package services

import "errors"

var (
	InternalServerError error = errors.New("internal server error")
	AlreadyExistsError  error = errors.New("already exists error")
)
