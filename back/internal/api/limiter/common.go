package limiter

var rateHandler *RateHandler

func SetRateHandler() {
	rateHandler = NewRateHandler()
	rateHandler.Start()
}

func GetRateHandler() *RateHandler {
	return rateHandler
}
