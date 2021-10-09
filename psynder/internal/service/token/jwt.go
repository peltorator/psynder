package token

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"psynder/internal/domain/model"
	"strconv"
	"time"
)

// TODO: HANDLE ERRORS PROPERLY

type JwtHandler struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey

	expire time.Duration
}

type jwtClaims struct {
	Id string
	jwt.StandardClaims
}

func NewJwtHandler(privateBytes, publicBytes []byte, keyExpiration time.Duration) (*JwtHandler, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, err
	}
	return &JwtHandler{
		publicKey:  publicKey,
		privateKey: privateKey,
		expire:     keyExpiration,
	}, nil
}

func (j *JwtHandler) IssueToken(accountId model.AccountId) (AccessToken, error) {
	claims := jwtClaims{
		Id: accountId.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expire).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err := token.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}
	return AccessToken(accessToken), nil
}

func (j *JwtHandler) AccountIdByToken(token AccessToken) (model.AccountId, error) {
	var claims jwtClaims
	_, err := jwt.ParseWithClaims(token.String(), &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected token signing method")
		}
		return j.publicKey, nil
	})
	if err != nil {
		return 0, err
	}
	//claims, ok := token.Claims.(*jwtClaims)
	//if !ok {
	//	return "", errors.New("invalid token claims")
	//}
	id, err := strconv.ParseUint(claims.Id, 10, 64)
	if err != nil {
		return 0, err
	}
	return model.AccountId(id), nil
}