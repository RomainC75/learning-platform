package bootstrap

import (
	"language-learning/internal/api"
	"language-learning/internal/api/limiter"
)

func Bootstrap() {
	limiter.SetRateHandler()
	api.Serve()

}
