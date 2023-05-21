package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-co-op/gocron"
	ipsw_go "ipsw-go/ipsw-go"
	"ipsw-go/logger"
	"log"
	"strings"
	"time"
)

type config struct {
	Workers     int      `env:"Workers"`
	Identifiers []string `env:"Identifiers,required" envSeparator:":"`
	TZ          string   `env:"TZ,required"`
	CheckAt     string   `env:"CheckAt,required"`
	Mode        string   `env:"Mode"`
}

func main() {

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	logger.Init(cfg.Mode)

	logger.Info("ipsw-go developed by 0xJacky")
	logger.Info("Data source: Betahub.cn")

	loc, err := time.LoadLocation(cfg.TZ)

	checkAt := strings.ReplaceAll(cfg.CheckAt, ",", ";")

	logger.Info("Timezone", loc.String())
	logger.Info("Check at", checkAt)

	s := gocron.NewScheduler(loc)

	_, err = s.Every(1).Day().At(checkAt).StartImmediately().Do(ipsw_go.Worker, cfg.Workers, cfg.Identifiers)

	if err != nil {
		logger.Fatal("[Error] create cron job", err)
	}

	s.StartBlocking()
}
