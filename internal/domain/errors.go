package domain

import (
	"errors"
	"fmt"
)

var (
	ErrEmpty = errors.New("field empty")
	ErrNotEnoughMembers = errors.New("not enough members to add to team")
)

type RobotError struct {
	Op	string
	Field string
	Err error
}

func (e *RobotError) Error() string {
	return	fmt.Sprintf("[%s:%s] %s", e.Field, e.Op, e.Err) 
}

func (e *RobotError) Unwrap() error {
	return e.Err
}

func NewRobotErr(op, field string, err error) *RobotError {
	if err == nil {
		return nil
	}

	return &RobotError{
		Op: op,
		Field: field,
		Err: err,
	}
}