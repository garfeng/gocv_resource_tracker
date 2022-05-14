package tracker

import (
	"runtime"
	"sync/atomic"
	"time"
)

var (
	lastGCTime         = time.Now().Unix()
	heapNum            int64
	lastHeapNumAfterGC int64

	monitorStopped int64
)

const (
	forceGCPeriod     = 2 * 60 // 2 minutes
	forceGCNewHeapNum = 100
)

func addHeap(num int) {
	atomic.AddInt64(&heapNum, int64(num))
}
func removeHeap(num int) {
	atomic.AddInt64(&heapNum, -(int64(num)))
}

func CallGCIfRequired() {
	deltaTime := time.Now().Unix() - atomic.LoadInt64(&lastGCTime)
	deltaHeapNum := atomic.LoadInt64(&heapNum) - atomic.LoadInt64(&lastHeapNumAfterGC)

	if deltaTime > forceGCPeriod || deltaHeapNum > forceGCNewHeapNum {
		forceGC()
	}
}

func forceGC() {
	runtime.GC()
	atomic.StoreInt64(&lastHeapNumAfterGC, atomic.LoadInt64(&heapNum))
	atomic.StoreInt64(&lastGCTime, time.Now().Unix())
}

func monitor() {
	for atomic.LoadInt64(&monitorStopped) == 0 {
		CallGCIfRequired()
		<-time.After(time.Second * 5)
	}
}

func init() {
	go monitor()

	runtime.SetFinalizer(&monitorStopped, func(stopped *int64) {
		atomic.StoreInt64(stopped, 1)
	})
}
