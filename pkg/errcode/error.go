package errcode

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	code int      `json:"code"`
	msg  string   `json:"msg"`
	data []string `json:"data"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Fatalf("错误码 %d 已经存在，请更换一个", code)
	}

	codes[code] = msg

	return &Error{code: code, msg: msg}
}

func (e *Error) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.data
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.data = []string{}
	for _, d := range details {
		newError.data = append(newError.data, d)
	}

	return &newError
}

func (e *Error) StatusCode() int {

	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}

	return http.StatusInternalServerError
}
