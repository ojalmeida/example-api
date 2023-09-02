package database

import "errors"

var (
	ErrClientNotFound error = errors.New("client not found")
)