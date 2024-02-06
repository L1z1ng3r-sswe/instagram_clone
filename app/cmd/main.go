package main

import (
	"log"
	"os"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/app"
)

func main() {
	a := app.NewApp()

	if err := a.Run(os.Getenv("API_PORT"), a.Handler.SetUpRoutes()); err != nil {
		log.Fatal("Failed to run server: " + err.Error())
	}
}
