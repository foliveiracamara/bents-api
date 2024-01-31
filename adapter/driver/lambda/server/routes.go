package server

import (
	"net/http"

	"github.com/foliveiracamara/bents-api/adapter/driver/lambda/controller"
	"github.com/labstack/echo/v4"
)

func (s *Server) InitRoutes(isLambda string) interface{} {
	userController := controller.InitUserController(s.Echo.AcquireContext())
	eateryController := controller.InitEateryController(s.Echo.AcquireContext())
	s.corsConfig()

	main := s.Echo.Group("/api/v1")
	{
		main.GET("", s.HealthCheck)

		u := main.Group("/user")
		{
			u.POST("", s.controllerWrapper(userController.CreateUser))
			u.POST("/login", s.controllerWrapper(userController.LoginUser))
			u.GET("/:uuid", s.controllerWrapper(userController.GetUser))
		}

		e := main.Group("/eatery")
		{
			e.POST("", s.controllerWrapper(eateryController.CreateEatery))
			e.GET("/list", s.controllerWrapper(eateryController.FindEateries))
			e.GET("/:name", s.controllerWrapper(eateryController.GetEatery))
		}
	}

	if isLambda == "TRUE" {
		return s.Echo
	} else {
		s.Echo.Logger.Fatal(s.Echo.Start(":10000"))
	}

	return nil
}

func (s *Server) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Healthy",
		"code":    http.StatusOK,
	})
}
