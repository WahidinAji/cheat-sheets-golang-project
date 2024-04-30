package validations

type ErrValidation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}