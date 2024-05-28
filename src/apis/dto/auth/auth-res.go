package auth

type TokenResDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UserResDto struct {
	ID       string `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type AuthResDto struct {
	Tokens TokenResDto `json:"tokens"`
	User   UserResDto  `json:"user"`
}
