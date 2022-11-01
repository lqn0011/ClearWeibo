package main

import (
	"context"
	"os"
	"time"

	"clean-wb/app/config"
	"clean-wb/app/svc"
	"clean-wb/basic/conf"
	"clean-wb/basic/log"

	"github.com/urfave/cli/v2"
)

var (
	configFile string
	cfg        *config.Config
)

func Cfg() *config.Config {
	if cfg == nil {
		cfg = &config.Config{}
		conf.MustLoadConfig(configFile, &cfg)
	}
	return cfg
}

func main() {
	app := &cli.App{
		Name:  "clean-wb",
		Usage: "清空微博脚本",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "etc/clean-wb.yaml",
				Usage:       "the config file",
				Destination: &configFile,
			},
		},
		Action: func(c *cli.Context) error {
			return CleanAll(c.Context)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Err(err).Msg("start failed.")
		log.Fatal()
	}
}

func CleanAll(ctx context.Context) (err error) {
	scfg := Cfg()
	log.Init(scfg.Log)
	svcCtx := svc.NewServiceContext(*scfg)

	defer func() {
		svcCtx.Close()
		log.Close()
		time.Sleep(300 * time.Millisecond)
	}()

	log.Info().Msg("Start cleaning.")
	svcCtx.CleanAll(ctx)
	log.Info().Msg("End cleaning. Do some cleaning.")

	return
}
