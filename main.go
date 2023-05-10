package main

import (
	"fmt"
	"os"

	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/constants"
	"github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/routes"

	config "github.com/ainurbrr/go_mini-project_moh-ainur-bahtiar-rohman/tree/main/config"

	"github.com/labstack/echo/v4"
)

func main() {

	db := config.Init()
	e := echo.New()
	routes.Routes(e, db)
	port := os.Getenv("PORT")

	if port == "" {
		port = constants.DEFAULT_PORT
	}

	appPort := fmt.Sprintf(":%s", port)

	e.Logger.Fatal(e.Start(appPort))
}
