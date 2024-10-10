package pwlogger

import (
	"context"
	"log/slog"
	"strings"
)

// StreamType represents the type of output stream
type StreamType int

const (
	StdoutStream StreamType = iota
	StderrStream
)

func (st StreamType) String() string {
	switch st {
	case StdoutStream:
		return "stdout"
	case StderrStream:
		return "stderr"
	default:
		return "unknown"
	}
}

func (st StreamType) LogValue() slog.Value {
	return slog.StringValue(st.String())
}

// SlogWriter is a custom type that implements io.Writer
type SlogWriter struct {
	logger   *slog.Logger
	stream   StreamType
	cmdAttrs []slog.Attr
}

// Write implements the io.Writer interface
func (sw *SlogWriter) Write(p []byte) (n int, err error) {
	message := strings.TrimSpace(string(p))
	attrs := append(sw.cmdAttrs,
		slog.String("stream", sw.stream.String()),
	)
	sw.logger.LogAttrs(context.Background(), slog.LevelInfo, message, attrs...)
	return len(p), nil
}

// NewSlogWriter creates a new SlogWriter with the specified stream type and command attributes
func NewSlogWriter(logger *slog.Logger, stream StreamType, cmdAttrs ...slog.Attr) *SlogWriter {
	return &SlogWriter{
		logger:   logger,
		stream:   stream,
		cmdAttrs: cmdAttrs,
	}
}

func ErrAttr(err error) slog.Attr {
	return slog.Any("error", err)
}
