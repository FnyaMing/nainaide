package main

import (
	"encoding/json"
	"io"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/FnyaMing/nainaide/app"
	"github.com/FnyaMing/nainaide/baseapp"
	"github.com/FnyaMing/nainaide/client"
	"github.com/FnyaMing/nainaide/server"
	"github.com/FnyaMing/nainaide/store"
	sdk "github.com/FnyaMing/nainaide/types"
	"github.com/FnyaMing/nainaide/x/genaccounts"
	genaccscli "github.com/FnyaMing/nainaide/x/genaccounts/client/cli"
	genutilcli "github.com/FnyaMing/nainaide/x/genutil/client/cli"
	"github.com/FnyaMing/nainaide/x/staking"
)

// nainaided custom flags
const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(sdk.Bech32PrefixAccAddr, sdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(sdk.Bech32PrefixValAddr, sdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(sdk.Bech32PrefixConsAddr, sdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := app.nainaideContext
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "nainaided",
		Short:             "nainaide Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(genutilcli.InitCmd(ctx.ServerContext, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx.ServerContext, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx.ServerContext, cdc))
	rootCmd.AddCommand(genutilcli.GenTxCmd(ctx.ServerContext, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx.ServerContext, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccscli.AddGenesisAccountCmd(ctx.ServerContext, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.AddCommand(testnetCmd(ctx.ServerContext, cdc, app.ModuleBasics, genaccounts.AppModuleBasic{}))
	rootCmd.AddCommand(replayCmd())

	server.AddCommands(ctx.ServerContext, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "GA", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewnainaideApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool, jailWhiteList []string,
) (json.RawMessage, []tmtypes.GenesisValidator, error) {

	if height != -1 {
		gApp := app.NewnainaideApp(logger, db, traceStore, false, uint(1))
		err := gApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	gApp := app.NewnainaideApp(logger, db, traceStore, true, uint(1))
	return gApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
