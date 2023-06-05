package app

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Onelvay/HL-architecture/config"
	"github.com/Onelvay/HL-architecture/internal/handler"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/internal/service"
	"go.uber.org/zap"
)

func Run() {
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
		sugar.Fatalf("failed to init config:  %s", err)
	}

	repo, db, err := repository.New(repository.PostgresRepository())
	if err != nil {
		sugar.Fatalf("failed to start reoi: %s", err)
	}
	defer db.Close()

	d := service.Dependencies{
		CourseRepo: repo.Course,
		AuthRepo:   repo.Auth,
		OrderRepo:  repo.Order,
		UserRepo:   repo.User,
	}
	s := service.New(d)

	if err != nil {
		sugar.Errorf("failed to init repository %s", err)
	}

	server := handler.NewServer(cfg, s)
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
