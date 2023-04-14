package ginpkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	ReasonCode string `json:"reason_code"`
	Reason     string `json:"reason"`
	Data       any    `json:"data"`
}

func SuccessResponse(c *gin.Context, data any, message ...string) {
	r := &Response{
		Code:       http.StatusOK,
		Message:    "success",
		ReasonCode: "SUCCESS",
		Reason:     "",
		Data:       data,
	}
	if len(message) > 0 {
		r.Message = message[0]
	}
	c.AbortWithStatusJSON(http.StatusOK, r)
}

func BadRequestResponse(c *gin.Context, message string, reason ...string) {
	r := &Response{
		Code:       http.StatusBadRequest,
		Message:    message,
		ReasonCode: "BAD_REQUEST",
		Reason:     "",
		Data:       nil,
	}
	if len(reason) > 0 {
		r.Reason = reason[0]
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, r)
}

func ErrorResponse(c *gin.Context, message string, reason ...string) {
	r := &Response{
		Code:       http.StatusInternalServerError,
		Message:    message,
		ReasonCode: "SERVER_ERROR",
		Reason:     "",
		Data:       nil,
	}
	if len(reason) > 0 {
		r.Reason = reason[0]
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, r)
}
