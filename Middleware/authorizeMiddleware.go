package Middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	R "wPan/v1/Response"
	"wPan/v1/Utils"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		userInfo, err := Utils.ParseToken(token[7:])
		if err != nil {
			R.Response(c, http.StatusUnauthorized, R.AUTH_ERROR, nil, http.StatusUnauthorized)
			c.Abort()
			return
		}
		c.Set("userInfo", userInfo)
		c.Next()
		return
	}
}
