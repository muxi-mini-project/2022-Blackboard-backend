package auth

import (
	"blackboard/pkg/token"
	"errors"

	"github.com/gin-gonic/gin"
)

type Context struct {
	ID        string
	ExpiresAt int64 //过期时间
}

var (
	// ErrMissingHeader means the `Authorization` header was empty.
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero")
	// ErrTokenInvalid means the token is invalid.
	ErrTokenInvalid = errors.New("the token is invalid")
)

func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")
	if len(header) == 0 {
		return nil, ErrMissingHeader
	}
	return Parse(header)
}

func Parse(tokenStr string) (*Context, error) {
	t, err := token.ResolveToken(tokenStr)
	if err != nil {
		return nil, err
	}
	return &Context{
		ID:        t.StudentID,
		ExpiresAt: t.ExpiresAt,
	}, nil
}
