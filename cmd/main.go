package main

import (
	"fmt"
	"os"

	"github.com/devlongs/vanity-address-generator/generator"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run main.go <prefix>")
        os.Exit(1)
    }

    prefix := os.Args[1]
    address, privateKey := generator.GenerateVanityAddress(prefix)

    fmt.Printf("Vanity Address: %s\n", address)
    fmt.Printf("Private Key: %s\n", privateKey)
}