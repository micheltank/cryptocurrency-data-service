package domain

type Error struct {
	cause   error
	message string
	key     string
	detail  string
}

func NewError(cause error, message, key, detail string) error {
	return &Error{
		cause:   cause,
		message: message,
		key:     key,
		detail:  detail,
	}
}

func (e Error) Key() string {
	return e.key
}

func (e Error) Detail() string {
	return e.detail
}

func (e Error) Error() string {
	return e.message
}

func (e Error) Cause() error {
	return e.cause
}
