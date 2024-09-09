package routes

import (
	controller2 "awesomeProject/routine/controller"
	service2 "awesomeProject/routine/service"
	"awesomeProject/users/controller"
	"awesomeProject/users/service"
	"github.com/labstack/echo/v4"
)

func SetupUserRoutes(e *echo.Echo, Userservice *service.UserService) {

	userController := controller.NewUserController(Userservice)

	e.POST("/users/register", userController.AddUser)
	e.GET("/users/getUser", userController.GetUser)
	e.GET("/users/getUsers", userController.GetUserList)
	e.PUT("/users/updateUser", userController.UpdateUser)
	e.DELETE("/users/deleteUser", userController.DeleteUser)
}

func SetupRoutineRoutes(e *echo.Echo, RoutineService *service2.RoutineService) {

	routineController := controller2.NewRoutineController(RoutineService)

	e.GET("/routine/for", routineController.GetDataList)
	e.GET("/routine/goRoutine", routineController.GoRoutine)
}
