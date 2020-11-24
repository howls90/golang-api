package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	// return func(c *gin.Context) {
	c.Next()
	// 	for _, e := range c.Errors {
	// 		log.Error(e)
	// 	}
	// }

}
