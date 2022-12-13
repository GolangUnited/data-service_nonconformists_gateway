package server

import (
	"fmt"
	"go-united/gateway/internal/api/rest"
	"go-united/gateway/internal/service"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type RestServer struct {
	service service.Gateway
}

func New(serv service.Gateway) *RestServer {
	return &RestServer{
		service: serv,
	}
}

func isUuidCValid(uuidVal *string) bool {
	if uuidVal == nil {
		return false
	}
	if _, err := uuid.Parse(*uuidVal); err != nil {
		return false
	}
	return true
}

func (rs RestServer) GetCertificates(ctx echo.Context, params rest.GetCertificatesParams) error {
	if !isUuidCValid(params.UserId) {
		return echo.NewHTTPError(http.StatusBadRequest, "valid userId query parameter required as uuid string")
	}

	resp, err := rs.service.GetCertificates(*params.UserId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (rs RestServer) GetProfile(ctx echo.Context) error {
	fmt.Println("GetProfile not implemented")
	return nil
}
