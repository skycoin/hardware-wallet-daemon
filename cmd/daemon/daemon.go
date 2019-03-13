package main

import (
	"flag"
	"os"

	"github.com/skycoin/skycoin/src/util/logging"

	"github.com/skycoin/hardware-wallet-daemon/src/daemon"
)

var (
	logger = logging.MustGetLogger("hw-daemon")

	config = daemon.NewConfig(
		6430,
		"$HOME/.skycoin")
)

func init() {
	config.RegisterFlags()
}

func main() {
	flag.Parse()

	d := daemon.NewDaemon(config, logger)

	// parse config values
	if err := d.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	if err := d.Run(); err != nil {
		os.Exit(1)
	}
}
