package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

type Block struct {
	PrevHash   []byte
	Generation uint64
	Difficulty uint8
	Data       string
	Proof      uint64
	Hash       []byte
}

// Create new initial (generation 0) block, not setting the .Hash value.
func Initial(difficulty uint8) Block {
	// TODO
	b := new(Block)
	b.PrevHash = make([]byte, 32)
	b.Generation = uint64(0)
	b.Difficulty = difficulty
	b.Data = ""
	return *b
}

// Create new block to follow this block, with provided data, not setting the .Hash value.
func (prev_block Block) Next(data string) Block {
	// TODO
	b := new(Block)
	b.PrevHash = prev_block.Hash
	b.Generation = prev_block.Generation + uint64(1)
	b.Difficulty = prev_block.Difficulty
	b.Data = data
	return *b
}

// String that we hash for this block.
func (blk Block) hashString() string {
	return blk.hashStringProof(blk.Proof)
}

// String that we hash for this block, if we had blk.Proof == proof.
func (blk Block) hashStringProof(proof uint64) string {
	// TODO
	var result string

	if blk.Proof == proof {
		pHash := hex.EncodeToString(blk.PrevHash)
		generation := strconv.Itoa(int(blk.Generation))
		difficulty := strconv.Itoa(int(blk.Difficulty))
		data := blk.Data
		p := strconv.Itoa(int(blk.Proof))

		result = pHash + ":" + generation + ":" + difficulty + ":" + data + ":" + p
	}

	return result
}

// Calculate hash as if blk.Proof == proof.
// Separated from .CalcHash so we can test many proof values without
// modifying the Block.
func (blk Block) calcHashProof(proof uint64) []byte {
	// TODO
	var result []byte

	if blk.Proof == proof {
		h := sha256.New()
		h.Write([]byte(blk.hashString()))
		result = h.Sum(nil)
	}

	return result
}

// Calculate the block's hash.
func (blk Block) CalcHash() []byte {
	return blk.calcHashProof(blk.Proof)
}

// Would this hash end in enough null bits, if blk.Proof == proof?
func (blk Block) validHashProof(proof uint64) bool {
	// TODO
	result := false

	if blk.Hash != nil {
		if blk.Proof == proof {
			nBytes := blk.Difficulty / 8
			nBits := blk.Difficulty % 8
			for i := len(blk.Hash) - int(nBytes); i < len(blk.Hash); i++ {
				if blk.Hash[i] != '\x00' {
					return result
				}
			}

			if blk.Hash[len(blk.Hash)-int(nBytes)-1]%(1<<nBits) != 0 {
				return result
			}
			result = true
		}
	}

	return result
}

// Is this block's hash valid?
func (blk Block) ValidHash() bool {
	return blk.validHashProof(blk.Proof)
}

// Set the proof-of-work and calculate the block's "true" hash.
func (blk *Block) SetProof(proof uint64) {
	blk.Proof = proof
	blk.Hash = blk.CalcHash()
}
