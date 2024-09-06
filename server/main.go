package main

import (
	"io"
	"log"
	"os"

	"lifelogger/server/router"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("empty .env " + err.Error())
	}

	//config.InitializeDB()

	multiWriter := io.MultiWriter(os.Stdout)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := router.Initialize(multiWriter)

	// Start the server on port 3000

	log.Fatal(app.Listen(":22250"))
}
