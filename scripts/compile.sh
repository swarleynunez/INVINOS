#!/usr/bin/env bash

# Compile Solidity contracts (solc 0.8.21)
./bin/macos/solc --evm-version paris --optimize --optimize-runs=1000000 --abi --bin -o bin/contracts --overwrite contracts/Traceability.sol

# Generate Go bindings
./bin/macos/abigen --abi bin/contracts/Traceability.abi --bin bin/contracts/Traceability.bin --type Traceability --pkg bindings --out core/bindings/traceability.go
./bin/macos/abigen --abi bin/contracts/Auth.abi --bin bin/contracts/Auth.bin --type Auth --pkg bindings --out core/bindings/auth.go
./bin/macos/abigen --abi bin/contracts/EntityInfo.abi --bin bin/contracts/EntityInfo.bin --type EntityInfo --pkg bindings --out core/bindings/entity_info.go
