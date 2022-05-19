package exception

import "errors"

var (
	ErrInvalidLogin = errors.New("unauthentication")
)

type ErrorUnauthenticated struct {
	err     error
	message string
}

func (r *ErrorUnauthenticated) Error() string {
	return r.err.Error()
}

func (r ErrorUnauthenticated) Is(target error) bool {
	return target == ErrInvalidLogin
}

func NewErrorUnauthenticated(message string) error {
	return &ErrorUnauthenticated{
		err:     errors.New("unauthentication"),
		message: message,
	}
}
