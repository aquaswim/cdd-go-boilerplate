package errorx

import "fmt"

type Errorx struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Data    *interface{} `json:"error,omitempty"`
	Edited  bool         `json:"edited"`
}

func New(code string, message string) error {
	return &Errorx{Code: code, Message: message}
}

func WrapError(code string, err error) error {
	return &Errorx{Code: code, Message: fmt.Sprint(err)}
}

func NewEdited(code string, message string) error {
	return &Errorx{Code: code, Message: message, Edited: true}
}

func WrapErrorEdited(code string, err error) error {
	return &Errorx{Code: code, Message: fmt.Sprint(err), Edited: true}
}

func NewWithData(code string, message string, data *interface{}) error {
	return &Errorx{Code: code, Message: message, Data: data}
}

func WrapErrorWithData(code string, err error, data *interface{}) error {
	return &Errorx{Code: code, Message: fmt.Sprint(err), Data: data}
}

func (e *Errorx) Error() string {
	return fmt.Sprintf("code: %s, message: %s, data: %+v", e.Code, e.Message, e.Data)
}
