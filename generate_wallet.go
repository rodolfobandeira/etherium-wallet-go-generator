package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {
	// Generaing the private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println("Private key: ", hexutil.Encode(privateKeyBytes)[2:])

	// Generating the public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// Here is the raw public key. We won't share this
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("Public key: ", hexutil.Encode(publicKeyBytes)[4:])

	// Here is the address we'll use to receive etherium and the address to share with the world
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println("Address to receive etherium: ", address)

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("Address Hex: ", hexutil.Encode(hash.Sum(nil)[12:]))
}
