package pharaoh

type Error struct {
	msg string
}

func newError(msg string) Error {
	return Error{msg}
}

func (e Error) Error() string {
	return e.msg
}
