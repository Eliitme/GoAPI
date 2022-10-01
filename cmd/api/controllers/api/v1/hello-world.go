package v1

import (
	u "azure/api/cmd/api/apiHelpers"

	"github.com/gin-gonic/gin"
)

func Helloworld(c *gin.Context) {

	u.Respond(c.Writer, u.Message(1, "Hello World v1"))

}
