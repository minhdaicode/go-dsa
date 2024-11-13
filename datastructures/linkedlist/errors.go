package linkedlist

import "errors"

var (
	ErrEmpty         = errors.New("list is empty")
	ErrPosOutOfRange = errors.New("position out of range")
)
