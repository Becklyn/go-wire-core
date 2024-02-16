package grpc

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/fraym/golog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func New(lifecycle *app.Lifecycle, logger golog.Logger) *grpc.Server {
	addr := env.StringWithDefault("GRPC_ADDR", "tcp://0.0.0.0:9000")
	uri, err := url.Parse(addr)
	if err != nil {
		logger.Fatal().WithError(err).Write()
	}

	listener, err := net.Listen(uri.Scheme, uri.Host)
	if err != nil {
		logger.Fatal().WithError(err).Write()
	}

	keepaliveOptions := grpc.KeepaliveParams(keepalive.ServerParameters{
		Time:    time.Minute,
		Timeout: 3 * time.Second,
	})

	keepaliveEnforcementOptions := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		MinTime:             30 * time.Second,
		PermitWithoutStream: true,
	})

	server := grpc.NewServer(
		keepaliveOptions,
		keepaliveEnforcementOptions,
	)

	lifecycle.OnStart(func(ctx context.Context) error {
		go func() {
			reflection.Register(server)

			if err := server.Serve(listener); err != nil {
				logger.Fatal().WithError(err).Write()
			}
		}()
		logger.Info().WithField("address", fmt.Sprintf("%s://%s", uri.Scheme, uri.Host)).Write("gRPC server listening")
		return nil
	})

	lifecycle.OnStopLast(func(ctx context.Context) error {
		stopped := make(chan bool)
		go func() {
			server.GracefulStop()
			close(stopped)
		}()

		t := time.NewTimer(5 * time.Second)
		select {
		case <-t.C:
			server.Stop()
		case <-stopped:
			t.Stop()
		}

		return nil
	})

	return server
}
