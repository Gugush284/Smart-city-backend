package Constants

import (
	"errors"
)

type ctxKey int8

var (
	ErrRecordNotFound           = errors.New("record not found")
	ErrIncorrectLoginOrPassword = errors.New("incorrect login or password")
	ErrNotAuthenticated         = errors.New("not authenticated")
	ErrSqlIdNil                 = errors.New("sql return id = 0")
)

const SessionName = "activesession"
const HostUrl = "https://a8790-70b7.s.d-f.pw"

const (
	CtxKeyUser ctxKey = iota
	CtxKeyId
)
