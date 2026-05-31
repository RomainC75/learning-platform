package limiter

import (
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	"sync"
	"time"
)

// voir RWMutex ?
var (
	refreshTime   = time.Second * 2
	refreshNumber = 1
)

type RateHandler struct {
	tokens        int
	timeGenerator shared_domain_time.TimeGenerator
	sync.Mutex
}

func NewRateHandler(tg shared_domain_time.TimeGenerator) *RateHandler {
	return &RateHandler{
		timeGenerator: tg,
	}
}

func (rh *RateHandler) Start(tick <-chan time.Time) chan int {
	stopChn := make(chan int)
	go func() {
		for range tick {
			select {
			case <-stopChn:
				return
			default:
				rh.Mutex.Lock()
				if rh.tokens < refreshNumber {
					rh.tokens = refreshNumber
				}
				rh.Mutex.Unlock()
			}
		}
	}()
	return stopChn
}

func (rh *RateHandler) UseToken() bool {
	if rh.tokens > 0 {
		rh.Mutex.Lock()
		rh.tokens--
		rh.Mutex.Unlock()
		return true
	}
	return false
}
