package httpentity

type EmailLoginRequest struct {
	Email    string `json:"email" validate:"required,customEmail"`
	Password string `json:"password" validate:"omitempty,min=8,max=20"`
}

func (input *EmailLoginRequest) Validate() []FieldError {
	return validate(input)
}

type UserTokenResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int    `json:"expires_in,omitempty"`
}

type UserRefreshRequest struct {
	UserId       string `json:"user_id" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (u *UserRefreshRequest) Validate() []FieldError {
	return validate(u)
}
