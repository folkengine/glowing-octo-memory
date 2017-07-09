package main

import (
	"os"
	"fmt"
)

type Pinky struct {
	merchId	string
	pubKey 	string
	privKey	string
}

func PinkyFactory() Pinky {
	pinky := Pinky{
		merchId: os.Getenv("BRAINTREE_MERCH_ID"),
		pubKey:  os.Getenv("BRAINTREE_PUB_KEY"),
		privKey: os.Getenv("BRAINTREE_PRIV_KEY"),
	}

	if pinky.merchId == "" {
		fmt.Println("Please set BRAINTREE_MERCH_ID environment variable")
		os.Exit(1)
	}

	if pinky.pubKey == "" {
		fmt.Println("Please set BRAINTREE_PUB_KEY environment variable")
		os.Exit(1)
	}

	if pinky.pubKey == "" {
		fmt.Println("Please set BRAINTREE_PRIV_KEY environment variable")
		os.Exit(1)
	}

	return pinky
}
