package main

import "fmt"

func main() {
	go func() {
		fmt.Printf("Run SHELTERS")
		run_shelters()

	}()
	fmt.Printf("Run ALL")
	run()
}