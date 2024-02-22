package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Hash []byte

type MerkleNode struct {
	Left      *MerkleNode
	Right     *MerkleNode
	HashValor Hash
}

func CalculateHash(data []byte) Hash {
	hash := sha256.Sum256(data)
	return hash[:]
}

func BuildMerkleTree(transactions [][]byte) *MerkleNode {
	var nodes []MerkleNode

	//create leaf nodes (transactions)
	for _, tx := range transactions {
		hash := CalculateHash(tx)
		nodes = append(nodes, MerkleNode{nil, nil, hash})
	}

	//build the tree by combining the hashes
	for len(nodes) > 1 {
		var level []MerkleNode

		for i := 0; i < len(nodes); i += 2 {
			if i+1 == len(nodes) { //odd case
				level = append(level, nodes[i])
			} else { //even case
				hash := CalculateHash(append(nodes[i].HashValor, nodes[i+1].HashValor...))
				level = append(level, MerkleNode{&nodes[i], &nodes[i+1], hash})
			}
		}
		nodes = level
	}

	return &nodes[0] //returns the root node
}

func main() {
	//transactions
	transactions := [][]byte{
		[]byte("Tx1"),
		[]byte("Tx2"),
		[]byte("Tx3"),
		[]byte("Tx4"),
	}

	//build merkle tree
	arbol := BuildMerkleTree(transactions)

	//print the root of the Merkle tree
	fmt.Printf("merkle tree root (Hash): %s\n", hex.EncodeToString(arbol.HashValor))
}
