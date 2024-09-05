package routes

import (
	"awesomeProject/users/controller"
	"awesomeProject/users/service"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, Userservice *service.UserService) {

	userController := controller.NewUserController(Userservice)

	e.POST("/users/register", userController.AddUser)
	e.GET("/users/getUser", userController.GetUser)
	e.GET("/users/getUsers", userController.GetUserList)
	e.PUT("/users/updateUser", userController.UpdateUser)
	e.DELETE("/users/deleteUser", userController.DeleteUser)
}
