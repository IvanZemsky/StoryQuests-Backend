package repository

import (
	"context"
	"time"
)

const defaultRequestTimeout = time.Second * 30

func NewRequestTimeoutContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultRequestTimeout)
	return ctx, cancel
}

// time.Second * time.Duration(timeout)
func NewCustomRequestTimeoutContext(timeout int) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(timeout))
	return ctx, cancel
}
