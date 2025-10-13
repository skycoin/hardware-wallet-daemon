package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/skycoin/hardware-wallet-daemon/src/api"
	"github.com/skycoin/hardware-wallet-daemon/src/daemon"
	"github.com/skycoin/skycoin/src/util/logging"
)

var (
	Version = "0.1.0"
	Commit  = ""
	Branch  = ""

	appConfig = daemon.NewAppConfig(9510, "$HOME/.skycoin")
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hardware-wallet-daemon",
	Short: "Hardware wallet daemon for Skycoin",
	Long:  `A daemon that provides an HTTP API for interfacing with Skycoin hardware wallets`,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logging.MustGetLogger("hw-daemon")

		d := daemon.NewDaemon(daemon.Config{
			App: appConfig,
			Build: api.BuildInfo{
				Version: Version,
				Commit:  Commit,
				Branch:  Branch,
			},
		}, logger)

		if err := d.ParseConfig(); err != nil {
			logger.Error(err)
			os.Exit(1)
		}

		if err := d.Run(); err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.Flags().IntVarP(&appConfig.WebInterfacePort, "web-interface-port", "p", appConfig.WebInterfacePort, "port to serve web interface on")
	rootCmd.Flags().StringVarP(&appConfig.WebInterfaceAddr, "web-interface-addr", "a", appConfig.WebInterfaceAddr, "addr to serve web interface on")
	rootCmd.Flags().BoolVar(&appConfig.EnableCSRF, "enable-csrf", appConfig.EnableCSRF, "enable CSRF check")
	rootCmd.Flags().BoolVar(&appConfig.DisableHeaderCheck, "disable-header-check", appConfig.DisableHeaderCheck, "disables the host, origin and referer header checks")
	rootCmd.Flags().StringVar(&appConfig.HostWhitelist, "host-whitelist", appConfig.HostWhitelist, "Hostnames to whitelist in the Host header check. Only applies when the web interface is bound to localhost.")

	rootCmd.Flags().BoolVar(&appConfig.ColorLog, "color-log", appConfig.ColorLog, "Add terminal colors to log output")
	rootCmd.Flags().StringVarP(&appConfig.LogLevel, "log-level", "l", appConfig.LogLevel, "Choices are: debug, info, warn, error, fatal, panic")
	rootCmd.Flags().BoolVar(&appConfig.LogToFile, "logtofile", appConfig.LogToFile, "log to file")

	rootCmd.Flags().BoolVar(&appConfig.ProfileCPU, "profile-cpu", appConfig.ProfileCPU, "enable cpu profiling")
	rootCmd.Flags().StringVar(&appConfig.ProfileCPUFile, "profile-cpu-file", appConfig.ProfileCPUFile, "where to write the cpu profile file")
	rootCmd.Flags().BoolVar(&appConfig.HTTPProf, "http-prof", appConfig.HTTPProf, "run the HTTP profiling interface")
	rootCmd.Flags().StringVar(&appConfig.HTTPProfHost, "http-prof-host", appConfig.HTTPProfHost, "hostname to bind the HTTP profiling interface to")

	rootCmd.Flags().StringVarP(&appConfig.DataDirectory, "data-dir", "d", appConfig.DataDirectory, "directory to store app data (defaults to ~/.skycoin)")

	rootCmd.Flags().StringVarP(&appConfig.DaemonMode, "daemon-mode", "m", appConfig.DaemonMode, "Choices are: USB or EMULATOR")

	rootCmd.Version = fmt.Sprintf("%s (commit: %s, branch: %s)", Version, Commit, Branch)
}
