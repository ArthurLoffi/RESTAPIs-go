package middleware

import (
	"net/http"
	errorf "restapis-go/internal/error"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiters = make(map[string]*rate.Limiter)

func Limiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		if _, exists := limiters[ip]; !exists {
			limiters[ip] = rate.NewLimiter(1, 10)
		}
		limiter := limiters[ip]

		if !limiter.Allow() {
			errorf.New(http.StatusTooManyRequests, "Too many requests")
			c.AbortWithStatus(http.StatusTooManyRequests)
			return
		}

		c.Next()
	}
}