// Dagger CI Tooling - Build
package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

	"dagger.io/dagger"
)

func main() {
	// define build matrix
	geese := []string{"linux", "darwin", "windows"}
	goarches := []string{"amd64", "arm64"}

	ctx := context.Background()
	// initialize dagger client
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}

	// go1.21.5
	golangVer := runtime.Version()
	golangVer = strings.ReplaceAll(golangVer, "go", "")
	fmt.Println("Golang Ver: " + golangVer)

	// get reference to the local project
	src := client.Host().Directory(".")

	// create empty directory to put build outputs
	outputs := client.Directory()

	golang := client.Container().
		// get golang image
		From("golang:latest").
		// mount source code into golang image
		WithDirectory("/src", src).
		WithWorkdir("/src").
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golangVer)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golangVer)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithExec([]string{"rm", "-rf", "./build"})

	for _, goos := range geese {
		for _, goarch := range goarches {
			// create a directory for each OS and architecture
			outputPath := fmt.Sprintf("build/%s/%s/", goos, goarch)

			build := golang.
				// set GOARCH and GOOS in the build environment
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch)

			entries, err := ioutil.ReadDir("./cmd/")
			if err != nil {
				log.Fatal(err)
			}

			for _, entry := range entries {
				if entry.IsDir() {
					mainFile := "./cmd/" + entry.Name() + "/main.go"
					build = build.WithExec(
						[]string{"go", "build", "-o", outputPath + entry.Name(), mainFile},
					)
				}
			}

			// add build to outputs
			outputs = outputs.WithDirectory(outputPath, build.Directory(outputPath))
		}
	}

	// write build artifacts to host
	ok, err := outputs.Export(ctx, ".")
	if err != nil {
		panic(err)
	}

	if !ok {
		panic("did not export files")
	}
}
