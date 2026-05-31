package limiter

import (
	shared_infra "language-learning/internal/modules/shared/infra"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLimiter(t *testing.T) {
	t.Run("should block when out of token", func(t *testing.T) {

		mytime := shared_infra.NewDeterministicTimeGenerator(time.Now())
		rl := NewRateHandler(mytime)
		tickCh := make(chan time.Time)
		rl.Start(tickCh)

		tickCh <- time.Now()
		for i := 0; i < refreshNumber; i++ {
			assert.Equal(t, true, rl.UseToken())
		}
		assert.Equal(t, false, rl.UseToken())

		tickCh <- time.Now()
		time.Sleep(time.Second)
		for i := 0; i < refreshNumber; i++ {
			assert.Equal(t, true, rl.UseToken())
		}
		assert.Equal(t, false, rl.UseToken())
	})
}
