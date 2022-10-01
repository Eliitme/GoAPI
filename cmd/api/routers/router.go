package routers

import (
	"azure/api/cmd/api/middlewares"

	"github.com/gin-gonic/gin"

	apiV1 "azure/api/cmd/api/controllers/api/v1"
	apiV2 "azure/api/cmd/api/controllers/api/v2"
)

func RunRouter() *gin.Engine {

	r := gin.Default()

	r.Static("/storage", "storage")

	r.Static("/template", "template")

	r.Use(func(c *gin.Context) {
		// add header Access-Control-Allow-Origin
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}

	})

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.UserMiddlewares())
	{
		v1.GET("/hello-world", apiV1.Helloworld)
	}

	v2 := r.Group("/api/v2")

	v2.GET("hello-world", apiV2.Helloworld)

	return r

}
