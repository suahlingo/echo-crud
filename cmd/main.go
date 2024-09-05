package main

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/routes"
	"awesomeProject/users/service"
	"github.com/labstack/echo/v4"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	config.LoadConfig()

	if err := models.Migrate(config.DB); err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	e := echo.New()
	userService := service.NewUserService(config.DB)
	routes.SetupRouteServiceRoutes(e, routeService)

	routes.SetupUserRoutes(e, userService)

	e.Logger.Fatal(e.Start(":8080"))

}
