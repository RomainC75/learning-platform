package bootstrap

import (
	"language-learning/internal/api"
	"language-learning/internal/api/limiter"
	validatorHandler "language-learning/internal/api/validator"
)

func Bootstrap() {
	validatorHandler.SetValidator()
	limiter.SetRateHandler()
	api.Serve()

}
