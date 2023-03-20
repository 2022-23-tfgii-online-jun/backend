package api

import (
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func SentryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				sentry.CaptureException(e.Err)
			}
			sentry.Flush(2 * time.Second)
		}
	}
}
