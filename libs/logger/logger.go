package logger

import (
	"context"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"log"
	"os"
	"syscall"
)

type ctxKey string

const (
	// LogContextKey ...
	LogContextKey = "logger"

	logTagKey = "tag"
)

// New ...
func New() *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	err := syscall.Dup2(int(os.Stdout.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		log.Println("error while copying from stderr to stdout while initializing logging")
	}

	logger := zerolog.New(os.Stdout).With().
		Timestamp().
		Caller().
		Logger()
	return &logger
}

// NewContextWithLogger creates new context with logger
func NewContextWithLogger() context.Context {
	l := New()
	ctx := context.TODO()
	return context.WithValue(ctx, ctxKey(LogContextKey), l)
}

// WithLogger add logger to context and returns a new copy
func WithLogger(ctx context.Context, l *zerolog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey(LogContextKey), l)
}

// FromContext returns existing logger from context, otherwise returns new logger
func FromContext(ctx context.Context) *zerolog.Logger {
	if l, ok := ctx.Value(ctxKey(LogContextKey)).(*zerolog.Logger); ok {
		return l
	}

	return New()
}

// FromContextWithTag returns the existing logger from context with log tag
func FromContextWithTag(ctx context.Context, tag string) *zerolog.Logger {
	l := FromContext(ctx).With().Str(logTagKey, tag).Logger()
	return &l
}
