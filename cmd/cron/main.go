package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"

	myCron "github.com/avraam311/event-booker/internal/infra/cron"
	repository "github.com/avraam311/event-booker/internal/repository/events"
	service "github.com/avraam311/event-booker/internal/service/events"

	"github.com/wb-go/wbf/config"
	"github.com/wb-go/wbf/dbpg"
	"github.com/wb-go/wbf/zlog"

	"github.com/robfig/cron/v3"
)

const (
	configFilePath = "config/local.yaml"
	envFilePath    = ".env"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	zlog.Init()
	cfg := config.New()
	if err := cfg.LoadEnvFiles(envFilePath); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("failed to load env file")
	}
	cfg.EnableEnv("")
	if err := cfg.LoadConfigFiles(configFilePath); err != nil {
		zlog.Logger.Fatal().Err(err).Msg("failed to load config file")
	}

	opts := &dbpg.Options{
		MaxOpenConns:    cfg.GetInt("db.max_open_conns"),
		MaxIdleConns:    cfg.GetInt("db.max_idle_conns"),
		ConnMaxLifetime: cfg.GetDuration("db.conn_max_lifetime"),
	}
	slavesDNSs := []string{}
	masterDNS := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.GetString("DB_USER"), cfg.GetString("DB_PASSWORD"),
		cfg.GetString("DB_HOST"), cfg.GetString("DB_PORT"),
		cfg.GetString("DB_NAME"), cfg.GetString("DB_SSL_MODE"),
	)
	db, err := dbpg.New(masterDNS, slavesDNSs, opts)
	if err != nil {
		zlog.Logger.Fatal().Err(err).Msg("failed to connect to database")
	}

	repo := repository.NewRepository(db)
	srvc := service.NewService(repo)
	crn := cron.New()
	myCrn := myCron.New(crn)
	myCrn.AddTask(cfg.GetString("cron.spec"), func() { srvc.CleanExpiredBooks(ctx) })
	go myCrn.Start()

	zlog.Logger.Info().Msg("cron is running")

	<-ctx.Done()
	zlog.Logger.Info().Msg("shutdown signal received")

	myCrn.Stop()

	if err := db.Master.Close(); err != nil {
		zlog.Logger.Printf("failed to close master DB: %v", err)
	}
	for i, s := range db.Slaves {
		if err := s.Close(); err != nil {
			zlog.Logger.Printf("failed to close slave DB %d: %v", i, err)
		}
	}
}
