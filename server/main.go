package main

import (
	"fmt"
	"os"

	"github.com/alephshahor/Mirlo/server/repositories"
	"github.com/joho/godotenv"
)

func main() {
	configFilePath := os.Getenv("GOPATH") + "/src/github.com/alephshahor/Mirlo/server/.env"
	if err := godotenv.Load(configFilePath); err != nil {
		panic(fmt.Errorf("Fatal error loading .env file: %s \n", err))
	}

	if _, err := repositories.DB().Exec("SELECT 1"); err != nil {
		panic(err)
	}
}
