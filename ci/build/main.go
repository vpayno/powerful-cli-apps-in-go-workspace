// Dagger CI Tooling - Build
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strings"

	"dagger.io/dagger"
)

func main() {
	// define build matrix
	geese := []string{"linux", "darwin", "windows"}
	goarches := []string{
		"386",
		"amd64",
		"arm64",
		"mips",
		"mips64",
		"mips64le",
		"mipsle",
		"ppc64",
		"ppc64le",
	}

	goosArchMap := map[string][]string{}
	goosArchMap["linux"] = []string{
		"386",
		"amd64",
		"arm",
		"mips",
		"mipsle",
		"mips64le",
		"ppc64",
		"ppc64le",
	}
	goosArchMap["darwin"] = []string{"amd64", "arm64"}
	goosArchMap["windows"] = []string{"amd64", "arm64"}

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
	repo := client.Host().Directory(".")

	// create empty directory to put build outputs
	outputs := client.Directory()

	golang := client.Container().
		// get golang image
		From("golang:latest").
		// mount source code into golang image
		WithDirectory("/repo", repo).
		WithWorkdir("/repo").
		WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golangVer)).
		WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
		WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golangVer)).
		WithEnvVariable("GOCACHE", "/go/build-cache").
		WithExec([]string{"rm", "-rf", "./build"}).
		WithExec([]string{"pwd"}).
		WithExec([]string{"ls", "-lvh"})

	for _, goos := range geese {
		for _, goarch := range goarches {
			if !slices.Contains(goosArchMap[goos], goarch) {
				continue
			}

			// create a directory for each OS and architecture
			outputPath := fmt.Sprintf("build/%s/%s/", goos, goarch)

			build := golang.
				// set GOARCH and GOOS in the build environment
				WithEnvVariable("GOOS", goos).
				WithEnvVariable("GOARCH", goarch)

			fmt.Printf(
				"GOOS: %s\tGOARCH: %s\n",
				goos,
				goarch,
			)

			entries, err := os.ReadDir("./cmd/")
			if err != nil {
				log.Fatal(err)
			}

			for _, entry := range entries {
				if entry.IsDir() {
					mainFile := "./cmd/" + entry.Name() + "/main.go"
					build = build.WithExec(
						[]string{
							"go",
							"build",
							"-o",
							outputPath + entry.Name(),
							mainFile,
						},
					)
				}
			}

			// add build to outputs
			outputs = outputs.WithDirectory(outputPath, build.Directory(outputPath))
		}
	}

	// write build artifacts to host
	out, err := outputs.Export(ctx, ".")
	if err != nil {
		panic(err)
	}

	if out != "" {
		panic("did not export files")
	}
}
