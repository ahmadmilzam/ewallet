package store

import "errors"

var ErrNoResult = errors.New("no results")
var ErrUniqueViolation = errors.New("violate unique constraint")
var ErrDBUnavailable = errors.New("db unavailable")
