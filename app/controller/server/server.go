package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/g-chicken/mah-jong/app/domain"
	"github.com/g-chicken/mah-jong/app/logger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	grpcSrv *grpcSrv
	logger  *logger.Logger
}

func NewServer(config *domain.Config) *Server {
	l := logger.NewLogger("server")
	l.Info("server initializing...")

	grpcSrv := newGRPC(config.GetGRPCPort())

	l.Info("server initialized")

	return &Server{
		grpcSrv: grpcSrv,
		logger:  l,
	}
}

func (srv *Server) Run() error {
	srv.logger.Info("server running...")

	eg, c := errgroup.WithContext(context.Background())

	eg.Go(func() error { return srv.grpcSrv.run(c) })
	eg.Go(func() error { return srv.receiveSignal(c) })

	srv.logger.Info("server run")

	if err := eg.Wait(); err != nil {
		srv.logger.Error("catch error", zap.Error(err))

		return err
	}

	return nil
}

func (srv *Server) receiveSignal(c context.Context) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	select {
	case s := <-sig:
		srv.logger.Info("receive signal", zap.String("signal", s.String()))

		return fmt.Errorf("receive signal (signal = %s)", s.String()) // nolint:goerr113
	case <-c.Done():
		return nil
	}
}
