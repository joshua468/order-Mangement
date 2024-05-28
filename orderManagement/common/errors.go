package common

import "errors"

var (
	ErrNoItems = errors.New("items must have atleast one item")
)
