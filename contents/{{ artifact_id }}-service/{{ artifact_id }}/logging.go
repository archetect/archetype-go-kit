package {{artifact_id}}

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   Service
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingMiddleware{logger, s}
}

func (mw loggingMiddleware) Status(ctx context.Context) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "status",
			"input", "status",
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Status(ctx)
	return
}
