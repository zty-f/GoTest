package errors

const (
	SuccessCode = 0
	UnknownCode = 1
)

const (
	SuccessMessage = "Success"
	UnknownMessage = "Unknown"
)

func ErrUnknown(message string) *Error {
	if message == "" {
		message = UnknownMessage
	}
	return New(Stat_FAILED, UnknownCode, message)
}
