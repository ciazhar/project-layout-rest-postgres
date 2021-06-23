package response

import (
	"github.com/ciazhar/project-layout-rest-postgres/third_party/logger"
	"github.com/gofiber/fiber/v2"
	"strings"
)

type Response struct {
	Code    Code        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Code byte

const (
	CodeSuccess      Code = 1
	CodeUnknownError      = 2
	CodeBadRequest        = 3
	CodeNotExist          = 4
)

func Data(data interface{}) Response {
	return Response{
		Message: "succes",
		Code:    CodeSuccess,
		Data:    data,
	}
}

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(Data(data))
}

func Error(err error, code ...Code) error {
	if err == nil {
		return nil
	}

	if code != nil && len(code) == 1 {
		return fiber.NewError(int(code[0]), err.Error())
	}

	logger.Error(err.Error())

	switch r := err.Error(); {
	case strings.Contains(r, "MustExist") == true:
		return fiber.NewError(CodeNotExist, err.Error())
	}

	return fiber.NewError(CodeUnknownError, err.Error())
}
