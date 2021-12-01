package data

import (
	"encoding/json"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/peltorator/psynder/internal/domain"
	"github.com/peltorator/psynder/internal/domain/auth"
	"github.com/peltorator/psynder/internal/domain/swipe"
	"github.com/peltorator/psynder/internal/pagination"
	"github.com/peltorator/psynder/internal/serviceimpl/authservice"
	"github.com/peltorator/psynder/internal/serviceimpl/shelterservice"
	"github.com/peltorator/psynder/internal/serviceimpl/swipeservice"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
)

const sheltersNumber = 10
const dogsPerShelter = 20
const passwordLength = 10
const personsNumber = 100
const likeProbability = 0.8
const seed = 1337

var breed = regexp.MustCompile("https://images\\.dog\\.ceo/breeds/([^/-]*)(?:-([^/]*))?/.*")

func getBreedFromURL(url string) string {
	res := ""
	match := breed.FindStringSubmatch(url)
	for i := len(match) - 1; i > 0; i-- {
		res += strings.Title(match[i])
		if i > 1 && match[i] != "" {
			res += " "
		}
	}
	return res
}

func GenerateData(a *authservice.AuthService,
				  sw *swipeservice.SwipeService,
				  sh *shelterservice.ShelterService) {

	gofakeit.Seed(seed)

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
			Phone:     gofakeit.PhoneFormatted(),
		})
		if err != nil {
			return
		}

		r, err := http.Get(fmt.Sprintf("https://dog.ceo/api/breeds/image/random/%v", dogsPerShelter))
		if err != nil {
			return
		}

		type dogsType struct {
			Message []string `json:"message"`
			Status  string   `json:"status"`
		}

		var dogs dogsType
		err = json.NewDecoder(r.Body).Decode(&dogs)
		if err != nil {
			return
		}

		for j := 0; j < dogsPerShelter; j++ {
			dogURL := dogs.Message[j]
			breed := getBreedFromURL(dogURL)
			_, err := sh.AddPsyna(id, swipe.PsynaData{
				Name:        gofakeit.PetName(),
				Breed:       breed,
				Description: breed,
				PhotoLink:   dogURL,
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
		psynas, err := sw.BrowsePsynas(
			id,
			pagination.Info{Limit: int(info.Psynas)},
			domain.PsynaFilter{},
		)
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