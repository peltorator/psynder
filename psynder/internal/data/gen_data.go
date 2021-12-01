package data

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"github.com/peltorator/psynder/internal/serviceimpl/shelterservice"
	"github.com/peltorator/psynder/internal/serviceimpl/swipeservice"
	"math/rand"
)


const sheltersNumber = 10
const dogsPerShelter = 20
const passwordLength = 10
const personsNumber = 100
const likeProbability = 0.8


func GenerateData(a *authservice.AuthService,
				  sw *swipeservice.SwipeService,
				  sh *shelterservice.ShelterService) {
	for i := 0; i < sheltersNumber; i++ {
		id, err := a.Signup(auth.SignupArgs{
			Credentials: auth.Credentials{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, false, true, passwordLength),
			},
			Kind: domain.AccountKindShelter,
		})
		if err != nil {
			return
		}

		address := gofakeit.Address()
		err = sh.AddInfo(id, domain.ShelterInfo{
			AccountId: id,
			City:      address.City,
			Address:   address.Street,
			Phone:     gofakeit.Phone(),
		})
		if err != nil {
			return
		}

		for j := 0; j < dogsPerShelter; j++ {
			_, err := sh.AddPsyna(id, swipe.PsynaData{
				Name: gofakeit.PetName(),
				Description: gofakeit.Dog(),
				PhotoLink: gofakeit.ImageURL(640, 480),
			})
			if err != nil {
				return
			}
		}
	}

	for i := 0; i < personsNumber; i++ {
		id, err := a.Signup(auth.SignupArgs{
			Credentials: auth.Credentials{
				Email:    gofakeit.Email(),
				Password: gofakeit.Password(true, true, true, false, true, passwordLength),
			},
			Kind: domain.AccountKindPerson,
		})
		if err != nil {
			return
		}

		info, _ := sw.GetAllInfo()
		psynas, err := sw.BrowsePsynas(id, pagination.Info{
			Limit: int(info.Psynas),
		})
		if err != nil {
			return
		}

		for _, p := range psynas {
			if rand.Float64() < likeProbability {
				err := sw.RatePsyna(id, p.Id, 1)
				if err != nil {
					return
				}
			} else {
				err := sw.RatePsyna(id, p.Id, 2)
				if err != nil {
					return
				}
			}
		}
	}
}