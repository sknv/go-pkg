// Code generated by moq, but slightly modified to support generics.
// github.com/matryer/moq

package cascade

import (
	"context"
	"sync"
)

// Ensure, that StorageMock does implement cascade.Storage.
// If this is not the case, regenerate this file with moq.
var _ Storage[any, struct{}] = &StorageMock[any, struct{}]{}

// StorageMock is a mock implementation of cascade.Storage.
//
// 	func TestSomethingThatUsesStorage(t *testing.T) {
//
// 		// make and configure a mocked cascade.Storage
// 		mockedStorage := &StorageMock{
// 			DeleteFunc: func(ctx context.Context, deleteBy Q) error {
// 				panic("mock out the Delete method")
// 			},
// 			GetFunc: func(ctx context.Context, getBy Q) (R, error) {
// 				panic("mock out the Get method")
// 			},
// 			SetFunc: func(ctx context.Context, setBy Q, record R) error {
// 				panic("mock out the Set method")
// 			},
// 		}
//
// 		// use mockedStorage in code that requires cascade.Storage
// 		// and then make assertions.
//
// 	}
type StorageMock[Q any, R comparable] struct {
	// DeleteFunc mocks the Delete method.
	DeleteFunc func(ctx context.Context, deleteBy Q) error

	// GetFunc mocks the Get method.
	GetFunc func(ctx context.Context, getBy Q) (R, error)

	// SetFunc mocks the Set method.
	SetFunc func(ctx context.Context, setBy Q, record R) error

	// calls tracks calls to the methods.
	calls struct {
		// Delete holds details about calls to the Delete method.
		Delete []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// DeleteBy is the deleteBy argument value.
			DeleteBy Q
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// GetBy is the getBy argument value.
			GetBy Q
		}
		// Set holds details about calls to the Set method.
		Set []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// SetBy is the setBy argument value.
			SetBy Q
			// Record is the record argument value.
			Record R
		}
	}
	lockDelete sync.RWMutex
	lockGet    sync.RWMutex
	lockSet    sync.RWMutex
}

// Delete calls DeleteFunc.
func (mock *StorageMock[Q, R]) Delete(ctx context.Context, deleteBy Q) error {
	if mock.DeleteFunc == nil {
		panic("StorageMock.DeleteFunc: method is nil but Storage.Delete was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		DeleteBy Q
	}{
		Ctx:      ctx,
		DeleteBy: deleteBy,
	}
	mock.lockDelete.Lock()
	mock.calls.Delete = append(mock.calls.Delete, callInfo)
	mock.lockDelete.Unlock()
	return mock.DeleteFunc(ctx, deleteBy)
}

// DeleteCalls gets all the calls that were made to Delete.
// Check the length with:
//     len(mockedStorage.DeleteCalls())
func (mock *StorageMock[Q, R]) DeleteCalls() []struct {
	Ctx      context.Context
	DeleteBy Q
} {
	var calls []struct {
		Ctx      context.Context
		DeleteBy Q
	}
	mock.lockDelete.RLock()
	calls = mock.calls.Delete
	mock.lockDelete.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *StorageMock[Q, R]) Get(ctx context.Context, getBy Q) (R, error) {
	if mock.GetFunc == nil {
		panic("StorageMock.GetFunc: method is nil but Storage.Get was just called")
	}
	callInfo := struct {
		Ctx   context.Context
		GetBy Q
	}{
		Ctx:   ctx,
		GetBy: getBy,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(ctx, getBy)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedStorage.GetCalls())
func (mock *StorageMock[Q, R]) GetCalls() []struct {
	Ctx   context.Context
	GetBy Q
} {
	var calls []struct {
		Ctx   context.Context
		GetBy Q
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Set calls SetFunc.
func (mock *StorageMock[Q, R]) Set(ctx context.Context, setBy Q, record R) error {
	if mock.SetFunc == nil {
		panic("StorageMock.SetFunc: method is nil but Storage.Set was just called")
	}
	callInfo := struct {
		Ctx    context.Context
		SetBy  Q
		Record R
	}{
		Ctx:    ctx,
		SetBy:  setBy,
		Record: record,
	}
	mock.lockSet.Lock()
	mock.calls.Set = append(mock.calls.Set, callInfo)
	mock.lockSet.Unlock()
	return mock.SetFunc(ctx, setBy, record)
}

// SetCalls gets all the calls that were made to Set.
// Check the length with:
//     len(mockedStorage.SetCalls())
func (mock *StorageMock[Q, R]) SetCalls() []struct {
	Ctx    context.Context
	SetBy  Q
	Record R
} {
	var calls []struct {
		Ctx    context.Context
		SetBy  Q
		Record R
	}
	mock.lockSet.RLock()
	calls = mock.calls.Set
	mock.lockSet.RUnlock()
	return calls
}
