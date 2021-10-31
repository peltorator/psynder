package httpapi

import "github.com/peltorator/psynder/internal/domain/model"

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
	Limit     uint64 `json:"limit"`
	Offset     uint64 `json:"offset"`
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
