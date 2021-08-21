package log

import (
	"context"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	type input struct {
		logger logrus.FieldLogger
		fields logrus.Fields
	}

	tests := map[string]struct {
		input          input
		wantNullLogger bool
	}{
		"returns an empty logger if one was not provided": {
			input:          input{},
			wantNullLogger: true,
		},
		"returns the provided logger if exists": {
			input:          input{logger: logrus.StandardLogger()},
			wantNullLogger: false,
		},
		"returns the logger with provided fields if exists": {
			input:          input{logger: _nullLogger, fields: logrus.Fields{"any-field": "any-value"}},
			wantNullLogger: false,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := ToContext(context.Background(), tc.input.logger)
			AddFields(ctx, tc.input.fields)
			ctxLogger := Extract(ctx)
			assert.Equal(t, tc.wantNullLogger, ctxLogger == _nullLogger)
		})
	}
}
