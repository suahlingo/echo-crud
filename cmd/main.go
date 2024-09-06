package main

import (
	"awesomeProject/config"
	"awesomeProject/models"
	"awesomeProject/routes"
	service2 "awesomeProject/routine/service"
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
	routineService := service2.NewRoutineService("https://www.random.org/integers/?num=1&min=1&max=20&col=1&base=10&format=plain&rnd=new")

	routes.SetupUserRoutes(e, userService)
	routes.SetupRoutineRoutes(e, routineService)

	e.Logger.Fatal(e.Start(":8080"))

}
