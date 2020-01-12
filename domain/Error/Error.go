package Error

type Error struct {
	message    string
	code       int64
	level      string
	childError *Error
}

func New() (err *Error) {
	err = new(Error)
	return err
}

func (err *Error) SetMessage(message string) *Error {
	err.message = message
	return err
}

func (err *Error) Message() (message string) {
	message = err.message
	return
}

func (err *Error) SetCode(code int64) *Error {
	err.code = code
	return err
}

func (err *Error) Code() (code int64) {
	code = err.code
	return
}

func (err *Error) SetLevel(level string) *Error {
	err.level = level
	return err
}

func (err *Error) Level() (level string) {
	level = err.level
	return
}

func (err *Error) SetChildError(childError *Error) *Error {
	err.childError = childError
	return err
}

func (err *Error) ChildError() (childError *Error) {
	childError = err.childError
	return
}

func (err *Error) IsSetChildError() (isSet bool) {
	if err.childError != nil {
		isSet = true
	}
	return
}
