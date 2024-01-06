// Dagger CI Tooling - Linting
package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"

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

	// go1.21.5
	golangVer := runtime.Version()
	golangVer = strings.ReplaceAll(golangVer, "go", "")
	fmt.Println("Golang Ver: " + golangVer)

	// use a node:16-slim container
	// mount the source code directory on the host
	// at /src in the container
	source := client.Container().
		From("vpayno/ci-generic-debian:latest").
		WithDirectory("/src", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"build/"},
		}))

	// set the working directory in the container
	// install application dependencies
	source = source.WithWorkdir("/src").
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golangVer)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golangVer)).
		WithEnvVariable("GOCACHE", "/go/build-cache")

	runner := source.
		WithExec([]string{"gocritic", "check", "-enableAll", "./..."}).
		WithExec([]string{"gocyclo", "-over", "15", "."}).
		WithExec([]string{"golangci-lint", "run", "./..."}).
		WithExec([]string{"gosec", "./..."}).
		WithExec([]string{"govulncheck", "./..."}).
		WithExec([]string{"ineffassign", "./..."}).
		WithExec([]string{"revive", "./..."}).
		WithExec([]string{"staticcheck", "./..."})

	// run application tests
	out, err := runner.Stderr(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
