package e

import "errors"

// HttpError : custom http error object used in controller layer
type HttpError struct {
	Status  int
	Code    int
	Message string
	Details string
}

type WrapError interface {
	Code() int
	Error() string
	Message() string
}

type WrapErrorImpl struct {
	ErrorCode int
	Msg       string
	RootCause error
}

// Code implements WrapError
func (e *WrapErrorImpl) Code() int {
	return e.ErrorCode
}

// Error implements WrapError
func (e *WrapErrorImpl) Error() string {
	return e.RootCause.Error()
}

// Message implements WrapError
func (e *WrapErrorImpl) Message() string {
	return e.Msg
}

// NewError : create a new error instance
func NewError(code int, msg string, rootCause ...error) WrapError {
	errCodeMsg := GetErrorMsg(code)
	if msg == "" {
		msg = errCodeMsg
	}
	newErr := &WrapErrorImpl{
		ErrorCode: code,
		Msg:       msg,
	}
	if len(rootCause) > 0 {
		newErr.RootCause = rootCause[0]
	} else {
		newErr.RootCause = errors.New(msg)
	}
	return newErr
}
