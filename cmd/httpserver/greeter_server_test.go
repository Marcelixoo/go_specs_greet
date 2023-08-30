package main_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/Marcelixoo/go_specs_greet/adapters"
	"github.com/Marcelixoo/go_specs_greet/adapters/httpserver"
	"github.com/Marcelixoo/go_specs_greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

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
