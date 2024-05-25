// Dagger CI Tooling - Testing
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
	// at /repo in the container
	source := client.Container().
		From("vpayno/ci-generic-debian:latest").
		WithDirectory("/repo", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"build/"},
		}))

	source = source.WithWorkdir("/repo").
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golangVer)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golangVer)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithExec([]string{"pwd"}).
		WithExec([]string{"ls", "-lvh"})

	// set the working directory in the container
	// install application dependencies
	runner := source.WithExec([]string{"./scripts/go-test-with-coverage"})

	// run application tests
	stdout, err := runner.Stdout(ctx)
	if err != nil {
		panic(err)
	}

	// run application tests
	stderr, err := runner.Stderr(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Repeat("=", 79))
	fmt.Println()
	fmt.Println("Output for Job: ")
	fmt.Println()
	fmt.Println("[stdout]")
	fmt.Println()
	fmt.Println(stdout)
	fmt.Println()
	fmt.Println(strings.Repeat("-", 79))
	fmt.Println()
	fmt.Println("[stderr]")
	fmt.Println()
	fmt.Println(stderr)
	fmt.Println()
	fmt.Println(strings.Repeat("=", 79))
}
