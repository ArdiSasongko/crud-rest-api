package execption

type NotFound struct {
	Error string
}

func NewNotFoundError(error string) NotFound {
	return NotFound{
		Error: error,
	}
}
