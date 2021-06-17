/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabenc

import (
	"io"
	"time"

	zaplogfmt "github.com/sykesm/zap-logfmt"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// FormatEncoder is a zapcore.Encoder that formats log records according to a
// go-logging based format specifier.
type FormatEncoder struct {
	zapcore.Encoder
	formatters []Formatter
	pool       buffer.Pool
}

// Formatter is used to format and write data from a zap log entry.
type Formatter interface {
	Format(w io.Writer, entry zapcore.Entry, fields []zapcore.Field)
}

func NewFormatEncoder(formatters ...Formatter) *FormatEncoder {
	return &FormatEncoder{
		Encoder: zaplogfmt.NewEncoder(zapcore.EncoderConfig{
			MessageKey:    "",
			LevelKey:      "",
			TimeKey:       "",
			NameKey:       "",
			CallerKey:     "",
			StacktraceKey: "",
			LineEnding:    "\n",
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02T15:04:05.999Z07:00"))
			},
			EncodeDuration: zapcore.StringDurationEncoder},
		),
		formatters: formatters,
		pool:       buffer.NewPool(),
	}
}

func (f *FormatEncoder) Clone() zapcore.Encoder {
	return &FormatEncoder{
		Encoder:    f.Encoder.Clone(),
		formatters: f.formatters,
		pool:       f.pool,
	}
}

func (f *FormatEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	line := f.pool.Get()
	for _, formatter := range f.formatters {
		formatter.Format(line, entry, fields)
	}

	encodedFields, err := f.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	if line.Len() > 0 && encodedFields.Len() != 1 {
		line.AppendString(" ")
	}
	line.AppendString(encodedFields.String())
	encodedFields.Free()

	return line, nil
}
