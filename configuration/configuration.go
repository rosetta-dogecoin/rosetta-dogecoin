// Copyright 2020 Coinbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package configuration

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/coinbase/rosetta-bitcoin/configuration"
	"github.com/rosetta-dogecoin/rosetta-dogecoin/dogecoin"

	"github.com/coinbase/rosetta-sdk-go/storage/encoder"
	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	// mainnetConfigPath is the path of the Bitcoin
	// configuration file for mainnet.
	mainnetConfigPath = "/app/bitcoin-mainnet.conf"

	// testnetConfigPath is the path of the Bitcoin
	// configuration file for testnet.
	testnetConfigPath = "/app/bitcoin-testnet.conf"

	// Zstandard compression dictionaries
	transactionNamespace         = "transaction"
	testnetTransactionDictionary = "/app/testnet-transaction.zstd"
	mainnetTransactionDictionary = "/app/mainnet-transaction.zstd"

	mainnetRPCPort = 22555
	testnetRPCPort = 44555

	// min prune depth is 288:
	// https://github.com/bitcoin/bitcoin/blob/ad2952d17a2af419a04256b10b53c7377f826a27/src/validation.h#L84
	pruneDepth = int64(10000) //nolint

	// min prune height (on mainnet):
	// https://github.com/bitcoin/bitcoin/blob/62d137ac3b701aae36c1aa3aa93a83fd6357fde6/src/chainparams.cpp#L102
	minPruneHeight = int64(100000) //nolint

	// attempt to prune once an hour
	pruneFrequency = 60 * time.Minute

	// DataDirectory is the default location for all
	// persistent data.
	DataDirectory = "/data"

	bitcoindPath = "dogecoind"
	indexerPath  = "indexer"

	// allFilePermissions specifies anyone can do anything
	// to the file.
	allFilePermissions = 0777
)

// LoadConfiguration attempts to create a new Configuration
// using the ENVs in the environment.
func LoadConfiguration(baseDirectory string) (*configuration.Configuration, error) {
	config := &configuration.Configuration{}
	config.Pruning = &configuration.PruningConfiguration{
		Frequency: pruneFrequency,
		Depth:     pruneDepth,
		MinHeight: minPruneHeight,
	}

	modeValue := configuration.Mode(os.Getenv(configuration.ModeEnv))
	switch modeValue {
	case configuration.Online:
		config.Mode = configuration.Online
		config.IndexerPath = path.Join(baseDirectory, indexerPath)
		if err := ensurePathExists(config.IndexerPath); err != nil {
			return nil, fmt.Errorf("%w: unable to create indexer path", err)
		}

		config.BitcoindPath = path.Join(baseDirectory, bitcoindPath)
		if err := ensurePathExists(config.BitcoindPath); err != nil {
			return nil, fmt.Errorf("%w: unable to create bitcoind path", err)
		}
	case configuration.Offline:
		config.Mode = configuration.Offline
	case "":
		return nil, errors.New("MODE must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid mode", modeValue)
	}

	networkValue := os.Getenv(configuration.NetworkEnv)
	switch networkValue {
	case configuration.Mainnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: dogecoin.Blockchain,
			Network:    dogecoin.MainnetNetwork,
		}
		config.GenesisBlockIdentifier = dogecoin.MainnetGenesisBlockIdentifier
		config.Params = dogecoin.MainnetParams
		config.Currency = dogecoin.MainnetCurrency
		config.ConfigPath = mainnetConfigPath
		config.RPCPort = mainnetRPCPort
		config.Compressors = []*encoder.CompressorEntry{
			{
				Namespace:      transactionNamespace,
				DictionaryPath: mainnetTransactionDictionary,
			},
		}
	case configuration.Testnet:
		config.Network = &types.NetworkIdentifier{
			Blockchain: dogecoin.Blockchain,
			Network:    dogecoin.TestnetNetwork,
		}
		config.GenesisBlockIdentifier = dogecoin.TestnetGenesisBlockIdentifier
		config.Params = dogecoin.TestnetParams
		config.Currency = dogecoin.TestnetCurrency
		config.ConfigPath = testnetConfigPath
		config.RPCPort = testnetRPCPort
		config.Compressors = []*encoder.CompressorEntry{
			{
				Namespace:      transactionNamespace,
				DictionaryPath: testnetTransactionDictionary,
			},
		}
	case "":
		return nil, errors.New("NETWORK must be populated")
	default:
		return nil, fmt.Errorf("%s is not a valid network", networkValue)
	}

	portValue := os.Getenv(configuration.PortEnv)
	if len(portValue) == 0 {
		return nil, errors.New("PORT must be populated")
	}

	port, err := strconv.Atoi(portValue)
	if err != nil || len(portValue) == 0 || port <= 0 {
		return nil, fmt.Errorf("%w: unable to parse port %s", err, portValue)
	}
	config.Port = port

	return config, nil
}

// ensurePathsExist directories along
// a path if they do not exist.
func ensurePathExists(path string) error {
	if err := os.MkdirAll(path, os.FileMode(allFilePermissions)); err != nil {
		return fmt.Errorf("%w: unable to create %s directory", err, path)
	}

	return nil
}
