package chat

type errorCode int

const (
	EINTERNAL errorCode = iota
	EDUPLICATE
)

type Error struct {
	Code errorCode
	Err  error
}

func (e Error) Error() string {
	return e.Err.Error()
}

func NewError(code errorCode, err error) error {
	return Error{Code: code, Err: err}
}

func ErrorCode(err error) errorCode {
	if err, ok := err.(Error); ok {
		return err.Code
	}

	return EINTERNAL
}
