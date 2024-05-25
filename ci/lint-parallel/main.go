// Dagger CI Tooling - Linting
package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"dagger.io/dagger"
)

func main() {
	cpuCount := runtime.NumCPU()
	runtime.GOMAXPROCS(cpuCount) // defaults to runtime.NumCPU

	fmt.Printf("Runner CPU/Core count: %d\n", cpuCount)

	wg := &sync.WaitGroup{}

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

	jobEntries := [][]string{
		{"pwd"},
		{"ls", "-lvh"},
		{"gocritic", "check", "-enableAll", "./..."},
		{"gocyclo", "-over", "15", "."},
		{"golangci-lint", "run", "./..."},
		{"gosec", "./..."},
		{"govulncheck", "./..."},
		{"ineffassign", "./..."},
		{"revive", "./..."},
		{"staticcheck", "./..."},
	}

	// wg.Add(len(jobEntries))

	messages := make(chan string)

	fmt.Println(strings.Repeat("=", 79))
	fmt.Println()

	for _, jobEntry := range jobEntries {
		fmt.Println(strings.Repeat("-", 79))
		fmt.Println("Starting job: ", jobEntry)
		fmt.Println()

		go func(job []string) {
			defer wg.Done()

			wg.Add(1)

			// mount the source code directory on the host at /repo in the container
			source := client.Container().
				From("vpayno/ci-generic-debian:latest").
				WithDirectory("/repo", client.Host().Directory(".", dagger.HostDirectoryOpts{
					Exclude: []string{"build/"},
				}))

			// set the working directory in the container
			// install application dependencies
			source = source.WithWorkdir("/repo").
				WithMountedCache("/go/pkg/mod", client.CacheVolume("go-mod-"+golangVer)).
				WithEnvVariable("GOMODCACHE", "/go/pkg/mod").
				WithMountedCache("/go/build-cache", client.CacheVolume("go-build"+golangVer)).
				WithEnvVariable("GOCACHE", "/go/build-cache")

			runner := source.WithExec(job)

			// run application tests
			stdout, err := runner.Stdout(ctx)
			if err != nil {
				panic(err)
			}

			stderr, err := runner.Stderr(ctx)
			if err != nil {
				panic(err)
			}

			message := "Output for Job: "
			message += fmt.Sprintln(job)
			message += "\n[stdout]\n"
			message += stdout
			message += "\n"
			message += strings.Repeat("-", 79)
			message += "\n[stderr]\n"
			message += stderr

			messages <- message
		}(jobEntry)
	}

	fmt.Println()
	fmt.Println(strings.Repeat("=", 79))
	fmt.Println()

	go func() {
		fmt.Println(strings.Repeat("=", 79))
		fmt.Println()
		fmt.Println("Job Output")

		for message := range messages {
			fmt.Println()
			fmt.Println(strings.Repeat("-", 79))
			fmt.Println()
			fmt.Println(message)
			fmt.Println()
		}

		fmt.Println(strings.Repeat("=", 79))
	}()

	wg.Wait()
}
