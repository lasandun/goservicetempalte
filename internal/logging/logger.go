package logging

import (
	"log/slog"
	"os"
)

func InitJSONLogger(level slog.Leveler) {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})
	slog.SetDefault(slog.New(handler))
}
