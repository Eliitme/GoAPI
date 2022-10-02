package api

import (
	"azure/api/cmd/api/routers"
	"azure/api/config"

	database "azure/api/internal/database"
)

func RunApi() {

	database.Connect()

	e := config.Getenv()

	r := routers.RunRouter()

	r.Run(":" + e.Port)
}
