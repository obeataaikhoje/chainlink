package utils

import (
	"context"
	"math"
	mrand "math/rand"
	"strings"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/services"
)

// WithJitter adds +/- 10% to a duration
func WithJitter(d time.Duration) time.Duration {
	// #nosec
	if d == 0 {
		return 0
	}
	// ensure non-zero arg to Intn to avoid panic
	max := math.Max(float64(d.Abs())/5.0, 1.)
	// #nosec - non critical randomness
	jitter := mrand.Intn(int(max))
	jitter = jitter - (jitter / 2)
	return time.Duration(int(d) + jitter)
}

// ContextFromChan creates a context that finishes when the provided channel
// receives or is closed.
// When channel closes, the ctx.Err() will always be context.Canceled
// NOTE: Spins up a goroutine that exits on cancellation.
// REMEMBER TO CALL CANCEL OTHERWISE IT CAN LEAD TO MEMORY LEAKS
func ContextFromChan(chStop <-chan struct{}) (context.Context, context.CancelFunc) {
	return services.StopRChan(chStop).NewCtx()
}

// ContextWithDeadlineFn returns a copy of the parent context with the deadline modified by deadlineFn.
// deadlineFn will only be called if the parent has a deadline.
// The new deadline must be sooner than the old to have an effect.
func ContextWithDeadlineFn(ctx context.Context, deadlineFn func(orig time.Time) time.Time) (context.Context, context.CancelFunc) {
	cancel := func() {}
	if d, ok := ctx.Deadline(); ok {
		if m := deadlineFn(d); m.Before(d) {
			ctx, cancel = context.WithDeadline(ctx, m)
		}
	}
	return ctx, cancel
}

func IsZero[C comparable](val C) bool {
	var zero C
	return zero == val
}

// AddHexPrefix adds the prefix (0x) to a given hex string.
func AddHexPrefix(str string) string {
	if len(str) < 2 || len(str) > 1 && strings.ToLower(str[0:2]) != "0x" {
		str = "0x" + str
	}
	return str
}
