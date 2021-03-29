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
	"os"
	"path"
	"testing"

	"github.com/coinbase/rosetta-bitcoin/configuration"
	"github.com/rosetta-dogecoin/rosetta-dogecoin/dogecoin"

	"github.com/coinbase/rosetta-sdk-go/storage/encoder"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/coinbase/rosetta-sdk-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfiguration(t *testing.T) {
	tests := map[string]struct {
		Mode    string
		Network string
		Port    string

		cfg *configuration.Configuration
		err error
	}{
		"no envs set": {
			err: errors.New("MODE must be populated"),
		},
		"only mode set": {
			Mode: string(configuration.Online),
			err:  errors.New("NETWORK must be populated"),
		},
		"only mode and network set": {
			Mode:    string(configuration.Online),
			Network: configuration.Mainnet,
			err:     errors.New("PORT must be populated"),
		},
		"all set (mainnet)": {
			Mode:    string(configuration.Online),
			Network: configuration.Mainnet,
			Port:    "1000",
			cfg: &configuration.Configuration{
				Mode: configuration.Online,
				Network: &types.NetworkIdentifier{
					Network:    dogecoin.MainnetNetwork,
					Blockchain: dogecoin.Blockchain,
				},
				Params:                 dogecoin.MainnetParams,
				Currency:               dogecoin.MainnetCurrency,
				GenesisBlockIdentifier: dogecoin.MainnetGenesisBlockIdentifier,
				Port:                   1000,
				RPCPort:                mainnetRPCPort,
				ConfigPath:             mainnetConfigPath,
				Pruning: &configuration.PruningConfiguration{
					Frequency: pruneFrequency,
					Depth:     pruneDepth,
					MinHeight: minPruneHeight,
				},
				Compressors: []*encoder.CompressorEntry{
					{
						Namespace:      transactionNamespace,
						DictionaryPath: mainnetTransactionDictionary,
					},
				},
			},
		},
		"all set (testnet)": {
			Mode:    string(configuration.Online),
			Network: configuration.Testnet,
			Port:    "1000",
			cfg: &configuration.Configuration{
				Mode: configuration.Online,
				Network: &types.NetworkIdentifier{
					Network:    dogecoin.TestnetNetwork,
					Blockchain: dogecoin.Blockchain,
				},
				Params:                 dogecoin.TestnetParams,
				Currency:               dogecoin.TestnetCurrency,
				GenesisBlockIdentifier: dogecoin.TestnetGenesisBlockIdentifier,
				Port:                   1000,
				RPCPort:                testnetRPCPort,
				ConfigPath:             testnetConfigPath,
				Pruning: &configuration.PruningConfiguration{
					Frequency: pruneFrequency,
					Depth:     pruneDepth,
					MinHeight: minPruneHeight,
				},
				Compressors: []*encoder.CompressorEntry{
					{
						Namespace:      transactionNamespace,
						DictionaryPath: testnetTransactionDictionary,
					},
				},
			},
		},
		"invalid mode": {
			Mode:    "bad mode",
			Network: configuration.Testnet,
			Port:    "1000",
			err:     errors.New("bad mode is not a valid mode"),
		},
		"invalid network": {
			Mode:    string(configuration.Offline),
			Network: "bad network",
			Port:    "1000",
			err:     errors.New("bad network is not a valid network"),
		},
		"invalid port": {
			Mode:    string(configuration.Offline),
			Network: configuration.Testnet,
			Port:    "bad port",
			err:     errors.New("unable to parse port bad port"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			newDir, err := utils.CreateTempDir()
			assert.NoError(t, err)
			defer utils.RemoveTempDir(newDir)

			os.Setenv(configuration.ModeEnv, test.Mode)
			os.Setenv(configuration.NetworkEnv, test.Network)
			os.Setenv(configuration.PortEnv, test.Port)

			cfg, err := LoadConfiguration(newDir)
			if test.err != nil {
				assert.Nil(t, cfg)
				assert.Contains(t, err.Error(), test.err.Error())
			} else {
				test.cfg.IndexerPath = path.Join(newDir, "indexer")
				test.cfg.BitcoindPath = path.Join(newDir, "dogecoind")
				assert.Equal(t, test.cfg, cfg)
				assert.NoError(t, err)
			}
		})
	}
}
