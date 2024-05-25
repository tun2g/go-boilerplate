package auth

type LoginReqDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
