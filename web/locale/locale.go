// Package locale provides localization functionality for the web panel.
package locale

import "github.com/gin-gonic/gin"

// LocalizerMiddleware returns a middleware that handles localization from headers/cookies.
func LocalizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
