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

//nolint:gomnd
package dogecoin

import (
	"math/big"
	"time"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

// These variables are the chain proof-of-work limit parameters for each default
// network.
var (
	// bigOne is 1 represented as a big.Int. It is defined here to avoid
	// the overhead of creating it multiple times.
	bigOne = big.NewInt(1)

	// mainPowLimit is the highest proof of work value a Dogecoin block can
	// have for the main network. It is the value 2^236 - 1.
	mainPowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 236), bigOne)

	// testNet3PowLimit is the highest proof of work value a Dogecoin block
	// can have for the test network (version 3). It is the value 2^236 - 1.
	testNet3PowLimit = new(big.Int).Sub(new(big.Int).Lsh(bigOne, 236), bigOne)
)

// Constants used to indicate the message dogecoin network.
const (
	// MainNet represents the main dogecoin network.
	MainNet wire.BitcoinNet = 0xc0c0c0c0

	// TestNet3 represents the test network (version 3).
	TestNet3 wire.BitcoinNet = 0xdcb7c1fc
)

// MainNetParams defines the network parameters for the main Bitcoin network.
var MainNetParams = chaincfg.Params{
	Name:        "mainnet",
	Net:         wire.MainNet,
	DefaultPort: "22556",
	DNSSeeds: []chaincfg.DNSSeed{
		{Host: "seed.multidoge.org", HasFiltering: true},
		{Host: "seed2.multidoge.org", HasFiltering: false},
	},

	// Chain parameters
	GenesisBlock:             &genesisBlock,
	GenesisHash:              &genesisHash,
	PowLimit:                 mainPowLimit,
	PowLimitBits:             0x3bffffff,
	BIP0034Height:            1034383, // 251f53c3e66f122347b7667c45aaadfffb52c2a05b8c80edcce68450b639f7f6
	BIP0065Height:            3464751, // 34cd2cbba4ba366f47e5aa0db5f02c19eba2adf679ceb6653ac003bdc9a0ef1f
	BIP0066Height:            1034383, // 251f53c3e66f122347b7667c45aaadfffb52c2a05b8c80edcce68450b639f7f6
	CoinbaseMaturity:         240,
	SubsidyReductionInterval: 100000,
	TargetTimespan:           time.Minute, // 1 minute
	TargetTimePerBlock:       time.Minute, // 1 minute
	RetargetAdjustmentFactor: 4,           // 25% less, 400% more
	ReduceMinDifficulty:      false,
	MinDiffReductionTime:     0,
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{
		{Height: 0, Hash: newHashFromStr("bb0a78264637406b6360aad926284d544d7049f45189db5664f3c4d07350559e")},
		{Height: 483173, Hash: newHashFromStr("a804201ca0aceb7e937ef7a3c613a9b7589245b10cc095148c4ce4965b0b73b5")},
		{Height: 591117, Hash: newHashFromStr("5f6b93b2c28cedf32467d900369b8be6700f0649388a7dbfd3ebd4a01b1ffad8")},
		{Height: 658924, Hash: newHashFromStr("ed6c8324d9a77195ee080f225a0fca6346495e08ded99bcda47a8eea5a8a620b")},
		{Height: 703635, Hash: newHashFromStr("839fa54617adcd582d53030a37455c14a87a806f6615aa8213f13e196230ff7f")},
		{Height: 1000000, Hash: newHashFromStr("1fe4d44ea4d1edb031f52f0d7c635db8190dc871a190654c41d2450086b8ef0e")},
		{Height: 1202214, Hash: newHashFromStr("a2179767a87ee4e95944703976fee63578ec04fa3ac2fc1c9c2c83587d096977")},
		{Height: 1250000, Hash: newHashFromStr("b46affb421872ca8efa30366b09694e2f9bf077f7258213be14adb05a9f41883")},
		{Height: 1500000, Hash: newHashFromStr("0caa041b47b4d18a4f44bdc05cef1a96d5196ce7b2e32ad3e4eb9ba505144917")},
		{Height: 1750000, Hash: newHashFromStr("8042462366d854ad39b8b95ed2ca12e89a526ceee5a90042d55ebb24d5aab7e9")},
		{Height: 2000000, Hash: newHashFromStr("d6acde73e1b42fc17f29dcc76f63946d378ae1bd4eafab44d801a25be784103c")},
		{Height: 2250000, Hash: newHashFromStr("c4342ae6d9a522a02e5607411df1b00e9329563ef844a758d762d601d42c86dc")},
		{Height: 2500000, Hash: newHashFromStr("3a66ec4933fbb348c9b1889aaf2f732fe429fd9a8f74fee6895eae061ac897e2")},
		{Height: 2750000, Hash: newHashFromStr("473ea9f625d59f534ffcc9738ffc58f7b7b1e0e993078614f5484a9505885563")},
		{Height: 3062910, Hash: newHashFromStr("113c41c00934f940a41f99d18b2ad9aefd183a4b7fe80527e1e6c12779bd0246")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1900, // 95% of MinerConfirmationWindow
	MinerConfirmationWindow:       2000, //
	Deployments: [chaincfg.DefinedDeployments]chaincfg.ConsensusDeployment{
		chaincfg.DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1199145601, // January 1, 2008 UTC
			ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		chaincfg.DeploymentCSV: {
			BitNumber:  0,
			StartTime:  1462060800, // May 1, 2016 UTC
			ExpireTime: 1493596800, // May 1, 2017 UTC
		},
		chaincfg.DeploymentSegwit: {
			BitNumber:  1,
			StartTime:  1479168000, // November 15, 2016 UTC
			ExpireTime: 0,          // Disabled
		},
	},

	// Mempool parameters
	RelayNonStdTxs: false,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "doge", // Planned for 0.21

	// Address encoding magics
	PubKeyHashAddrID:        0x1e, // 30, starts with D
	ScriptHashAddrID:        0x16, // 22, starts with 9 or A
	PrivateKeyID:            0x9e, // 158, starts with 6 (uncompressed) or Q (compressed)
	WitnessPubKeyHashAddrID: 0x00, // Unimplemented
	WitnessScriptHashAddrID: 0x00, // Unimplemented

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x02, 0xfa, 0xca, 0xfd}, // starts with dgub
	HDPrivateKeyID: [4]byte{0x02, 0xfa, 0xc3, 0x98}, // starts with dgpv

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 3,
}

// TestNet3Params defines the network parameters for the test Bitcoin network
// (version 3). Not to be confused with the regression test network, this
// network is sometimes simply called "testnet".
var TestNet3Params = chaincfg.Params{
	Name:        "testnet3",
	Net:         wire.TestNet3,
	DefaultPort: "44556",
	DNSSeeds: []chaincfg.DNSSeed{
		{Host: "testseed.jrn.me.uk", HasFiltering: false},
	},

	// Chain parameters
	GenesisBlock:             &testNet3GenesisBlock,
	GenesisHash:              &testNet3GenesisHash,
	PowLimit:                 testNet3PowLimit,
	PowLimitBits:             0x3bffffff,
	BIP0034Height:            708658,  // 21b8b97dcdb94caa67c7f8f6dbf22e61e0cfe0e46e1fff3528b22864659e9b38
	BIP0065Height:            1854705, // 955bd496d23790aba1ecfacb722b089a6ae7ddabaedf7d8fb0878f48308a71f9
	BIP0066Height:            708658,  // 21b8b97dcdb94caa67c7f8f6dbf22e61e0cfe0e46e1fff3528b22864659e9b38
	CoinbaseMaturity:         240,
	SubsidyReductionInterval: 100000,
	TargetTimespan:           time.Minute, // 1 minute
	TargetTimePerBlock:       time.Minute, // 1 minute
	RetargetAdjustmentFactor: 4,           // 25% less, 400% more
	ReduceMinDifficulty:      true,
	MinDiffReductionTime:     time.Minute * 2, // TargetTimePerBlock * 2
	GenerateSupported:        false,

	// Checkpoints ordered from oldest to newest.
	Checkpoints: []chaincfg.Checkpoint{
		{Height: 0, Hash: newHashFromStr("bb0a78264637406b6360aad926284d544d7049f45189db5664f3c4d07350559e")},
		{Height: 483173, Hash: newHashFromStr("a804201ca0aceb7e937ef7a3c613a9b7589245b10cc095148c4ce4965b0b73b5")},
		{Height: 591117, Hash: newHashFromStr("5f6b93b2c28cedf32467d900369b8be6700f0649388a7dbfd3ebd4a01b1ffad8")},
		{Height: 658924, Hash: newHashFromStr("ed6c8324d9a77195ee080f225a0fca6346495e08ded99bcda47a8eea5a8a620b")},
		{Height: 703635, Hash: newHashFromStr("839fa54617adcd582d53030a37455c14a87a806f6615aa8213f13e196230ff7f")},
		{Height: 1000000, Hash: newHashFromStr("1fe4d44ea4d1edb031f52f0d7c635db8190dc871a190654c41d2450086b8ef0e")},
		{Height: 1202214, Hash: newHashFromStr("a2179767a87ee4e95944703976fee63578ec04fa3ac2fc1c9c2c83587d096977")},
		{Height: 1250000, Hash: newHashFromStr("b46affb421872ca8efa30366b09694e2f9bf077f7258213be14adb05a9f41883")},
		{Height: 1500000, Hash: newHashFromStr("0caa041b47b4d18a4f44bdc05cef1a96d5196ce7b2e32ad3e4eb9ba505144917")},
		{Height: 1750000, Hash: newHashFromStr("8042462366d854ad39b8b95ed2ca12e89a526ceee5a90042d55ebb24d5aab7e9")},
		{Height: 2000000, Hash: newHashFromStr("d6acde73e1b42fc17f29dcc76f63946d378ae1bd4eafab44d801a25be784103c")},
		{Height: 2250000, Hash: newHashFromStr("c4342ae6d9a522a02e5607411df1b00e9329563ef844a758d762d601d42c86dc")},
		{Height: 2500000, Hash: newHashFromStr("3a66ec4933fbb348c9b1889aaf2f732fe429fd9a8f74fee6895eae061ac897e2")},
		{Height: 2750000, Hash: newHashFromStr("473ea9f625d59f534ffcc9738ffc58f7b7b1e0e993078614f5484a9505885563")},
		{Height: 3062910, Hash: newHashFromStr("113c41c00934f940a41f99d18b2ad9aefd183a4b7fe80527e1e6c12779bd0246")},
	},

	// Consensus rule change deployments.
	//
	// The miner confirmation window is defined as:
	//   target proof of work timespan / target proof of work spacing
	RuleChangeActivationThreshold: 1900, // 75% of MinerConfirmationWindow
	MinerConfirmationWindow:       2000,
	Deployments: [chaincfg.DefinedDeployments]chaincfg.ConsensusDeployment{
		chaincfg.DeploymentTestDummy: {
			BitNumber:  28,
			StartTime:  1199145601, // January 1, 2008 UTC
			ExpireTime: 1230767999, // December 31, 2008 UTC
		},
		chaincfg.DeploymentCSV: {
			BitNumber:  0,
			StartTime:  1456790400, // March 1, 2016 UTC
			ExpireTime: 1493596800, // May 1, 2017 UTC
		},
		chaincfg.DeploymentSegwit: {
			BitNumber:  1,
			StartTime:  1462060800, // May 1, 2016 UTC
			ExpireTime: 0,          // Disabled
		},
	},

	// Mempool parameters
	RelayNonStdTxs: true,

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173.
	Bech32HRPSegwit: "tdge", // Planned for 0.21

	// Address encoding magics
	PubKeyHashAddrID:        0x71, // 113 starts with n
	ScriptHashAddrID:        0xc4, // 196 starts with 2
	PrivateKeyID:            0xf1, // 241 starts with 9 (uncompressed) or c (compressed)
	WitnessPubKeyHashAddrID: 0x00, // Unimplemented
	WitnessScriptHashAddrID: 0x00, // Unimplemented

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xcf}, // starts with tpub
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with tprv

	// BIP44 coin type used in the hierarchical deterministic path for
	// address generation.
	HDCoinType: 1,
}

// newHashFromStr converts the passed big-endian hex string into a
// chainhash.Hash.  It only differs from the one available in chainhash in that
// it panics on an error since it will only (and must only) be called with
// hard-coded, and therefore known good, hashes.
//
// Copied from btcd/chaincfg (it is not exported)
func newHashFromStr(hexStr string) *chainhash.Hash {
	hash, err := chainhash.NewHashFromStr(hexStr)
	if err != nil {
		// Ordinarily I don't like panics in library code since it
		// can take applications down without them having a chance to
		// recover which is extremely annoying, however an exception is
		// being made in this case because the only way this can panic
		// is if there is an error in the hard-coded hashes.  Thus it
		// will only ever potentially panic on init and therefore is
		// 100% predictable.
		panic(err)
	}
	return hash
}
