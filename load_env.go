package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Environment struct {
	dbString string
	port     string
}

func loadEnv() (Environment, error) {
	err := godotenv.Load()
	if err != nil {
		return Environment{}, err
	}
	port := os.Getenv("PORT")
	dbString := os.Getenv("DB_STRING")

	if port == "" {
		return Environment{}, errors.New("Port variable is not set")
	}

	if dbString == "" {
		return Environment{}, errors.New("DB_STRING env not set")
	}

	return Environment{
		dbString: dbString,
		port:     port,
	}, nil

}
