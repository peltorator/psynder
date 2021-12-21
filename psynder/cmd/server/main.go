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
		fmt.Printf("Run AUTH\n")
		run_auth(a, aCfg)
	}()

	sw, swCfg := getSwipeService()
	go func() {
		fmt.Printf("Run SWIPES\n")
		run_swipes(a, sw, swCfg)
	}()

	sh, shCfg := getShelterService()
	go func() {
		fmt.Printf("Run SHELTERS\n")
		run_shelters(a, sh, shCfg)
	}()

	go func() {
		fmt.Printf("Generate Data\n")
		data.GenerateData(a, sw, sh)
	}()

	for {
		time.Sleep(10)
	}

}
