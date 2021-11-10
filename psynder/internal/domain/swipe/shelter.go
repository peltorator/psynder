package swipe

import "github.com/peltorator/psynder/internal/domain"

type ShelterData struct {
	City string `json:"city"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

type Shelter struct {
	Id domain.AccountId `json:"account_id"`
	ShelterData
}
