package github.com/yicheng20110203/locker

type IError interface {
	Error() string
	GetCode() int
}

const (
	ErrorCodeHasBeenUsed = 1 << iota
)

type ErrorWrap struct {
	errCode int
	errMsg  string
}

func NewErrorWrap(errCode int, errMsg string) ErrorWrap {
	return ErrorWrap{
		errCode: errCode,
		errMsg:  errMsg,
	}
}

func (ew ErrorWrap) GetCode() int {
	return ew.errCode
}

func (ew ErrorWrap) Error() string {
	return ew.errMsg
}

var (
	ErrorLockerHasBeenUsed ErrorWrap = NewErrorWrap(ErrorCodeHasBeenUsed, "其他线程已获得锁")
)
