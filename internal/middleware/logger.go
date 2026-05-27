package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Vai chamar a rota de verdade
		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()

		// Retorna o metódo, url da requisição e duração dela
		fmt.Printf("[%s] %s | %d | %v\n", c.Request.Method, c.Request.URL.Path, status, duration)
	}
}