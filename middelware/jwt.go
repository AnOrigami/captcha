package middelware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"captcha/util"
)

const SessionCookieName = "_session_"

// 用于中间件和请求方法中claims的传递
type ContextKey[K, T any] struct{}

func (key ContextKey[K, T]) WithValue(ctx context.Context, value T) context.Context {
	return context.WithValue(ctx, key, value)
}

func (key ContextKey[K, T]) Value(ctx context.Context) T {
	return ctx.Value(key).(T)
}

type ContextJWTClaims struct {
	ContextKey[ContextJWTClaims, jwt.RegisteredClaims]
}

func HTTPMiddlewareJWT(jwt *util.JWT) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		// 从请求的 Cookie 中获取 JWT
		cookie, _ := c.Request.Cookie(SessionCookieName)

		// 验证 JWT 并获取声明信息
		claims, err := jwt.Verify(ctx, cookie.Value)
		if err != nil {
			return
		}
		// 将声明信息存储到请求的上下文中
		ctx = ContextJWTClaims{}.WithValue(ctx, claims)
		c.Request = c.Request.WithContext(ctx)
		// 继续处理下一个中间件或请求处理函数
		c.Next()
	}
}
