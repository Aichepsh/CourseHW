package payload

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type SendResponse struct {
	Email string `json:"email" validate:"required,email"`
}

type VerifyRequest struct {
	Email string `json:"email" validate:"required,email"`
}
