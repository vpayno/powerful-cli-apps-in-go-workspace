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
	golang_ver := runtime.Version()
	golang_ver = strings.Replace(golang_ver, "go", "", -1)
	fmt.Println("Golang Ver: " + golang_ver)

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
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golang_ver)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golang_ver)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithExec([]string{"rm", "-rf", "./build"})

	for _, goos := range geese {
		for _, goarch := range goarches {
			// create a directory for each OS and architecture
			output_path := fmt.Sprintf("build/%s/%s/", goos, goarch)

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
					main_file := "./cmd/" + entry.Name() + "/main.go"
					build = build.WithExec(
						[]string{"go", "build", "-o", output_path + entry.Name(), main_file},
					)
				}
			}

			// add build to outputs
			outputs = outputs.WithDirectory(output_path, build.Directory(output_path))
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
