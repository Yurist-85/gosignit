package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/sarulabs/di/v2"
	"github.com/sirupsen/logrus"

	"github.com/yurist-85/gosignit/internal/api"
	"github.com/yurist-85/gosignit/internal/config"
	"github.com/yurist-85/gosignit/internal/handlers/coins"
	"github.com/yurist-85/gosignit/internal/handlers/txs"
	"github.com/yurist-85/gosignit/internal/services/clock"
	"github.com/yurist-85/gosignit/internal/services/trustwallet"
)

const (
	appTag              = "tx-api"
	defaultReadTimeout  = 10 * time.Second
	defaultWriteTimeout = 10 * time.Second
)

// Injected during building process.
var (
	gitCommit string
	buildDate string
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	logrus.WithFields(logrus.Fields{"revision": gitCommit, "build_date": buildDate}).Println()

	ctn, err := newApp()
	if err != nil {
		logrus.WithError(err).Fatalln("initialize app container")
	}

	defer func() {
		if err := ctn.DeleteWithSubContainers(); err != nil {
			logrus.WithError(err).Errorln("destroy app container")
		}
	}()

	handler := ctn.Get(api.DefinitionName).(api.HandlerInterface)
	e := handler.GetEcho()
	cfg := ctn.Get(config.DefinitionName).(config.ConfigInterface)
	logrus.WithField("host", cfg.HostAndPort()).Info("starting HTTP server...")

	go func() {
		s := &http.Server{
			Addr:         cfg.HostAndPort(),
			ReadTimeout:  defaultReadTimeout,
			WriteTimeout: defaultWriteTimeout,
		}
		s.SetKeepAlivesEnabled(false)

		if err := e.StartServer(s); err != nil {
			logrus.WithFields(logrus.Fields{logrus.ErrorKey: err, "host": cfg.Host()}).Fatalln("start http server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	logrus.WithField("signal", <-quit).Info("graceful shutdown by sigterm")

	// nolint:gomnd // intent usage
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

// newApp creates DI container with declared dependencies.
func newApp() (di.Container, error) {
	builder, err := di.NewBuilder()
	if err != nil {
		return nil, errors.Wrap(err, "new di builder")
	}

	if err := builder.Add(
		di.Def{
			Name: "app",
			Build: func(ctn di.Container) (interface{}, error) {
				return appTag, nil
			},
		},
		config.Definition,
		// API Handlers
		api.Definition,
		api.DefinitionEcho,
		coins.Definition,
		txs.Definition,
		// Services
		clock.Definition,
		trustwallet.Definition,
	); err != nil {
		return nil, errors.Wrap(err, "add dependencies to container")
	}

	return builder.Build(), nil
}
