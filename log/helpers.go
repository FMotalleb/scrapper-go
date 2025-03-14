// Package log contains utilities for logging system
package log

import (
	"log/slog"
	"os"
)

func SetupLogger(level string) error {
	var lvl slog.Level
	if err := lvl.UnmarshalText([]byte(level)); err != nil {
		return err
	}

	opts := &slog.HandlerOptions{}
	opts.Level = lvl
	newRoot := slog.New(slog.NewTextHandler(os.Stderr, opts))
	slog.SetDefault(newRoot)
	return nil
}
