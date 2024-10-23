package chat

type errorCode int

const (
	EINTERNAL errorCode = iota
	EDUPLICATE
	EUNAUTHORIZED
	EFOREIGNKEY
)

type Error struct {
	Code      errorCode
	ErrString string
}

func (e Error) Error() string {
	return e.ErrString
}

func NewError(code errorCode, errString string) error {
	return Error{Code: code, ErrString: errString}
}

func ErrorCode(err error) errorCode {
	if err, ok := err.(Error); ok {
		return err.Code
	}

	return EINTERNAL
}
