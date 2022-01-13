package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseMessage struct {
	Code          int         `json:"code"`
	Message       string      `json:"message"`
	Data          interface{} `json:"data,omitempty"`
	CorrelationID string      `json:"correlationId,omitempty"`
}

func SendResponse(ctx echo.Context, response ResponseMessage) error {
	if response.Code == 0 {
		response.Code = http.StatusOK
	}

	if len(response.Message) == 0 {
		response.Message = "success"
	}

	ctx.Response().Header().Set(echo.HeaderServer, "Backend Menus Apps")
	return ctx.JSON(response.Code, response)
}
