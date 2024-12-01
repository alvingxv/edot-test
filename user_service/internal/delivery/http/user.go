package http

import (
	"net/http"
	"user-service/internal/interfaces/delivery"
	"user-service/internal/interfaces/usecase"
	"user-service/pkg/dto"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jsoniter "github.com/json-iterator/go"
	"go.elastic.co/apm/v2"
)

type userHandler struct {
	userUsecase usecase.UserUsecase
	validate    *validator.Validate
}

func NewUserHandler(userUsecase usecase.UserUsecase, validate *validator.Validate) delivery.UserHandler {
	return &userHandler{
		userUsecase: userUsecase,
		validate:    validate,
	}
}

func (h *userHandler) Register(c *fiber.Ctx) error {
	apmSpan, ctx := apm.StartSpan(c.Context(), "Register", "Handler")
	defer apmSpan.End()

	reqBody := c.Body()

	var reqStruct usecase.RegisterRequest

	err := jsoniter.Unmarshal(reqBody, &reqStruct)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		c.Status(http.StatusUnprocessableEntity)
		c.JSON(dto.NewError(http.StatusBadRequest, "FM", "error unmarshall", err))
		return nil
	}

	err = h.validate.Struct(reqStruct)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		c.Status(http.StatusBadRequest)
		c.JSON(dto.NewError(http.StatusBadRequest, "FM", "error unmarshall", err))
		return nil
	}

	resp := h.userUsecase.Register(ctx, &reqStruct)

	c.Status(resp.HttpCode)
	c.JSON(resp)
	return nil
}

func (h *userHandler) Login(c *fiber.Ctx) error {
	apmSpan, ctx := apm.StartSpan(c.Context(), "Register", "Handler")
	defer apmSpan.End()

	reqBody := c.Body()

	var reqStruct usecase.LoginRequest

	err := jsoniter.Unmarshal(reqBody, &reqStruct)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		c.Status(http.StatusUnprocessableEntity)
		c.JSON(dto.NewError(http.StatusBadRequest, "FM", "error unmarshall", err))
		return nil
	}

	err = h.validate.Struct(reqStruct)
	if err != nil {
		apm.CaptureError(ctx, err).Send()
		c.Status(http.StatusBadRequest)
		c.JSON(dto.NewError(http.StatusBadRequest, "VE", "Validation Error", err))
		return nil
	}

	resp := h.userUsecase.Login(ctx, &reqStruct)

	c.Status(resp.HttpCode)
	c.JSON(resp)
	return nil
}
