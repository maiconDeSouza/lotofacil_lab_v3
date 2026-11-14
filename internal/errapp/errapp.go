package errapp

import "errors"

var (
	ErrNoDocuments = errors.New("Documento não encontrado")
)

type ErrAppRes struct {
	Err     error
	Message string
	Code    int
}

func NewErrAppRes(err error, msg string, code int) *ErrAppRes {
	return &ErrAppRes{
		Err:     err,
		Message: msg,
		Code:    code,
	}
}
