package main

import (
	"fmt"
	"github.com/fykaa/go-crypto/decrypt"
	"github.com/fykaa/go-crypto/encrypt"
)

func main() {
	fmt.Println(encrypt.Nimbus("WINNER"))

	fmt.Println(decrypt.Nimbus(encrypt.Nimbus("LOOSER")))
}
