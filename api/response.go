package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HTTPRequestValidationFailed = "HTTP Request Validation failed."
	UnknownGitProviderError     = "An error occured during Git Provider request."
)

// https://datatracker.ietf.org/doc/html/rfc7807
type ErrorModel struct {
	Type   string `json:"type,omitempty"`
	Title  string `json:"title,omitempty"`
	Status int    `json:"status,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func (e *ErrorModel) Error() string {
	return e.Detail
}

func (e *ErrorModel) GetStatus() int {
	return e.Status
}

type ContentTypeFilter interface {
	ContentType(string) string
}

type StatusError interface {
	GetStatus() int
	Error() string
}

var NewError = func(status int, msg string) StatusError {
	return &ErrorModel{
		Status: status,
		Title:  http.StatusText(status),
		Detail: msg,
	}
}

func RespondError(ctx *gin.Context, status int, msg string, errs ...error) {
	err := NewError(status, msg)

	ctx.Header("Content-Type", "application/problem+json")
	ctx.JSON(err.GetStatus(), gin.H{"error": err.Error()})
}

func RespondJSON(c *gin.Context, status int, payload interface{}) {
	c.JSON(status, payload)
}
