package server

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type grpcSrv struct {
	srv    *grpc.Server
	port   int
	logger *logger.Logger
}

func newGRPC(port int) *grpcSrv {
	l := logger.NewLogger("grpc-server")
	l.Info("grpc server initializing...")

	srv := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_zap.StreamServerInterceptor(l.GetRawLogger()),
				errorHandlingStreamIntercepter(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_zap.UnaryServerInterceptor(l.GetRawLogger()),
				errorHandlingUnaryIntercepter(),
			),
		),
	)

	grpc_health_v1.RegisterHealthServer(srv, health.NewServer())

	reflection.Register(srv)

	l.Info("grpc server initialized")

	return &grpcSrv{
		srv:    srv,
		port:   port,
		logger: l,
	}
}

func (srv *grpcSrv) run(c context.Context) error {
	srv.logger.Info("grpc server running...")

	errCh := make(chan error, 1)

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", srv.port))
		if err != nil {
			errCh <- err

			return
		}

		if err := srv.srv.Serve(lis); err != nil {
			errCh <- err

			return
		}
	}()

	srv.logger.Info("grpc server run", zap.Int("port", srv.port))

	select {
	case err := <-errCh:
		srv.logger.Error("catch error", zap.Error(err))

		return err
	case <-c.Done():
		srv.logger.Info("grpc server stopping...")
		srv.srv.GracefulStop()
		srv.logger.Info("grpc server stopped")

		return nil
	}
}

func errorHandlingUnaryIntercepter() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err == nil {
			return resp, nil
		}

		var (
			iae = domain.InvalidArgumentError{}
			nfe = domain.NotFoundError{}
			ce  = domain.ConflictError{}
		)

		switch {
		case errors.As(err, &iae):
			return resp, status.New(codes.InvalidArgument, err.Error()).Err()
		case errors.As(err, &nfe):
			return resp, status.New(codes.NotFound, err.Error()).Err()
		case errors.As(err, &ce):
			return resp, status.New(codes.AlreadyExists, err.Error()).Err()
		default:
			return resp, status.New(codes.Unknown, fmt.Sprintf("Unknown: %v", err)).Err()
		}
	}
}

func errorHandlingStreamIntercepter() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		err := handler(srv, ss)
		if err == nil {
			return nil
		}

		var (
			iae = domain.InvalidArgumentError{}
			nfe = domain.NotFoundError{}
			ce  = domain.ConflictError{}
		)

		switch {
		case errors.As(err, &iae):
			return status.New(codes.InvalidArgument, err.Error()).Err()
		case errors.As(err, &nfe):
			return status.New(codes.NotFound, err.Error()).Err()
		case errors.As(err, &ce):
			return status.New(codes.AlreadyExists, err.Error()).Err()
		default:
			return status.New(codes.Unknown, fmt.Sprintf("Unknown: %s", err)).Err()
		}
	}
}
