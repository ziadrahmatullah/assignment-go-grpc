package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func WithTimeout(c *gin.Context) {
	ctx := c.Request.Context()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	c.Request = c.Request.WithContext(ctx)
	c.Next()
}

func WithTimeoutInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	ctx2, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	res, err := handler(ctx2, req)
	return res, err 
}
