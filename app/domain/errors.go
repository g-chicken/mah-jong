package domain

import (
	"errors"
	"fmt"
)

var (
	errNilPlayer             = errors.New("player is nil")
	errNilHand               = errors.New("hand is nil")
	errNilPlayerScore        = errors.New("player score is nil")
	errNilHalfRoundGameScore = errors.New("half round game score is nil")
)

type InvalidArgumentError struct {
	msg string
}

func NewInvalidArgumentError(msg string) InvalidArgumentError {
	return InvalidArgumentError{msg: msg}
}

func (e InvalidArgumentError) Error() string {
	return fmt.Sprintf("Invalid Argument: %s", e.msg)
}

type NotFoundError struct {
	msg string
}

func NewNotFoundError(msg string) NotFoundError {
	return NotFoundError{msg: msg}
}

func (e NotFoundError) Error() string {
	return fmt.Sprintf("Not Found: %s", e.msg)
}

type ConflictError struct {
	msg string
}

func NewConflictError(msg string) ConflictError {
	return ConflictError{msg: msg}
}

func (e ConflictError) Error() string {
	return fmt.Sprintf("Conflict: %s", e.msg)
}

type IllegalForeignKeyConstraintError struct {
	msg string
}

func NewIllegalForeignKeyConstraintError(msg string) IllegalForeignKeyConstraintError {
	return IllegalForeignKeyConstraintError{msg: msg}
}

func (e IllegalForeignKeyConstraintError) Error() string {
	return fmt.Sprintf("Illegal Foreign Key Constraint: %s", e.msg)
}
