package slogx

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"

	"github.com/fatih/color"
)

// Атрибут "ошибка" для slog
func Error(err error) slog.Attr {
	attr := slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}

	return attr
}

type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := r.Time.Format("[15:04:05.000]")
	msg := color.HiCyanString(r.Message)

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	handler := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts),
		l:       log.New(out, "", 0),
	}

	return handler
}
