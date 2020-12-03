package blockchain

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBlockInitialAndNext(t *testing.T) {
	b := Initial(uint8(16))
	assert.Equal(t, b.PrevHash, make([]byte, 32))
	assert.Equal(t, b.Generation, uint64(0))
	assert.Equal(t, b.Difficulty, uint8(16))
	assert.Equal(t, b.Data, "")
	b1 := b.Next("Next block message")
	assert.Equal(t, b1.Generation, uint64(1))
	assert.Equal(t, b1.Difficulty, b.Difficulty)
	assert.Equal(t, b1.Data, "Next block message")
}

func TestHashString(t *testing.T) {
	b := Initial(uint8(16))
	b.SetProof(uint64(56231))
	s := b.hashString()
	assert.Equal(t, s, "0000000000000000000000000000000000000000000000000000000000000000:0:16::56231")
	b1 := b.Next("message")
	b1.SetProof(uint64(2159))
	s1 := b1.hashString()
	assert.Equal(t, s1, "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000:1:16:message:2159")
}

func TestCalcHash(t *testing.T) {
	b := Initial(uint8(16))
	b.SetProof(uint64(56231))
	s := hex.EncodeToString(b.CalcHash())
	assert.Equal(t, s, "6c71ff02a08a22309b7dbbcee45d291d4ce955caa32031c50d941e3e9dbd0000")
	b1 := b.Next("message")
	b1.SetProof(uint64(2159))
	s1 := hex.EncodeToString(b1.CalcHash())
	assert.Equal(t, s1, "9b4417b36afa6d31c728eed7abc14dd84468fdb055d8f3cbe308b0179df40000")
}

func TestValidHashes(t *testing.T) {
	b := Initial(16)
	assert.Equal(t, b.ValidHash(), false)
	b.SetProof(uint64(56231))
	assert.Equal(t, b.ValidHash(), true)
	b1 := Initial(19)
	b1.SetProof(87745)
	assert.Equal(t, b1.ValidHash(), true)
	b2 := b1.Next("hash example 1234")
	b2.SetProof(1407891)
	assert.Equal(t, b2.ValidHash(), true)
	b3 := Initial(19)
	b3.SetProof(346082)
	assert.Equal(t, b3.ValidHash(), false)
}

func TestMining(t *testing.T) {
	b0 := Initial(7)
	b0.Mine(1)
	assert.Equal(t, b0.Proof, uint64(385))
	assert.Equal(t, hex.EncodeToString(b0.Hash), "379bf2fb1a558872f09442a45e300e72f00f03f2c6f4dd29971f67ea4f3d5300")

	b1 := b0.Next("this is an interesting message")
	b1.Mine(1)
	assert.Equal(t, b1.Proof, uint64(20))
	assert.Equal(t, hex.EncodeToString(b1.Hash), "4a1c722d8021346fa2f440d7f0bbaa585e632f68fd20fed812fc944613b92500")

	b2 := b1.Next("this is not interesting")
	b2.Mine(1)
	assert.Equal(t, b2.Proof, uint64(40))
	assert.Equal(t, hex.EncodeToString(b2.Hash), "ba2f9bf0f9ec629db726f1a5fe7312eb76270459e3f5bfdc4e213df9e47cd380")
}
