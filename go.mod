module github.com/FnyaMing/nainaide

go 1.15

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.21.0-beta
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/ledger-cosmos-go v0.11.1
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.5.0
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-isatty v0.0.13
	github.com/otiai10/copy v1.6.0
	github.com/pelletier/go-toml v1.9.1
	github.com/pkg/errors v0.9.1
	github.com/rakyll/statik v0.1.7
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/iavl v0.0.0-00010101000000-000000000000
	github.com/tendermint/tendermint v0.34.10
	github.com/tendermint/tm-db v0.6.4
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/tendermint/iavl => github.com/barkisnet/iavl v0.12.4-barkis