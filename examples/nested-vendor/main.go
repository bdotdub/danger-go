package main

import (
	"log"
	"os"
	"regexp"

	danger "github.com/bdotdub/danger-go"
)

var (
	nestedVendorRegex = regexp.MustCompile("vendor/.*/vendor")
)

func main() {
	danger, results, err := danger.Danger()
	if err != nil {
		log.Fatal(err)
	}
	defer results.Flush(os.Stdout)

	if hasNestedVendor(danger.Git.CreatedFiles) || hasNestedVendor(danger.Git.ModifiedFiles) {
		results.Fail("PR includes `vendor` directory of `vendor` libraries.\nUsing the `-v` flag with glide will strip them out.")
	}
}

func hasNestedVendor(files []string) bool {
	for _, file := range files {
		if nestedVendorRegex.MatchString(file) {
			return true
		}
	}

	return false
}
