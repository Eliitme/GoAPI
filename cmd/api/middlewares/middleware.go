package middlewares

import (
	azure_authentication "azure/api/service/azure"

	"github.com/gin-gonic/gin"
)

/*
UserMiddlewares function to add auth
*/
func UserMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {

		//Code for middlewares

		// get header
		token := c.Request.Header.Get("Authorization")

		// if token is empty
		if token == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		// if token is not empty
		// call azure authentication
		authen := azure_authentication.Authenticate(token)

		// if token is invalid
		if authen.Access_token == "" {
			c.JSON(401, gin.H{
				"message": "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
