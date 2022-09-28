package grpc

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/env"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

func New(lifecycle *app.Lifecycle, logger *logrus.Logger) *grpc.Server {
	addr := env.StringWithDefault("GRPC_ADDR", "tcp://0.0.0.0:9000")
	uri, err := url.Parse(addr)
	if err != nil {
		logger.Fatal(err)
	}

	listener, err := net.Listen(uri.Scheme, uri.Host)
	if err != nil {
		logger.Fatal(err)
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
				logger.Fatal(err)
			}
		}()
		logger.WithFields(logrus.Fields{
			"address": fmt.Sprintf("%s://%s", uri.Scheme, uri.Host),
		}).Info("gRPC server listening")
		return nil
	})

	lifecycle.OnStop(func(ctx context.Context) error {
		server.GracefulStop()
		return nil
	})

	return server
}
