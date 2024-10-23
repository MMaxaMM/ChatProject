package chat

type errorCode int

const (
	EINTERNAL     errorCode = iota // Неизвестная ошибка
	EDUPLICATE                     // Пользователь с таким username уже существует
	EUNAUTHORIZED                  // Неверное имя пользователя или пароль
	EFOREIGNKEY                    // Нарушение ограничений базы данных, несуществующий внешний ключ
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
