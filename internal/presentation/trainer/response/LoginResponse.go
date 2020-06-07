package response

type LoginResponse struct {
	Token  *string `json:"token,omitempty"`
}

func NewLoginResponse(token *string) *LoginResponse {
	return &LoginResponse{
		Token: token,
	}
}
