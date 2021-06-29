package errors

var (
	// client 1001 - 1999
	ClientContextTimeout = NewError(1001, "Client Context Timeout")
	ClientOptionError = NewError(1001, "Client Option Error")
	// server 2001 - 2999
	ServerContextTimeout = NewError(2001, "Server Context Timeout")
	// others
	UnknownNetworkType = NewError(3001, "Unknown Network type")

)

type Error interface {
	ID() int

	Error() string
}

func NewError(id int, desc string) Error {
	return &error{
		id:   id,
		desc: desc,
	}
}

type error struct {
	id   int
	desc string
}

func (e *error) ID() int {
	return e.id
}

func (e *error) Error() string {
	return e.desc
}
