package routes

import (
	"awesomeProject/users/controller"
	"awesomeProject/users/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

func SetupUserRoutes(e *echo.Echo, Userservice *service.UserService) {

	userController := controller.NewUserController(Userservice)

	e.POST("/users/register", userController.AddUser)
	e.GET("/users/getUser", userController.GetUser)
	e.GET("/users/getUsers", userController.GetUserList)
	e.PUT("/users/updateUser", userController.UpdateUser)
	e.DELETE("/users/deleteUser", userController.DeleteUser)
}

func SetupRouteServiceRoutes(e *echo.Echo, routeService *service.RouteService) {
	e.POST("/for", func(c echo.Context) error {
		var req service.RefreshRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request")
		}

		// 랜덤 값을 반환하는 서비스 호출
		resp := routeService.GetRandomValue(req)
		return c.JSON(http.StatusOK, resp)
	})
}
