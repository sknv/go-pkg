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

type Storage[Q any, R comparable] interface {
	// Get should return (zero value, nil) pair if a record is not found.
	Get(ctx context.Context, getBy Q) (R, error)
	Set(ctx context.Context, setBy Q, record R) error
	Delete(ctx context.Context, deleteBy Q) error
}

// Cascade manages multiple storages allowing to get and setReversed records in cascade manner.
type Cascade[Q any, R comparable] struct {
	storages []Storage[Q, R]
}

// NewCascade returns a new instance.
func NewCascade[Q any, R comparable](storages ...Storage[Q, R]) *Cascade[Q, R] {
	return &Cascade[Q, R]{
		storages: storages,
	}
}

// Get fetches a record from the cascade, setting the value for the missing storages.
func (c *Cascade[Q, R]) Get(ctx context.Context, getBy Q) (R, error) {
	logger := log.Extract(ctx)

	var (
		result R
		err    error
	)
	foundIndex := -1
	for i, st := range c.storages {
		result, err = st.Get(ctx, getBy)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"error": err,
				"index": i,
			}).Error("get record from storage")
			continue
		}

		var zeroR R
		if result == zeroR {
			continue
		}

		foundIndex = i
		break
	}

	if foundIndex < 0 {
		if err == nil {
			err = ErrNotFound
		}

		var zeroR R
		return zeroR, err
	}

	// Sync storages
	if err = c.setReversed(ctx, getBy, result, foundIndex-1); err != nil {
		logger.WithError(err).Error("sync storages")
	}

	return result, nil
}

// Set sets the record for all the storages.
func (c *Cascade[Q, R]) Set(ctx context.Context, setBy Q, record R) error {
	return c.setReversed(ctx, setBy, record, len(c.storages)-1)
}

// Delete deletes the record from all the storages.
func (c *Cascade[Q, R]) Delete(ctx context.Context, deleteBy Q) error {
	for i := len(c.storages) - 1; i >= 0; i-- { // reversed order
		st := c.storages[i]
		if err := st.Delete(ctx, deleteBy); err != nil {
			return fmt.Errorf("delete record from storage with index %d: %w", i, err)
		}
	}
	return nil
}

func (c *Cascade[Q, R]) setReversed(ctx context.Context, setBy Q, record R, upperIndex int) error {
	for i := upperIndex; i >= 0; i-- { // reversed order
		st := c.storages[i]
		if err := st.Set(ctx, setBy, record); err != nil {
			return fmt.Errorf("set record for storage with index %d: %w", i, err)
		}
	}
	return nil
}
