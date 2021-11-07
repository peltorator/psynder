package tokenissuer

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"strconv"
	"time"
)

type jwtTokenIssuer struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey

	expire time.Duration
}

type jwtClaims struct {
	Uid string
	jwt.StandardClaims
}

func NewJWT(privateBytes, publicBytes []byte, keyExpiration time.Duration) (*jwtTokenIssuer, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return nil, err
	}
	return &jwtTokenIssuer{
		publicKey:  publicKey,
		privateKey: privateKey,
		expire:     keyExpiration,
	}, nil
}

func (j *jwtTokenIssuer) IssueToken(uid domain.AccountId) (auth.Token, error) {
	claims := jwtClaims{
		Uid: uid.String(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expire).Unix(),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedTok, err := tok.SignedString(j.privateKey)
	if err != nil {
		return "", err
	}
	return auth.NewTokenFromString(signedTok), nil
}

func (j *jwtTokenIssuer) AccountIdByToken(tok auth.Token) (domain.AccountId, error) {
	var claims jwtClaims
	_, err := jwt.ParseWithClaims(tok.String(), &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected tok signing method")
		}
		return j.publicKey, nil
	})
	if err != nil {
		return 0, err
	}
	uid, err := strconv.ParseUint(claims.Uid, 10, 64)
	if err != nil {
		return 0, err
	}
	return domain.AccountId(uid), nil
}
