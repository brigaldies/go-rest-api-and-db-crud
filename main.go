package main

import (
	"github.com/brigaldies/go-rest-api-and-db-crud/internal/utils"
	"github.com/brigaldies/go-rest-api-and-db-crud/internal/webserver"
)

func main() {

	a := webserver.App{}
	a.Initialize(
		utils.AppGetEnv("APP_DB_USERNAME", "postgres"),
		utils.AppGetEnv("APP_DB_PASSWORD", "secret"),
		utils.AppGetEnv("APP_DB_NAME", "postgres"),
	)
	a.Run(":8081")
}
