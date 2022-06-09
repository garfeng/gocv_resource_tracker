package tracker

import (
	"runtime"
	"sync"
)

type ResourceTracker struct {
	dataList             []Closer
	dataListCloseWithErr []CloserWithError
	mutex                sync.Mutex
}

// NewResourceTracker Returns a ResourceTracker, Close() by manual is not required.
func NewResourceTracker() *ResourceTracker {
	rt := &ResourceTracker{
		dataList:             []Closer{},
		dataListCloseWithErr: []CloserWithError{},
	}
	runtime.SetFinalizer(rt, func(r *ResourceTracker) {
		r.Close()
	})
	return rt
}

func (r *ResourceTracker) SetFinalizerClose() {
	runtime.SetFinalizer(r, func(r *ResourceTracker) {
		r.Close()
	})
}

func (r *ResourceTracker) Close() {
	r.mutex.Lock()
	removeHeap(len(r.dataList))
	for _, v := range r.dataList {
		v.Close()
	}
	r.dataList = r.dataList[:0]

	removeHeap(len(r.dataListCloseWithErr))
	for _, v := range r.dataListCloseWithErr {
		v.Close()
	}
	r.dataListCloseWithErr = r.dataListCloseWithErr[:0]

	r.mutex.Unlock()
}

func (r *ResourceTracker) TrackCloser(data ...Closer) {
	addHeap(len(data))
	r.mutex.Lock()
	r.dataList = append(r.dataList, data...)
	r.mutex.Unlock()
}

func (r *ResourceTracker) TrackCloseError(data ...CloserWithError) {
	addHeap(len(data))
	r.mutex.Lock()
	r.dataListCloseWithErr = append(r.dataListCloseWithErr, data...)
	r.mutex.Unlock()
}

func (r *ResourceTracker) Track(data ...any) {
	errClosers := []CloserWithError{}
	closers := []Closer{}

	for _, v := range data {
		if e, ok := v.(CloserWithError); ok {
			errClosers = append(errClosers, e)
		} else if e2, ok2 := v.(Closer); ok2 {
			closers = append(closers, e2)
		} else {
			panic("tracked data is not closer")
		}
	}

	addHeap(len(data))
	r.mutex.Lock()
	r.dataListCloseWithErr = append(r.dataListCloseWithErr, errClosers...)
	r.dataList = append(r.dataList, closers...)
	r.mutex.Unlock()
}

type CloserWithError interface {
	Close() error
}

type Closer interface {
	Close()
}
