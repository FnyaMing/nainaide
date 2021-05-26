//nolint
package app

import (
	"io"

	"github.com/tendermint/tendermint/libs/log"
	dbm "github.com/tendermint/tm-db"

	"github.com/FnyaMing/nainaide/baseapp"
	sdk "github.com/FnyaMing/nainaide/types"
	"github.com/FnyaMing/nainaide/x/staking"
)

var (
	genesisFile        string
	paramsFile         string
	exportParamsPath   string
	exportParamsHeight int
	exportStatePath    string
	exportStatsPath    string
	seed               int64
	initialBlockHeight int
	numBlocks          int
	blockSize          int
	enabled            bool
	verbose            bool
	lean               bool
	commit             bool
	period             int
	onOperation        bool // TODO Remove in favor of binary search for invariant violation
	allInvariants      bool
	genesisTime        int64
)

// DONTCOVER

// NewnainaideAppUNSAFE is used for debugging purposes only.
//
// NOTE: to not use this function with non-test code
func NewnainaideAppUNSAFE(logger log.Logger, db dbm.DB, traceStore io.Writer, loadLatest bool,
	invCheckPeriod uint, baseAppOptions ...func(*baseapp.BaseApp),
) (bapp *nainaideApp, keyMain, keyStaking *sdk.KVStoreKey, stakingKeeper staking.Keeper) {

	bapp = NewnainaideApp(logger, db, traceStore, loadLatest, invCheckPeriod, baseAppOptions...)
	return bapp, bapp.keys[baseapp.MainStoreKey], bapp.keys[staking.StoreKey], bapp.stakingKeeper
}
