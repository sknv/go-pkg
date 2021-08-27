package cascade

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/sknv/go-pkg/cascade/mock"
)

//go:generate mockgen -source=cascade.go -destination=mock/cascade.go -package=mock

func TestGet(t *testing.T) {
	bottomErr := errors.New("bottom")

	type storageContract struct {
		result interface{}
		err    error
	}

	tests := map[string]struct {
		storages []storageContract
		want     error
	}{
		"gets a record from the bottom layer successfully": {
			storages: []storageContract{
				{result: nil, err: errors.New("any")},
				{result: struct{}{}, err: nil},
			},
			want: nil,
		},
		"gets a record from the top layer successfully": {
			storages: []storageContract{
				{result: struct{}{}, err: nil},
				{result: nil, err: bottomErr},
			},
			want: nil,
		},
		"returns an `ErrNotFound`": {
			storages: []storageContract{
				{result: nil, err: nil},
				{result: nil, err: nil},
			},
			want: ErrNotFound,
		},
		"returns an error for the bottom layer": {
			storages: []storageContract{
				{result: nil, err: nil},
				{result: nil, err: bottomErr},
			},
			want: bottomErr,
		},
	}

	ctrl := gomock.NewController(t)

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var storages []Storage
			for _, st := range tc.storages {
				storage := mock.NewMockStorage(ctrl)
				storage.EXPECT().Get(gomock.Any(), gomock.Any()).
					Return(st.result, st.err).
					AnyTimes()
				storage.EXPECT().Set(gomock.Any(), gomock.Any()).AnyTimes()
				storages = append(storages, storage)
			}
			cascade := NewCascade(storages...)

			_, err := cascade.Get(context.Background(), struct{}{})
			assert.ErrorIs(t, err, tc.want, "errors do not match")
		})
	}
}

//nolint:dupl // false positive
func TestSet(t *testing.T) {
	bottomErr, topErr := errors.New("bottom"), errors.New("top")

	type storageContract struct {
		err error
	}

	tests := map[string]struct {
		storages []storageContract
		want     error
	}{
		"sets a record successfully": {
			storages: []storageContract{
				{err: nil},
				{err: nil},
			},
			want: nil,
		},
		"returns an error for the bottom layer": {
			storages: []storageContract{
				{err: topErr},
				{err: bottomErr},
			},
			want: bottomErr,
		},
		"returns an error for the top layer": {
			storages: []storageContract{
				{err: topErr},
				{err: nil},
			},
			want: topErr,
		},
	}

	ctrl := gomock.NewController(t)

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var storages []Storage
			for _, st := range tc.storages {
				storage := mock.NewMockStorage(ctrl)
				storage.EXPECT().Set(gomock.Any(), gomock.Any()).
					Return(st.err).
					AnyTimes()
				storages = append(storages, storage)
			}
			cascade := NewCascade(storages...)

			err := cascade.Set(context.Background(), struct{}{})
			assert.ErrorIs(t, err, tc.want, "errors do not match")
		})
	}
}

//nolint:dupl // false positive
func TestDelete(t *testing.T) {
	bottomErr, topErr := errors.New("bottom"), errors.New("top")

	type storageContract struct {
		err error
	}

	tests := map[string]struct {
		storages []storageContract
		want     error
	}{
		"deletes a record successfully": {
			storages: []storageContract{
				{err: nil},
				{err: nil},
			},
			want: nil,
		},
		"returns an error for the bottom layer": {
			storages: []storageContract{
				{err: topErr},
				{err: bottomErr},
			},
			want: bottomErr,
		},
		"returns an error for the top layer": {
			storages: []storageContract{
				{err: topErr},
				{err: nil},
			},
			want: topErr,
		},
	}

	ctrl := gomock.NewController(t)

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var storages []Storage
			for _, st := range tc.storages {
				storage := mock.NewMockStorage(ctrl)
				storage.EXPECT().Delete(gomock.Any(), gomock.Any()).
					Return(st.err).
					AnyTimes()
				storages = append(storages, storage)
			}
			cascade := NewCascade(storages...)

			err := cascade.Delete(context.Background(), struct{}{})
			assert.ErrorIs(t, err, tc.want, "errors do not match")
		})
	}
}
