package main

import (
	"os"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/app"
)

func main() {
	a := app.NewApp()
	a.Run(os.Getenv("API_PORT"))
}
