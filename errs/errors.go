package errs

type Apperror struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func (e Apperror) AsMessage() *Apperror {
	return &Apperror{
		Message: e.Message,
	}
}
func NewErrorNotFound(meg string) *Apperror {
	return &Apperror{
		Message: meg,
		Code:    404,
	}
}
func NewUnexpectedServerError(meg string) *Apperror {
	return &Apperror{
		Message: meg,
		Code:    501,
	}
}
