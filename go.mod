module github.com/FnyaMing/nainaide

go 1.12

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20190523000118-16327141da8c
	github.com/cosmos/go-bip39 v0.0.0-20180819234021-555e2067c45d
	github.com/cosmos/ledger-cosmos-go v0.10.3
	github.com/gogo/protobuf v1.3.0
	github.com/golang/mock v1.3.1
	github.com/gorilla/mux v1.7.3
	github.com/mattn/go-isatty v0.0.8
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/otiai10/copy v1.0.1
	github.com/otiai10/curr v0.0.0-20190513014714-f5a3d24e5776 // indirect
	github.com/pelletier/go-toml v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/common v0.6.0 // indirect
	github.com/prometheus/procfs v0.0.3 // indirect
	github.com/rakyll/statik v0.1.6
	github.com/rcrowley/go-metrics v0.0.0-20190704165056-9c2d0518ed81 // indirect
	github.com/snikch/goodman v0.0.0-20171125024755-10e37e294daa
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.4.0
	github.com/stumble/gorocksdb v0.0.3 // indirect
	github.com/tendermint/btcd v0.1.1
	github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5
	github.com/tendermint/go-amino v0.15.0
	github.com/tendermint/iavl v0.12.4
	github.com/tendermint/tendermint v0.32.7
	github.com/tendermint/tm-db v0.2.0
	golang.org/x/net v0.0.0-20190827160401-ba9fcec4b297 // indirect
	golang.org/x/sys v0.0.0-20190804053845-51ab0e2deafa // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20190701230453-710ae3a149df // indirect
	gopkg.in/yaml.v2 v2.2.2
)

replace github.com/tendermint/iavl => github.com/barkisnet/iavl v0.12.4-barkis