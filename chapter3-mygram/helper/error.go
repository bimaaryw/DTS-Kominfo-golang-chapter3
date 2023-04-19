package helper

import (
	"errors"
	"fmt"
)

func ErrorRequestMessages(field, tag, param string) (message error) {
	switch tag {
	case "required":
		message = errors.New(fmt.Sprintf("%s is required", field))
	case "email":
		message = errors.New(fmt.Sprintf("%s must be valid email", field))
	case "gt":
		message = errors.New(fmt.Sprintf("%s must be grather than %s", field, param))
	case "min":
		message = errors.New(fmt.Sprintf("%s minimum %s characters", field, param))
	default:
		message = errors.New("Error validated request")
	}

	return
}
