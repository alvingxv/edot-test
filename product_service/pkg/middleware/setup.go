package middleware

import (
	"time"
	"product-service/pkg/log"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {

	start := time.Now()
	log.LogRequest(c, start)

	// Process the request
	c.Next()

	responseStatus := c.Response().StatusCode()
	responseBody := string(c.Response().Body()) // Capture the response body
	responseTime := time.Since(start)

	// Use goroutine for response logging
	go func() {
		log.LogResponse(responseStatus, responseBody, responseTime)

	}()

	return nil
}
