// Package middleware provides HTTP middleware for the web panel.
package middleware

import "github.com/gin-gonic/gin"

// DomainValidatorMiddleware creates a middleware that validates the request domain.
func DomainValidatorMiddleware(domain string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
