package httpapi

// /signup

type postSignupRequest struct {
Email    string `json:"email"`
Password string `json:"password"`
}

type postSignupResponseSuccess struct {
Token string `json:"token"`
}

type postSignupResponseError struct {
Error string `json:"error"`
}

// /login

type postLoginRequest struct {
Email    string `json:"email"`
Password string `json:"password"`
}

type postLoginResponseSuccess struct {
Token string `json:"token"`
}

type postLoginResponseError struct {
Error string `json:"error"`
}