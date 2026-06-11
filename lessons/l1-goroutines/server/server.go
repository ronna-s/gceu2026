package main

import (
	"context"
	"io"
	"log/slog"
	"net"
	"os"
	"sync"
	"time"
)

type contextValuesKey struct{}

func contextWithValues(ctx context.Context, key, value any) context.Context {
	v, ok := ctx.Value(contextValuesKey{}).(map[any]any)
	if !ok {
		return context.WithValue(ctx, contextValuesKey{}, map[any]any{key: value})
	}

	// Copy-on-write so child contexts do not mutate parent values.
	next := make(map[any]any, len(v)+1)
	for k, val := range v {
		next[k] = val
	}
	next[key] = value
	return context.WithValue(ctx, contextValuesKey{}, next)
}

func loggerFromContext(ctx context.Context) *slog.Logger {
	values, ok := ctx.Value(contextValuesKey{}).(map[any]any)
	if !ok || len(values) == 0 {
		return logger
	}

	// Convert map to alternating key-value pairs for slog.With()
	var attrs []any
	for k, v := range values {
		attrs = append(attrs, k, v)
	}
	return logger.With(attrs...)
}

func main() {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		slog.Error("failed to create a listener", slog.Any("error", err))
	}

	ctx := contextWithValues(context.Background(), "address", l.Addr().String())
	loggerFromContext(ctx).Info("listening")
	go func() {
		<-time.After(20 * time.Second)
		l.Close()
	}()
	serve(ctx, l)
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level: slog.LevelInfo,
}))

func handle(ctx context.Context, conn net.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	if deadline, ok := ctx.Deadline(); ok {
		_ = conn.SetReadDeadline(deadline)
		defer conn.SetReadDeadline(time.Time{}) // optional: clear deadline
	}
	b, err := io.ReadAll(conn)
	if err != nil {
		return err
	}
	loggerFromContext(ctx).InfoContext(ctx, "handle", slog.Any("bytes received", len(b)))
	return nil
}

func serve(ctx context.Context, l net.Listener) error {
	var (
		wg sync.WaitGroup
	)
	defer wg.Wait()
	for id := 0; ; id++ {
		conn, err := l.Accept()
		if err != nil {
			loggerFromContext(ctx).ErrorContext(ctx, "failed to accept connection", slog.Any("err", err))
			return err
		}
		childCtx := contextWithValues(ctx, "connection", id)
		wg.Go(func() {
			if err := handle(childCtx, conn); err != nil {
				loggerFromContext(childCtx).ErrorContext(childCtx, "failed to handle connection", slog.Any("error", err))
			}
		})
	}
}
