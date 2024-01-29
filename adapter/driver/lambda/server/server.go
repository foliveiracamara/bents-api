package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo *echo.Echo
}

func (s *Server) corsConfig() {
	allowedFrontEndUrls := make([]string, 1)
	allowedFrontEndUrls[0] = "" // Should be bents-front-end url, like "https://bents-front-end.s3.amazonaws.com

	allowedHeaders := make([]string, 2)
	allowedHeaders[0] = "Content-Type"
	allowedHeaders[1] = "Authorization"

	s.Echo.Use(middleware.Logger())
	s.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: allowedFrontEndUrls,
		AllowHeaders: allowedHeaders,
	}))
	
	// s.Echo.Use(middlewares.AuthMiddleware)
}

func (s *Server) controllerWrapper(f func(c echo.Context) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := f(c)
		if err != nil {
			return err
		}
		return nil
	}
}