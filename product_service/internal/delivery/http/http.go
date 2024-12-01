package http

import (
	"context"
	"product-service/config"
	"product-service/constant"
	"product-service/internal/app"
	"product-service/pkg/middleware"

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

	productHandler := NewProdcutHandler(app.Usecases.ProductUsecase, validate)

	r.Use(apmfiber.Middleware())
	r.Use(middleware.LoggingMiddleware)

	r.Get(constant.RouteApiV1+"/products", productHandler.GetProducts)
	// r.Post(constant.RouteApiV1+"/login", userHandler.Login)

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
