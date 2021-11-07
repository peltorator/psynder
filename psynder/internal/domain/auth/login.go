package auth

type Token string

func (t Token) String() string {
	return string(t)
}

func NewTokenFromString(s string) Token {
	return Token(s)
}

type Credentials struct {
	Email string
	Password string
}