package httpapi

import "psynder/internal/domain/model"

// /signup

type postSignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type postSignupResponseSuccess struct {
	Token string `json:"token"`
}

// /login

type postLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type postLoginResponseSuccess struct {
	Token string `json:"token"`
}

// /loadpsynas
type postLoadPsynasRequest struct {
	Count     uint64 `json:"count"`
}

type postLoadPsynasResponseSuccess struct {
	Psynas []model.Psyna `json:"psynas"`
}

// /likepsyna

type postLikePsynaRequest struct {
	PsynaId uint64 `json:"psyna_id"`
}

// /getfavoritepsynas

type postGetFavoritePsynasResponseSuccess struct {
	Psynas []model.Psyna `json:"psynas"`
}
