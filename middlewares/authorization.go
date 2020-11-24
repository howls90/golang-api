package middlewares

import (
	"net/http"
	"strings"

	"api/libs"

	"github.com/gin-gonic/gin"
)

func extractToken(token string) string {
	if strArr := strings.Split(token, " "); len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Verify(c *gin.Context) {

	token := extractToken(c.GetHeader("Authorization"))

	params, err := libs.Verify(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
		c.Abort()
		return
	}

	c.Set("userParams", params)

	// Pass on to the next-in-chain
	c.Next()
}
