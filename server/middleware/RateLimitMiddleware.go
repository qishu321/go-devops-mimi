package middleware

import (
	"time"

	systemRsp "go-devops-mimi/server/model/system/response"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func RateLimitMiddleware(fillInterval time.Duration, capacity int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucket(fillInterval, capacity)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			systemRsp.Fail(c, nil, "访问限流")
			c.Abort()
			return
		}
		c.Next()
	}
}
