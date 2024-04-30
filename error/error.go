package error

import "errors"

var (
	ErrPoolInv    = errors.New("invalid pool")
	ErrConnClose  = errors.New("connection closed")
	ErrConnInv    = errors.New("invalid connection")
	ErrNotExists  = errors.New("user id was not found")
	ErrExists     = errors.New("email was already exists")
	ErrConnFailed = errors.New("connection failed")
	ErrQuery      = errors.New("execute query error")
	ErrBeginTx    = errors.New("begin transaction error")
	ErrScan       = errors.New("scan error")
	ErrCommit     = errors.New("commit error")
	ErrExec       = errors.New("query exec error")
	ErrRollback   = errors.New("rollback error")

	ErrHashPass        = errors.New("hash password error")
	ErrPasswordIsEmpty = errors.New("password is empty")
	ErrPasswordIsShort = errors.New("password is short")
)


type CustomError struct {
	
}