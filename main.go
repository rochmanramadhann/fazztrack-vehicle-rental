package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/rochmanramadhann/fazztrack-vehicle/src/configs/command"
)

func main() {
	if err := command.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
