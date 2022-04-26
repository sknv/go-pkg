package cascade

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

//go:generate moq -out mocks_test.go -fmt goimports . Storage

func TestCascade_Delete(t *testing.T) {
	type storageContract struct {
		err error
	}
	type fields struct {
		storages []storageContract
	}
	type args struct {
		deleteBy string
	}

	bottomErr, topErr := errors.New("bottom"), errors.New("top")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when a record is set successfully it returns no error",
			fields: fields{
				storages: []storageContract{
					{err: nil},
					{err: nil},
				},
			},
			args: args{
				deleteBy: "key",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "when setting a record in the bottom layer failed it returns the error",
			fields: fields{
				storages: []storageContract{
					{err: topErr},
					{err: bottomErr},
				},
			},
			args: args{
				deleteBy: "key",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, bottomErr)
			},
		},
		{
			name: "when setting a record in the top layer failed it returns the error",
			fields: fields{
				storages: []storageContract{
					{err: topErr},
					{err: nil},
				},
			},
			args: args{
				deleteBy: "key",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, topErr)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storages := make([]Storage[string, string], 0, len(tt.fields.storages))
			for _, st := range tt.fields.storages {
				st := st
				storage := &StorageMock[string, string]{
					DeleteFunc: func(_ context.Context, deleteBy string) error {
						assert.Equal(t, tt.args.deleteBy, deleteBy)
						return st.err
					},
				}
				storages = append(storages, storage)
			}

			c := Cascade[string, string]{
				storages: storages,
			}
			err := c.Delete(context.Background(), tt.args.deleteBy)
			tt.wantErr(t, err)
		})
	}
}

func TestCascade_Get(t *testing.T) {
	type storageContract struct {
		out string
		err error
	}
	type fields struct {
		storages []storageContract
	}
	type args struct {
		getBy string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when a record is found in the bottom layer it gets it successfully",
			fields: fields{
				storages: []storageContract{
					{out: "", err: errors.New("any")},
					{out: "value", err: nil},
				},
			},
			args: args{
				getBy: "key",
			},
			want: "value",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "when a record is found in the top layer it gets it successfully",
			fields: fields{
				storages: []storageContract{
					{out: "value", err: nil},
					{out: "", err: errors.New("any")},
				},
			},
			args: args{
				getBy: "key",
			},
			want: "value",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "when a record is not found it returns ErrNotFound",
			fields: fields{
				storages: []storageContract{
					{out: "", err: nil},
					{out: "", err: nil},
				},
			},
			args: args{
				getBy: "key",
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, ErrNotFound)
			},
		},
		{
			name: "when a record is not found and there is an error it returns an error",
			fields: fields{
				storages: []storageContract{
					{out: "", err: nil},
					{out: "", err: errors.New("any")},
				},
			},
			args: args{
				getBy: "key",
			},
			want: "",
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storages := make([]Storage[string, string], 0, len(tt.fields.storages))
			for _, st := range tt.fields.storages {
				st := st
				storage := &StorageMock[string, string]{
					GetFunc: func(_ context.Context, getBy string) (string, error) {
						assert.Equal(t, tt.args.getBy, getBy)
						return st.out, st.err
					},
					SetFunc: func(_ context.Context, setBy string, record string) error {
						assert.Equal(t, tt.args.getBy, setBy)
						return nil
					},
				}
				storages = append(storages, storage)
			}

			c := Cascade[string, string]{
				storages: storages,
			}
			got, err := c.Get(context.Background(), tt.args.getBy)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCascade_Set(t *testing.T) {
	type storageContract struct {
		err error
	}
	type fields struct {
		storages []storageContract
	}
	type args struct {
		setBy  string
		record string
	}

	bottomErr, topErr := errors.New("bottom"), errors.New("top")

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "when a record is set successfully it returns no error",
			fields: fields{
				storages: []storageContract{
					{err: nil},
					{err: nil},
				},
			},
			args: args{
				setBy:  "key",
				record: "value",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "when setting a record in the bottom layer failed it returns the error",
			fields: fields{
				storages: []storageContract{
					{err: topErr},
					{err: bottomErr},
				},
			},
			args: args{
				setBy:  "key",
				record: "value",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, bottomErr)
			},
		},
		{
			name: "when setting a record in the top layer failed it returns the error",
			fields: fields{
				storages: []storageContract{
					{err: topErr},
					{err: nil},
				},
			},
			args: args{
				setBy:  "key",
				record: "value",
			},
			wantErr: func(t assert.TestingT, err error, msgAndArgs ...interface{}) bool {
				return assert.ErrorIs(t, err, topErr)
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			storages := make([]Storage[string, string], 0, len(tt.fields.storages))
			for _, st := range tt.fields.storages {
				st := st
				storage := &StorageMock[string, string]{
					SetFunc: func(_ context.Context, setBy string, record string) error {
						assert.Equal(t, tt.args.setBy, setBy)
						assert.Equal(t, tt.args.record, record)
						return st.err
					},
				}
				storages = append(storages, storage)
			}

			c := Cascade[string, string]{
				storages: storages,
			}
			err := c.Set(context.Background(), tt.args.setBy, tt.args.record)
			tt.wantErr(t, err)
		})
	}
}
