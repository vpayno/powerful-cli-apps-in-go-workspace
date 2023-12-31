// Dagger CI Tooling - Testing
package main

import (
	"context"
	"fmt"
	"os"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()

	// initialize Dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// use a node:16-slim container
	// mount the source code directory on the host
	// at /src in the container
	source := client.Container().
		From("vpayno/ci-generic-debian:latest").
		WithDirectory("/src", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"ci/", "build/"},
		}))

	// set the working directory in the container
	// install application dependencies
	runner := source.WithWorkdir("/src").
		WithExec([]string{"./scripts/go-test-with-coverage"})

	// run application tests
	out, err := runner.Stderr(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
