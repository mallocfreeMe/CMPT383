package main

import (
  "fmt"
  "blockchain"
  "encoding/hex"
)

func main() {
  b0 := blockchain.Initial(7)
  b0.Mine(1)
  fmt.Println(b0.Proof, hex.EncodeToString(b0.Hash))
  b1 := b0.Next("this is an interesting message")
  b1.Mine(1)
  fmt.Println(b1.Proof, hex.EncodeToString(b1.Hash))
  b2 := b1.Next("this is not interesting")
  b2.Mine(1)
  fmt.Println(b2.Proof, hex.EncodeToString(b2.Hash))
}
