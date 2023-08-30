package main_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/Marcelixoo/go_specs_greet/adapters"
	"github.com/Marcelixoo/go_specs_greet/adapters/httpserver"
	"github.com/Marcelixoo/go_specs_greet/specifications"
	"github.com/alecthomas/assert/v2"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestGreeterServer(t *testing.T) {
	var (
		port           = "8080"
		dockerFilePath = "./cmd/httpserver/Dockerfile"
		baseURL        = fmt.Sprintf("http://localhost:%s", port)
		driver         = httpserver.Driver{
			BaseURL: baseURL,
			Client:  &http.Client{Timeout: 1 * time.Second},
		}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, driver)
}

func setupTestContainers(t testing.TB) {
	t.Helper()

	req := testcontainers.ContainerRequest{
		FromDockerfile: testcontainers.FromDockerfile{
			Context:       "../../.",
			Dockerfile:    "./cmd/httpserver/Dockerfile",
			PrintBuildLog: true, // set to false for less spam
		},
		ExposedPorts: []string{"8080:8080"},
		WaitingFor:   wait.ForHTTP("/").WithPort("8080"),
	}
	ctx := context.Background()

	container, err := testcontainers.GenericContainer(
		ctx,
		testcontainers.GenericContainerRequest{
			ContainerRequest: req,
			Started:          true,
		},
	)
	assert.NoError(t, err)

	t.Cleanup(func() {
		assert.NoError(t, container.Terminate(ctx))
	})
}
