package grpcserver

import (
	"context"

	"github.com/Marcelixoo/go_specs_greet/domain/interactions"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, r *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(r.Name)}, nil
}
