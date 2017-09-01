package main

import (
	"log"
	"os"
	"strings"

	danger "github.com/bdotdub/danger-go"
)

func main() {
	danger, results, err := danger.Danger()
	if err != nil {
		log.Fatal(err)
	}
	defer results.Flush(os.Stdout)

	if strings.TrimSpace(danger.PullRequest.Body) == "" {
		results.Warn("PRs messages should not be empty.")
	}
}
