# danger-go

Use Go with [Danger](http://danger.systems)!

This is alpha software, and depends on other alpha software: DangerJS 2.x.

## How to write a Dangerfile

```go
package main

import (
	"log"
	"os"

	danger "github.com/bdotdub/danger-go"
)

func main() {
	danger, results, err := danger.Danger()
	if err != nil {
		log.Fatal(err)
	}
	defer results.Flush(os.Stdout)

	for _, file := range danger.Git.ModifiedFiles {
		// Check something with the files.
	}

	results.Message("✌️ Howdy!")
}
```

## How to use this

As of DangerJS v2.x, you're able to use non-JS Dangerfiles by using: `danger process [executable_name]`.
Danger runs the executable and communicates over standard in/out. See more here: [danger/danger-js#pull/341](https://github.com/danger/danger-js/pull/341)

You will need to write your Go Dangerfile and the compile it into a binary:

    $ go build ./build/my-dangerfile

Then, you can run `danger`:

    $ danger process ./build/my-dangerfile

## Examples

### Running locally

There are some sample Dangerfiles under the `examples` directory. Each comes with it's own fixture. You
can run the samples like so (using `hello-world` in this case):

    $ cat examples/hello-world/fixture.json | go run ./examples/hello-world/main.go

It should output the JSON document that Danger expects.

### Running on a PR

You can also run these examples on a real PR by running the following:

    $ go build -o ./build/hello-world ./examples/hello-world/main.go
    $ DANGER_GITHUB_API_TOKEN=your-github-token DANGER_FAKE_CI="YEP" DANGER_TEST_REPO=username/repo DANGER_TEST_PR=1 danger process ./build/hello-world

## LICENSE

MIT
