package http

import (
	"context"
	"user-service/config"
	"user-service/constant"
	"user-service/internal/app"
	"user-service/pkg/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.elastic.co/apm/module/apmfiber/v2"
)

type HttpServer struct {
	*fiber.App
}

func NewHttpServer(app *app.App) (*HttpServer, error) {
	validate := validator.New()
	r := fiber.New()

	userHandler := NewUserHandler(app.Usecases.UserUsecase, validate)

	r.Use(apmfiber.Middleware())
	r.Use(middleware.LoggingMiddleware)

	r.Post(constant.RouteApiV1+"/register", userHandler.Register)
	r.Post(constant.RouteApiV1+"/login", userHandler.Login)
	// todo validate-jwt this endpoint used by other services to ensure the validity of the jwt

	r.Get("/healthz", func(c *fiber.Ctx) error {
		c.Set("Content-Security-Policy", "default-src 'self'")
		c.Status(200)
		return c.SendString("ok")
	})

	return &HttpServer{
		r,
	}, nil
}

func (s *HttpServer) Run() error {
	return s.Listen(":" + config.Cfg.App.Port)
}

func (s *HttpServer) Stop(ctx context.Context) error {
	return s.Shutdown()
}
