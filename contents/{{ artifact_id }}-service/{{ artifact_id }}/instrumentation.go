package {{artifact_id}}

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           Service
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, countResult metrics.Histogram, s Service) Service {
	return &instrumentingMiddleware{
		requestCount:   counter,
		requestLatency: latency,
		countResult:    countResult,
		next:           s,
	}
}

func (mw instrumentingMiddleware) Status(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "status", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.Status(ctx)
	return
}
