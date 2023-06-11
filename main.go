package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpczerolog "github.com/cheapRoc/grpc-zerolog"
	_ "github.com/jnewmano/grpc-json-proxy/codec"
	"github.com/jukeizu/birthday/api/protobuf-spec/birthdaypb"
	"github.com/jukeizu/birthday/treediagram"
	"github.com/oklog/run"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	"github.com/shawntoffel/gossage"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

var Version = ""

var (
	flagMigrate = false
	flagVersion = false
	flagServer  = false
	flagHandler = false

	grpcPort       = "50052"
	httpPort       = "10002"
	dbAddress      = "root@localhost:26257"
	serviceAddress = "localhost:" + grpcPort
)

func parseConfig() {
	flag.StringVar(&grpcPort, "grpc.port", grpcPort, "grpc port for server")
	flag.StringVar(&httpPort, "http.port", httpPort, "http port for handler")
	flag.StringVar(&dbAddress, "db", dbAddress, "Database connection address")
	flag.StringVar(&serviceAddress, "service.addr", serviceAddress, "address of service if not local")
	flag.BoolVar(&flagServer, "server", false, "Run as server")
	flag.BoolVar(&flagHandler, "handler", false, "Run as handler")
	flag.BoolVar(&flagMigrate, "migrate", false, "Run db migrations")
	flag.BoolVar(&flagVersion, "v", false, "version")

	flag.Parse()
}

func main() {
	parseConfig()

	if flagVersion {
		fmt.Println(Version)
		os.Exit(0)
	}

	logger := zerolog.New(os.Stdout).With().Timestamp().
		Str("instance", xid.New().String()).
		Str("component", "birthday").
		Str("version", Version).
		Logger()

	grpcLoggerV2 := grpczerolog.New(logger.With().Str("transport", "grpc").Logger())
	grpclog.SetLoggerV2(grpcLoggerV2)

	if !flagServer && !flagHandler {
		flagServer = true
		flagHandler = true
	}

	if flagMigrate {
		birthdayRepository, err := NewRepository(dbAddress)
		if err != nil {
			logger.Error().Err(err).Caller().Msg("could not create birthday repository")
			os.Exit(1)
		}

		gossage.Logger = func(format string, a ...interface{}) {
			msg := fmt.Sprintf(format, a...)
			logger.Info().Str("component", "migrator").Msg(msg)
		}

		err = birthdayRepository.Migrate()
		if err != nil {
			logger.Error().Err(err).Caller().Msg("could not migrate birthday repository")
			os.Exit(1)
		}
	}

	g := run.Group{}

	if flagServer {
		birthdayRepository, err := NewRepository(dbAddress)
		if err != nil {
			logger.Error().Err(err).Caller().Msg("could not create birthday repository")
			os.Exit(1)
		}

		grpcServer := newGrpcServer(logger)

		reflection.Register(grpcServer)
		birthdayServer := NewServer(logger, birthdayRepository, grpcServer)
		birthdaypb.RegisterBirthdayServiceServer(grpcServer, birthdayServer)
		grpcAddr := ":" + grpcPort

		g.Add(func() error {
			return birthdayServer.Start(grpcAddr)
		}, func(error) {
			birthdayServer.Stop()
		})
	}

	if flagHandler {
		clientConn, err := grpc.Dial(serviceAddress, grpc.WithInsecure(),
			grpc.WithKeepaliveParams(
				keepalive.ClientParameters{
					Time:                30 * time.Second,
					Timeout:             10 * time.Second,
					PermitWithoutStream: true,
				},
			),
		)
		if err != nil {
			logger.Error().Err(err).Str("serviceAddress", serviceAddress).Msg("could not dial service address")
			os.Exit(1)
		}

		httpAddr := ":" + httpPort

		client := birthdaypb.NewBirthdayServiceClient(clientConn)
		handler := treediagram.NewHandler(logger, client, httpAddr)

		g.Add(func() error {
			return handler.Start()
		}, func(error) {
			err := handler.Stop()
			if err != nil {
				logger.Error().Err(err).Caller().Msg("couldn't stop handler")
			}
		})
	}

	cancel := make(chan struct{})
	g.Add(func() error {
		return interrupt(cancel)
	}, func(error) {
		close(cancel)
	})

	logger.Info().Err(g.Run()).Msg("stopped")
}

func interrupt(cancel <-chan struct{}) error {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-cancel:
		return errors.New("stopping")
	case sig := <-c:
		return fmt.Errorf("%s", sig)
	}
}

func newGrpcServer(logger zerolog.Logger) *grpc.Server {
	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    5 * time.Minute,
				Timeout: 10 * time.Second,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             5 * time.Second,
				PermitWithoutStream: true,
			},
		),
		LoggingInterceptor(logger),
	)

	return grpcServer
}
