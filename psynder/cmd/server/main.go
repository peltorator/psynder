package main

import (
	"fmt"
	"github.com/peltorator/psynder/internal/data"
	"time"
)

const ReadTimeout = 10 * time.Second
const WriteTimeout = 10 * time.Second

func main() {
	a, aCfg := getAuthService()
	go func() {
		run_auth(a, aCfg)
	}()

	sw, swCfg := getSwipeService()
	go func() {
		fmt.Printf("Run SWIPES")
		run_swipes(a, sw, swCfg)
	}()

	sh, shCfg := getShelterService()
	run_shelters(a, sh, shCfg)

	data.GenerateData(a, sw, sh)
}
