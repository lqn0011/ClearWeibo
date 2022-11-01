package config

import (
	"clean-wb/basic/log"
	"clean-wb/basic/rhttp"
)

type Config struct {
	Resty *rhttp.Config
	Log   *log.Config

	Uid    int64
	Cookie string
	XXsrfToken string
	Debug bool
}
