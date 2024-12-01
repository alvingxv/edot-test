package useruc

import (
	"context"
	"net/http"
	"user-service/internal/interfaces/usecase"
	"user-service/pkg/dto"

	"go.elastic.co/apm/v2"
)

func (uc *userUsecase) Register(ctx context.Context, req *usecase.RegisterRequest) *dto.Response {
	apmSpan, ctx := apm.StartSpan(ctx, "Register", "usecase")
	defer apmSpan.End()

	resp := dto.New()

	err := uc.userRepository.InsertUserToDB(ctx, req.Email, req.Name)
	if err != nil {
		resp.SetError(http.StatusNotFound, err.Status(), err.Message(), err)
		return resp
	}

	resp.SetSuccess(http.StatusOK, "00", "Success Register", nil)

	return resp
}
