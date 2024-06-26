package main

import (
	"flag"
	"os"

	"github.com/ViBiOh/absto/pkg/absto"
	"github.com/ViBiOh/flags"
	"github.com/ViBiOh/httputils/v4/pkg/health"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	"github.com/ViBiOh/httputils/v4/pkg/pprof"
	"github.com/ViBiOh/httputils/v4/pkg/telemetry"
	"github.com/ViBiOh/netatmo/v2/pkg/netatmo"
)

type configuration struct {
	logger    *logger.Config
	telemetry *telemetry.Config
	pprof     *pprof.Config
	health    *health.Config

	netatmo      *netatmo.Config
	absto        *absto.Config
	cipherSecret *string
}

func newConfig() configuration {
	fs := flag.NewFlagSet("netatmo", flag.ExitOnError)
	fs.Usage = flags.Usage(fs)

	config := configuration{
		logger:    logger.Flags(fs, "logger"),
		telemetry: telemetry.Flags(fs, "telemetry"),
		pprof:     pprof.Flags(fs, "pprof"),
		health:    health.Flags(fs, ""),

		netatmo:      netatmo.Flags(fs, ""),
		absto:        absto.Flags(fs, "storage"),
		cipherSecret: flags.New("CipherSecret", "Secret for ciphering token, 32 characters").DocPrefix("secret").String(fs, "", nil),
	}

	_ = fs.Parse(os.Args[1:])

	return config
}
