package grpcserver

import (
	"context"
	sync "sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Driver struct {
	Addr string

	once   sync.Once
	conn   *grpc.ClientConn
	client GreeterClient
}

func (d *Driver) Greet(name string) (string, error) {
	client, err := d.getGreeterClient()
	if err != nil {
		return "", err
	}

	greeting, err := client.Greet(context.Background(), &GreetRequest{
		Name: name,
	})
	if err != nil {
		return "", err
	}

	return greeting.Message, nil
}

func (d *Driver) getGreeterClient() (GreeterClient, error) {
	var err error

	d.once.Do(func() {
		d.conn, err = grpc.Dial(d.Addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		d.client = NewGreeterClient(d.conn)
	})

	return d.client, err
}
