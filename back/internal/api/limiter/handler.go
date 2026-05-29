package limiter

import (
	"fmt"
	"sync"
	"time"
)

// voir RWMutex ?
var (
	refreshTime   = time.Second * 2
	refreshNumber = 1
)

type RateHandler struct {
	started bool
	tokens  int
	sync.Mutex
}

func NewRateHandler() *RateHandler {
	return &RateHandler{}
}

func (rh *RateHandler) Start() chan int {
	stopChn := make(chan int)
	go func() {
		rh.started = true
		for {
			select {
			case <-stopChn:
				rh.started = false
				return
			default:
				rh.Mutex.Lock()
				if rh.tokens < refreshNumber {
					rh.tokens = refreshNumber
				}
				rh.Mutex.Unlock()
			}
			time.Sleep(refreshTime)
		}
	}()
	return stopChn
}

func (rh *RateHandler) UseToken() bool {
	fmt.Println(rh.started, rh.tokens)
	if rh.started && rh.tokens > 0 {
		rh.Mutex.Lock()
		rh.tokens--
		rh.Mutex.Unlock()
		return true
	}
	return false
}
