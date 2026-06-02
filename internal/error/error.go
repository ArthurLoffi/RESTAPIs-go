package error

import "github.com/gin-gonic/gin"

type AppError struct {
    Status  int    `json:"status"`
    Message string `json:"message"`
}

func New(status int, msg string) *AppError {
    return &AppError{Status: status, Message: msg}
}

func (e *AppError) Error() string {
    return e.Message
}

// internal/error/error.go
func Respond(c *gin.Context, err *AppError) {
    c.JSON(err.Status, err)
    c.Abort()
}