// Copyright 2021 Dogecoin Rosetta Developers
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

package bitcoin

import (
	"encoding/binary"
	"io"

	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/wire"
)

const vAuxPoW = 0x0100

// MerkleBranch defines a merkel branch
// https://en.bitcoin.it/wiki/Merged_mining_specification#Merkle_Branch
type MerkleBranch struct {
	BranchHashes   []*chainhash.Hash
	BranchSideMask int32
}

// AuxHeader defines an AuxPoW header
type AuxHeader struct {
	CoinbaseTx       wire.MsgTx
	BlockHash        chainhash.Hash
	CoinbaseBranch   MerkleBranch
	BlockchainBranch MerkleBranch
	ParentHeader     wire.BlockHeader
}

// AuxBlock defines a AuxPoW block message.
// It is used to deliver block and transaction information.
// https://en.bitcoin.it/wiki/Merged_mining_specification#Aux_proof-of-work_block
type AuxBlock struct {
	Header       wire.BlockHeader
	AuxPoW       AuxHeader
	Transactions []*wire.MsgTx
}

// Deserialize decodes a merkel branch from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (mb *MerkleBranch) Deserialize(r io.Reader) error {
	branchCount, err := wire.ReadVarInt(r, 0)
	if err != nil {
		return err
	}

	mb.BranchHashes = make([]*chainhash.Hash, 0, branchCount)
	for i := uint64(0); i < branchCount; i++ {
		hash := chainhash.Hash{}
		if _, err := io.ReadFull(r, hash[:]); err != nil {
			return err
		}
		mb.BranchHashes = append(mb.BranchHashes, &hash)
	}

	if err := binary.Read(r, binary.LittleEndian, &mb.BranchSideMask); err != nil {
		return err
	}

	return nil
}

// Deserialize decodes a AuxPow header from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (aux *AuxHeader) Deserialize(r io.Reader) error {
	if err := aux.CoinbaseTx.Deserialize(r); err != nil {
		return err
	}
	if _, err := io.ReadFull(r, aux.BlockHash[:]); err != nil {
		return err
	}
	if err := aux.CoinbaseBranch.Deserialize(r); err != nil {
		return err
	}
	if err := aux.BlockchainBranch.Deserialize(r); err != nil {
		return err
	}
	if err := aux.ParentHeader.Deserialize(r); err != nil {
		return err
	}
	return nil
}

// Deserialize decodes a AuxPoW block from r into the receiver using a format that is
// suitable for long-term storage such as a database while respecting the
// Version field in the block.
func (b *AuxBlock) Deserialize(r io.Reader) error {
	if err := b.Header.Deserialize(r); err != nil {
		return err
	}

	if (b.Header.Version & vAuxPoW) != 0 {
		if err := b.AuxPoW.Deserialize(r); err != nil {
			return err
		}
	}

	txCount, err := wire.ReadVarInt(r, 0)
	if err != nil {
		return err
	}

	b.Transactions = make([]*wire.MsgTx, 0, txCount)
	for i := uint64(0); i < txCount; i++ {
		tx := wire.MsgTx{}
		if err := tx.Deserialize(r); err != nil {
			return err
		}
		b.Transactions = append(b.Transactions, &tx)
	}

	return nil
}
