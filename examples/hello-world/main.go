package main

import (
	"log"
	"os"

	danger "github.com/bdotdub/danger-go"
)

func main() {
	_, results, err := danger.Danger()
	if err != nil {
		log.Fatal(err)
	}
	defer results.Flush(os.Stdout)

	results.Message("✌️ Howdy!")
}
