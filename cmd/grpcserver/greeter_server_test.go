package main_test

import (
	"fmt"
	"testing"

	"github.com/Marcelixoo/go_specs_greet/adapters"
	"github.com/Marcelixoo/go_specs_greet/adapters/grpcserver"
	"github.com/Marcelixoo/go_specs_greet/specifications"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	var (
		port           = "50051"
		dockerFilePath = "./cmd/grpcserver/Dockerfile"
		driver         = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, dockerFilePath)
	specifications.GreetSpecification(t, &driver)
}
