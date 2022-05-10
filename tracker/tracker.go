package tracker

import (
	"fmt"
	"runtime"
	"sync"
)

type ResourceTracker struct {
	dataList             []Closer
	dataListCloseWithErr []CloserWithError
	mutex                sync.Mutex
}

func (r *ResourceTracker) Close() {
	r.mutex.Lock()

	for _, v := range r.dataList {
		v.Close()
	}
	r.dataList = r.dataList[:0]

	for _, v := range r.dataListCloseWithErr {
		v.Close()
	}
	r.dataListCloseWithErr = r.dataListCloseWithErr[:0]

	r.mutex.Unlock()
	fmt.Println("all closed")
}

func (r *ResourceTracker) TrackCloser(data ...Closer) {
	r.mutex.Lock()
	r.dataList = append(r.dataList, data...)
	r.mutex.Unlock()
}

func (r *ResourceTracker) TrackCloseError(data ...CloserWithError) {
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

	r.mutex.Lock()
	r.dataListCloseWithErr = append(r.dataListCloseWithErr, errClosers...)
	r.dataList = append(r.dataList, closers...)
	r.mutex.Unlock()
}

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

type CloserWithError interface {
	Close() error
}

type Closer interface {
	Close()
}
