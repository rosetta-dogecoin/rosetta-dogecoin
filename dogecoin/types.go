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

package dogecoin

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	// Blockchain is Bitcoin.
	Blockchain string = "Bitcoin"

	// MainnetNetwork is the value of the network
	// in MainnetNetworkIdentifier.
	MainnetNetwork string = "Mainnet"

	// TestnetNetwork is the value of the network
	// in TestnetNetworkIdentifier.
	TestnetNetwork string = "Testnet3"

	// Decimals is the decimals value
	// used in Currency.
	Decimals = 8

	// SatoshisInBitcoin is the number of
	// Satoshis in 1 BTC (10^8).
	SatoshisInBitcoin = 100000000

	// TransactionHashLength is the length
	// of any transaction hash in Bitcoin.
	TransactionHashLength = 64
)

var (
	// MainnetGenesisBlockIdentifier is the genesis block for mainnet.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
	}

	// MainnetParams are the params for mainnet.
	MainnetParams = &MainNetParams

	// MainnetCurrency is the *types.Currency for mainnet.
	MainnetCurrency = &types.Currency{
		Symbol:   "BTC",
		Decimals: Decimals,
	}

	// TestnetGenesisBlockIdentifier is the genesis block for testnet.
	TestnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: "000000000933ea01ad0ee984209779baaec3ced90fa3f408719526f8d77f4943",
	}

	// TestnetParams are the params for testnet.
	TestnetParams = &TestNet3Params

	// TestnetCurrency is the *types.Currency for testnet.
	TestnetCurrency = &types.Currency{
		Symbol:   "tBTC",
		Decimals: Decimals,
	}
)
