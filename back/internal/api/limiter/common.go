package limiter

import (
	shared_infra "language-learning/internal/modules/shared/infra"
	"time"
)

var rateHandler *RateHandler

func SetRateHandler() {
	rateHandler = NewRateHandler(shared_infra.NewTimeGenerator())
	ticker := time.NewTicker(refreshTime).C
	rateHandler.Start(ticker)
}

func GetRateHandler() *RateHandler {
	return rateHandler
}
