package util

import (
	"fmt"
	"time"
)

type CacheError struct {
	What string
	Code int
}

type LogError struct {
	When time.Time
	What string
	Code int
}

func (e CacheError) Error() string {
	return fmt.Sprintf("%s, code: %d",
		e.What, e.Code)
}

func (e LogError) Error() string {
	return fmt.Sprintf("at %v, %s, code: %d",
		e.When, e.What, e.Code)
}

var (
	ErrorWrongType         = CacheError{"Wrong type", 999}
	ErrorIndexOutOfBounds  = CacheError{"Index out of Bounds", 998}
	ErrorInvalidTTLValue   = CacheError{"Invalid ttl value", 997}
	ErrorResponseOrBodyNil = CacheError{"Response or body nil", 996}
	ErrorBadRequest        = CacheError{"Bad request", 400}
	ErrorKeyNotFound       = CacheError{"Key not found", 404}
	ErrorDictKeyNotFound   = CacheError{"Key not found in dictionary", 404}
)
