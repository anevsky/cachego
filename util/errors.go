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
  ErrorBadRequest = CacheError{"Bad Request ", 400}
  ErrorKeyNotFound = CacheError{"Key Not Found", 404}
)
