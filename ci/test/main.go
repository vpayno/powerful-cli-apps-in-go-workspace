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
	golang_ver := runtime.Version()
	golang_ver = strings.Replace(golang_ver, "go", "", -1)
	fmt.Println("Golang Ver: " + golang_ver)

	// use a node:16-slim container
	// mount the source code directory on the host
	// at /src in the container
	source := client.Container().
		From("vpayno/ci-generic-debian:latest").
		WithDirectory("/src", client.Host().Directory(".", dagger.HostDirectoryOpts{
			Exclude: []string{"ci/", "build/"},
		}))

	source = source.WithWorkdir("/src").
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golang_ver)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golang_ver)).
		WithEnvVariable("GOCACHE", "/go/build-cache")

	// set the working directory in the container
	// install application dependencies
	runner := source.WithExec([]string{"./scripts/go-test-with-coverage"})

	// run application tests
	out, err := runner.Stderr(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}
