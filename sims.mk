#!/usr/bin/make -f

#########################################
### Simulations

SIMAPP = github.com/FnyaMing/nainaide/app

sim-nainaide-nondeterminism:
	@echo "Running nondeterminism test..."
	@go test -mod=readonly $(SIMAPP) -run TestAppStateDeterminism -Enabled=true \
		-NumBlocks=100 -BlockSize=200 -Commit=true -v -timeout 24h

sim-nainaide-custom-genesis-fast:
	@echo "Running custom genesis simulation..."
	@echo "By default, ${HOME}/.nainaided/config/genesis.json will be used."
	@go test -mod=readonly $(SIMAPP) -run TestFullnainaideSimulation -Genesis=${HOME}/.nainaided/config/genesis.json \
		-Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

sim-nainaide-fast:
	@echo "Running quick nainaide simulation. This may take several minutes..."
	@go test -mod=readonly $(SIMAPP) -run TestFullnainaideSimulation -Enabled=true -NumBlocks=100 -BlockSize=200 -Commit=true -Seed=99 -Period=5 -v -timeout 24h

sim-nainaide-import-export: runsim
	@echo "Running nainaide import/export simulation. This may take several minutes..."
	$(GOPATH)/bin/runsim $(SIMAPP) 25 5 TestnainaideImportExport

sim-nainaide-simulation-after-import: runsim
	@echo "Running nainaide simulation-after-import. This may take several minutes..."
	$(GOPATH)/bin/runsim $(SIMAPP) 25 5 TestnainaideSimulationAfterImport

sim-nainaide-custom-genesis-multi-seed: runsim
	@echo "Running multi-seed custom genesis simulation..."
	@echo "By default, ${HOME}/.nainaided/config/genesis.json will be used."
	$(GOPATH)/bin/runsim $(SIMAPP) -g ${HOME}/.nainaided/config/genesis.json 400 5 TestFullnainaideSimulation

sim-nainaide-multi-seed: runsim
	@echo "Running multi-seed nainaide simulation. This may take awhile!"
	$(GOPATH)/bin/runsim $(SIMAPP) 400 5 TestFullnainaideSimulation

sim-benchmark-invariants:
	@echo "Running simulation invariant benchmarks..."
	@go test -mod=readonly $(SIMAPP) -benchmem -bench=BenchmarkInvariants -run=^$ \
	-Enabled=true -NumBlocks=1000 -BlockSize=200 \
	-Commit=true -Seed=57 -v -timeout 24h

SIM_NUM_BLOCKS ?= 500
SIM_BLOCK_SIZE ?= 200
SIM_COMMIT ?= true
sim-nainaide-benchmark:
	@echo "Running nainaide benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullnainaideSimulation$$  \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h

sim-nainaide-profile:
	@echo "Running nainaide benchmark for numBlocks=$(SIM_NUM_BLOCKS), blockSize=$(SIM_BLOCK_SIZE). This may take awhile!"
	@go test -mod=readonly -benchmem -run=^$$ $(SIMAPP) -bench ^BenchmarkFullnainaideSimulation$$ \
		-Enabled=true -NumBlocks=$(SIM_NUM_BLOCKS) -BlockSize=$(SIM_BLOCK_SIZE) -Commit=$(SIM_COMMIT) -timeout 24h -cpuprofile cpu.out -memprofile mem.out


.PHONY: runsim sim-nainaide-nondeterminism sim-nainaide-custom-genesis-fast sim-nainaide-fast sim-nainaide-import-export \
	sim-nainaide-simulation-after-import sim-nainaide-custom-genesis-multi-seed sim-nainaide-multi-seed \
	sim-benchmark-invariants sim-nainaide-benchmark sim-nainaide-profile
