package app

import (
	"flag"
	"fmt"
	"github.com/Onelvay/HL-architecture/config"
	"github.com/Onelvay/HL-architecture/internal/api/rest"
	"github.com/Onelvay/HL-architecture/pkg/database"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	database.New()
	if err := config.ParseYaml(); err != nil {
		log.Fatalf(err.Error())
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	sugar := logger.Sugar()

	cfg, err := config.New()
	if err != nil {
		sugar.Errorf("failed to init config:  %s", err)
	}
	server := rest.NewServer(cfg)
	go func() {
		if err = server.Run(); err != nil {
			sugar.Error(err)
			return
		}
	}()
	sugar.Infof("listen on port %s", cfg.Http.Port)

	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	fmt.Println("Server was successful shutdown.")

}
