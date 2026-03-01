package main

import (
	"simonwalenkamp/go-up/checker"
	"simonwalenkamp/go-up/config"
)

func main() {
	config := config.Load()

	for _, url := range config.Urls {
		checker.CheckIfUp(url)
	}
}
