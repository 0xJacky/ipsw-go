package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-co-op/gocron"
	ipsw_go "ipsw-go/ipsw-go"
	"log"
	"strings"
	"time"
)

type config struct {
	Workers     int      `env:"Workers"`
	Identifiers []string `env:"Identifiers,required" envSeparator:":"`
	TZ          string   `env:"TZ,required"`
	CheckAt     string   `env:"CheckAt,required"`
}

func main() {
	log.Println("[Info] ipsw-go developed by 0xJacky")
	log.Println("[Info] Data source: Betahub.cn")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("%+v\n", err)
	}

	loc, err := time.LoadLocation(cfg.TZ)

	checkAt := strings.ReplaceAll(cfg.CheckAt, ",", ";")

	log.Println("[Info]", "Timezone", loc.String())
	log.Println("[Info]", "Check at", checkAt)

	s := gocron.NewScheduler(loc)

	_, err = s.Every(1).Day().At(checkAt).Do(ipsw_go.Worker, cfg.Workers, cfg.Identifiers)

	if err != nil {
		log.Fatalln("[Error] create cron job", err)
	}

	s.StartBlocking()
}
