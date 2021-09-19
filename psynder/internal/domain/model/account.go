package model

import "strconv"

type AccountId uint64
type PasswordHash []byte

func (id AccountId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type Account struct {
	Id AccountId
}