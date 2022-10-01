package api

import (
	"azure/api/cmd/api/routers"
	"azure/api/config"
)

func RunApi() {

	e := config.Getenv()

	println(e.Port)

	r := routers.RunRouter()

	r.Run(":" + e.Port)
}
