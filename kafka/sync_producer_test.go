package kafka

import (
	"context"
	"errors"
	"testing"

	"github.com/Shopify/sarama"
	"github.com/Shopify/sarama/mocks"
	"github.com/stretchr/testify/assert"
)

func TestSyncProducerPublish(t *testing.T) {
	type (
		producerContract struct {
			err error
		}

		input struct {
			producer producerContract
			topic    string
		}
	)

	tests := map[string]struct {
		input   input
		wantErr bool
	}{
		"returns an error if a producer returns an error": {
			input: input{
				producer: producerContract{err: errors.New("any")},
			},
			wantErr: true,
		},
		"returns successful result and applies options": {
			input: input{
				producer: producerContract{err: nil},
				topic:    "any-topic",
			},
			wantErr: false,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			prod := mocks.NewSyncProducer(t, nil)
			if tc.input.producer.err == nil {
				prod.ExpectSendMessageAndSucceed()
			} else {
				prod.ExpectSendMessageAndFail(tc.input.producer.err)
			}
			syncProd := SyncProducer{SyncProducer: prod}

			_, _, err := syncProd.Publish(context.Background(), &sarama.ProducerMessage{})
			assert.Equal(t, tc.wantErr, err != nil, "errors do not match, err=%s", err)
		})
	}
}

func TestSyncProducerIsHealthy(t *testing.T) {
	t.Run("returns false for empty sync producer", func(t *testing.T) {
		var prod SyncProducer
		assert.False(t, prod.IsHealthy())
	})
}

func TestSyncProducerClose(t *testing.T) {
	doneCtx, cancel := context.WithTimeout(context.Background(), -1)
	defer cancel()

	tests := map[string]struct {
		ctx     context.Context
		wantErr bool
	}{
		"closes sync producer successfully": {
			ctx:     context.Background(),
			wantErr: false,
		},
		"returns an error for done context": {
			ctx:     doneCtx,
			wantErr: true,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			prod := mocks.NewSyncProducer(t, nil)
			syncProd := SyncProducer{SyncProducer: prod}

			err := syncProd.Close(tc.ctx)
			assert.Equal(t, tc.wantErr, err != nil, "errors do not match, err=%s", err)
		})
	}
}
