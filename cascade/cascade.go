package cascade

import (
	"context"
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/sknv/go-pkg/log"
)

var (
	ErrNotFound = errors.New("record not found")
)

// Dependency contract

type Storage interface {
	// Get should return (nil, nil) pair if a record is not found.
	Get(ctx context.Context, record interface{}) (interface{}, error)
	Set(ctx context.Context, record interface{}) error
	Delete(ctx context.Context, record interface{}) error
}

// Cascade manages multiple storages allowing to get and setReversed records in cascade manner.
type Cascade struct {
	storages []Storage
}

// NewCascade returns a new instance.
func NewCascade(storages ...Storage) *Cascade {
	return &Cascade{
		storages: storages,
	}
}

// Get fetches a record from the cascade, setting the value for the missing storages.
func (c *Cascade) Get(ctx context.Context, record interface{}) (interface{}, error) {
	logger := log.Extract(ctx)

	var (
		result interface{}
		err    error
	)
	foundIndex := -1
	for i, st := range c.storages {
		result, err = st.Get(ctx, record)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"error": err,
				"index": i,
			}).Error("get record from storage")
			continue
		}

		if result == nil {
			continue
		}

		foundIndex = i
		break
	}

	if foundIndex < 0 {
		if err == nil {
			err = ErrNotFound
		}
		return nil, err
	}

	// Sync storages
	if err = c.setReversed(ctx, result, foundIndex-1); err != nil {
		logger.WithError(err).Error("sync storages")
	}

	return result, nil
}

// Set sets the record for all the storages.
func (c *Cascade) Set(ctx context.Context, record interface{}) error {
	return c.setReversed(ctx, record, len(c.storages)-1)
}

// Delete deletes the record from all the storages.
func (c *Cascade) Delete(ctx context.Context, record interface{}) error {
	for i := len(c.storages) - 1; i >= 0; i-- { // reversed order
		st := c.storages[i]
		if err := st.Delete(ctx, record); err != nil {
			return fmt.Errorf("delete record from storage with index %d: %w", i, err)
		}
	}
	return nil
}

func (c *Cascade) setReversed(ctx context.Context, record interface{}, upperIndex int) error {
	for i := upperIndex; i >= 0; i-- { // reversed order
		st := c.storages[i]
		if err := st.Set(ctx, record); err != nil {
			return fmt.Errorf("set record for storage with index %d: %w", i, err)
		}
	}
	return nil
}
