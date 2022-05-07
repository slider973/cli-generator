package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/slider973/cli-generator/config"
	"github.com/slider973/cli-generator/pkg"
	flag "github.com/spf13/pflag"
)

var (
	cfg        *config.Config = new(config.Config)
	cout       *config.Config
	cfgDefault *config.Config = config.DefaultConfig
	logger     hclog.Logger   = hclog.New(hclog.DefaultOptions)
	version    string         = "unversioned"
)

func init() {
	var err error

	flag.StringVar(&cfg.ConfigFile, "config", "", "CLI Generator configuration file path")
	flag.StringVar(&cfg.LogOptions.Level, "log.level", cfgDefault.LogOptions.Level, "Log level values allowed [trace, debug, info, warn, error, fatal]")
	flag.StringVar(&cfg.LogOptions.Format, "log.fmt", cfgDefault.LogOptions.Format, "Log format values allowed [logfmt, json]")
	flag.Parse()

	if cfg.ConfigFile != "" {
		if cout, err = cfg.Reload(logger); err != nil {
			fmt.Fprintf(os.Stderr, "impossible to load config file: %v\n", err)
			os.Exit(1)
		}
	} else {
		flag.PrintDefaults()
	}
}

func main() {
	var (
		log    hclog.Logger = cfg.LogOptions.LogFlagParse("satellite")
		client *pkg.Client  = pkg.NewClient(log)
	)

	log.With("version", version).Info("Start cli-generator")
	client.Start(cout.Generate)
}
