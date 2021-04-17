// Copyright 2020 Coinbase, Inc.
// Copyright 2021 Rosetta Dogecoin Developers
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
	// Blockchain is Dogecoin.
	Blockchain string = "Dogecoin"

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
	// Satoshis in 1 DOGE (10^8).
	SatoshisInBitcoin = 100000000

	// TransactionHashLength is the length
	// of any transaction hash in Dogecoin.
	TransactionHashLength = 64
)

var (
	// MainnetGenesisBlockIdentifier is the genesis block for mainnet.
	MainnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: "1a91e3dace36e2be3bf030a65679fe821aa1d6ef92e7c9902eb318182c355691",
	}

	// MainnetParams are the params for mainnet.
	MainnetParams = &MainNetParams

	// MainnetCurrency is the *types.Currency for mainnet.
	MainnetCurrency = &types.Currency{
		Symbol:   "DOGE",
		Decimals: Decimals,
	}

	// TestnetGenesisBlockIdentifier is the genesis block for testnet.
	TestnetGenesisBlockIdentifier = &types.BlockIdentifier{
		Hash: "bb0a78264637406b6360aad926284d544d7049f45189db5664f3c4d07350559e",
	}

	// TestnetParams are the params for testnet.
	TestnetParams = &TestNet3Params

	// TestnetCurrency is the *types.Currency for testnet.
	TestnetCurrency = &types.Currency{
		Symbol:   "DOGETEST",
		Decimals: Decimals,
	}
)
