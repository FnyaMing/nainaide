package server

import (
	"encoding/json"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"errors"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	tcmd "github.com/tendermint/tendermint/cmd/tendermint/commands"
	cfg "github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	tmflags "github.com/tendermint/tendermint/libs/cli/flags"
	"github.com/tendermint/tendermint/libs/log"
	pvm "github.com/tendermint/tendermint/privval"

	"github.com/FnyaMing/nainaide/app/config"
	"github.com/FnyaMing/nainaide/client/flags"
	"github.com/FnyaMing/nainaide/codec"
	"github.com/FnyaMing/nainaide/version"
)

//___________________________________________________________________________________

// PersistentPreRunEFn returns a PersistentPreRunE function for cobra
// that initailizes the passed in context with a properly configured
// logger and config object.
func PersistentPreRunEFn(context *config.nainaideContext) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		if cmd.Name() == version.Cmd.Name() {
			return nil
		}
		err := interceptLoadConfig(context)
		if err != nil {
			return err
		}
		logger := log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		logger, err = tmflags.ParseLogLevel(context.Config.LogLevel, logger, cfg.DefaultLogLevel())
		if err != nil {
			return err
		}
		if viper.GetBool(cli.TraceFlag) {
			logger = log.NewTracingLogger(logger)
		}
		logger = logger.With("module", "main")
		context.Logger = logger
		return nil
	}
}

// If a new config is created, change some of the default tendermint settings
func interceptLoadConfig(context *config.nainaideContext) (err error) {
	tmpConf := cfg.DefaultConfig()
	err = viper.Unmarshal(tmpConf)
	if err != nil {
		panic(err)
	}
	rootDir := tmpConf.RootDir
	configFilePath := filepath.Join(rootDir, "config/config.toml")

	context.Config, err = tcmd.ParseConfig() // NOTE: ParseConfig() creates dir/files as necessary.
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		// the following parse config is needed to create directories
		context.Config, _ = tcmd.ParseConfig() // NOTE: ParseConfig() creates dir/files as necessary.
		context.Config.ProfListenAddress = "localhost:6060"
		context.Config.P2P.RecvRate = 5120000
		context.Config.P2P.SendRate = 5120000
		context.Config.TxIndex.IndexAllTags = true
		context.Config.Consensus.TimeoutCommit = 5 * time.Second
		cfg.WriteConfigFile(configFilePath, context.Config)
		// Fall through, just so that its parsed into memory.
	}

	appConfigFilePath := filepath.Join(rootDir, "config/app.toml")
	if _, err := os.Stat(appConfigFilePath); os.IsNotExist(err) {
		appConf, _ := config.ParseConfig()
		config.WriteConfigFile(appConfigFilePath, appConf)
	} else {
		err = context.ParseAppConfigInPlace()
		if err != nil {
			return err
		}
	}

	viper.SetConfigName("app")
	err = viper.MergeInConfig()

	return
}

// add server commands
func AddCommands(
	ctx *config.ServerContext, cdc *codec.Codec,
	rootCmd *cobra.Command,
	appCreator AppCreator, appExport AppExporter) {

	rootCmd.PersistentFlags().String("log_level", ctx.Config.LogLevel, "Log level")

	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}

	tendermintCmd.AddCommand(
		ShowNodeIDCmd(ctx),
		ShowValidatorCmd(ctx),
		ShowAddressCmd(ctx),
		VersionCmd(ctx),
	)

	rootCmd.AddCommand(
		StartCmd(ctx, appCreator),
		UnsafeResetAllCmd(ctx),
		flags.LineBreak,
		tendermintCmd,
		ExportCmd(ctx, cdc, appExport),
		flags.LineBreak,
		version.Cmd,
	)
}

//___________________________________________________________________________________

// InsertKeyJSON inserts a new JSON field/key with a given value to an existing
// JSON message. An error is returned if any serialization operation fails.
//
// NOTE: The ordering of the keys returned as the resulting JSON message is
// non-deterministic, so the client should not rely on key ordering.
func InsertKeyJSON(cdc *codec.Codec, baseJSON []byte, key string, value json.RawMessage) ([]byte, error) {
	var jsonMap map[string]json.RawMessage

	if err := cdc.UnmarshalJSON(baseJSON, &jsonMap); err != nil {
		return nil, err
	}

	jsonMap[key] = value
	bz, err := codec.MarshalJSONIndent(cdc, jsonMap)

	return json.RawMessage(bz), err
}

// https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
// TODO there must be a better way to get external IP
func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if skipInterface(iface) {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			ip := addrToIP(addr)
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// TrapSignal traps SIGINT and SIGTERM and terminates the server correctly.
func TrapSignal(cleanupFunc func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		if cleanupFunc != nil {
			cleanupFunc()
		}
		exitCode := 128
		switch sig {
		case syscall.SIGINT:
			exitCode += int(syscall.SIGINT)
		case syscall.SIGTERM:
			exitCode += int(syscall.SIGTERM)
		}
		os.Exit(exitCode)
	}()
}

// UpgradeOldPrivValFile converts old priv_validator.json file (prior to Tendermint 0.28)
// to the new priv_validator_key.json and priv_validator_state.json files.
func UpgradeOldPrivValFile(config *cfg.Config) {
	if _, err := os.Stat(config.OldPrivValidatorFile()); !os.IsNotExist(err) {
		if oldFilePV, err := pvm.LoadOldFilePV(config.OldPrivValidatorFile()); err == nil {
			oldFilePV.Upgrade(config.PrivValidatorKeyFile(), config.PrivValidatorStateFile())
		}
	}
}

func skipInterface(iface net.Interface) bool {
	if iface.Flags&net.FlagUp == 0 {
		return true // interface down
	}
	if iface.Flags&net.FlagLoopback != 0 {
		return true // loopback interface
	}
	return false
}

func addrToIP(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	return ip
}

// DONTCOVER
