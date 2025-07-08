package utils

import (
	"encoding/json"
	"fmt"
	"github.com/qianqianzyk/AILesson-Planner/internal/logs"
	"github.com/qianqianzyk/AILesson-Planner/internal/types"
	"go.uber.org/zap"
	"net/http"
	"regexp"
)

const StatusOk = 200

type ApiError struct {
	StatusCode int
	Code       int
	Message    string
	Level      logs.Level
	Data       any
}

func (e ApiError) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

func (e ApiError) Response() *types.Base {
	return &types.Base{
		Code: e.Code,
		Msg:  e.Message,
	}
}

func NewError(code int, level logs.Level, msg string) ApiError {
	apiErr := ApiError{
		StatusCode: StatusOk,
		Code:       code,
		Message:    msg,
		Level:      level,
		Data:       nil,
	}
	return apiErr
}

func AbortWithException(apiErr ApiError, err error) error {
	LogError(&apiErr, err)
	return &apiErr
}

func MatchRegexp(pattern string, value string) bool {
	r := regexp.MustCompile(pattern)
	return r.MatchString(value)
}

func HandleError(w http.ResponseWriter, err error) {
	var apiError *ApiError
	if ok := AsApiError(err, &apiError); ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(apiError.StatusCode)
		resp, _ := json.Marshal(apiError.Response())
		w.Write(resp)
		return
	}
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func AsApiError(err error, target **ApiError) bool {
	apiErr, ok := err.(*ApiError)
	if ok {
		*target = apiErr
	}
	return ok
}

func LogError(apiErr *ApiError, err error) {
	logFields := []zap.Field{
		zap.Int("error_code", apiErr.Code),
		zap.Error(err),
	}
	logs.GetLogFunc(apiErr.Level)(apiErr.Message, logFields...)
}
