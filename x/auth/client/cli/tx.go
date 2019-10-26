package cli

import (
	"github.com/spf13/cobra"

	"github.com/barkisnet/barkis/client"
	"github.com/barkisnet/barkis/codec"
	"github.com/barkisnet/barkis/x/auth/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Auth transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	txCmd.AddCommand(
		GetMultiSignCommand(cdc),
		GetSignCommand(cdc),
	)
	return txCmd
}
