package main

import (
	"log"

	"github.com/michaelgbenle/fintech/cmd/server"
	"github.com/michaelgbenle/fintech/internal/repository"
)

func main() {
	//Gets the environment variables
	env := server.InitDBParams()

	//Initializes the database
	db, err := repository.Initialize(env.DbUrl)
	log.Println(env.DbUrl)
	if err != nil {
		return
	}

	//Runs the app
	server.Run(db, env.Port)
}
