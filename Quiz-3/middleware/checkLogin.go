package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk check login
func CekLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodDelete {
			username, password, ok := c.Request.BasicAuth()
			if !ok || username != "admin" || password != "admin" {
				if username == "admin" {
					if password != "password" {
						c.JSON(http.StatusMethodNotAllowed, gin.H{
							"error": "Password salah!",
						})
						c.Abort()
						return
					}
				} else if username == "editor" {
					if password != "secret" {
						c.JSON(http.StatusMethodNotAllowed, gin.H{
							"error": "Password salah!",
						})
						c.Abort()
						return
					}
				} else {
					c.JSON(http.StatusMethodNotAllowed, gin.H{
						"error": "Login dulu maszzeh!",
					})
					c.Abort()
					return
				}
			}

			c.Next()
		}
	}
}
