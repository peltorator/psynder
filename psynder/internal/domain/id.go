package domain

import "strconv"

type AccountId uint64

func (id AccountId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}

type PsynaId uint64

func (id PsynaId) String() string {
	return strconv.FormatUint(uint64(id), 10)
}