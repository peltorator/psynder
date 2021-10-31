package token

import "github.com/peltorator/psynder/internal/domain/model"

type AccessToken string

func (t AccessToken) String() string {
	return string(t)
}

type Issuer interface {
	IssueToken(accountId model.AccountId) (AccessToken, error)
	AccountIdByToken(token AccessToken) (model.AccountId, error)
}