package main

import (
	"github.com/andremartinsds/flash-cards-api/config"
	"github.com/andremartinsds/flash-cards-api/router"
)

func main() {
	err := config.Init()

	if err != nil {
		//TODO: change to logger
		panic(err)
	}

	router.Initilize()
}
