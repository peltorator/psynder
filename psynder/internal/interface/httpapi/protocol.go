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
	// TODO: get account id from token
	Count     uint64 `json:"count"`
	AccountId uint64 `json:"account_id"`
}

type postLoadPsynasResponseSuccess struct {
	Psynas []model.Psyna `json:"psynas"`
}

// /likepsyna

type postLikePsynaRequest struct {
	// TODO: get account id from token
	AccountId uint64 `json:"account_id"`
	PsynaId uint64 `json:"psyna_id"`
}

// /getfavoritepsynas

type postGetFavoritePsynasRequest struct {
	// TODO: get account id from token
	AccountId uint64 `json:"account_id"`
}

type postGetFavoritePsynasResponseSuccess struct {
	Psynas []model.Psyna `json:"psynas"`
}
