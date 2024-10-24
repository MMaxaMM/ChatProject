package slogx

import (
	"log/slog"
)

// Атрибут "ошибка" для slog
func Error(err error) slog.Attr {
	attr := slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}

	return attr
}
