package er

import (
	"fmt"
)

var (
	// rest api
	ErrorIsNotMatch        = New(0, "")
	ErrorSaving     = New(10, "ошибка сохранения")
	ErrorMetchData     = New(20, "ошибка сохранения")
)

type Error struct {
	ErrorCode        int    `json:"error_code"`
	ErrorDescription string `json:"error_description"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("err_type:%d , err_des:%s", e.ErrorCode, e.ErrorDescription)
}

func New(code int, desc string) error {
	return &Error{
		ErrorCode:        code,
		ErrorDescription: desc,
	}
}
