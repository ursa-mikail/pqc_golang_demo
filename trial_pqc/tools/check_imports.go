package main

import (
	"fmt"

	"github.com/cloudflare/circl/kem/kyber/kyber1024"
	"github.com/cloudflare/circl/kem/kyber/kyber512"
	"github.com/cloudflare/circl/kem/kyber/kyber768"
	"github.com/cloudflare/circl/sign/dilithium/mode2"
	"github.com/cloudflare/circl/sign/dilithium/mode3"
	"github.com/cloudflare/circl/sign/dilithium/mode5"
	"golang.org/x/crypto/sha3"
)

func main() {
	fmt.Println("=== CIRCL Import Check ===")

	// Test Kyber imports
	fmt.Println("✅ Kyber512:", kyber512.Scheme().Name())
	fmt.Println("✅ Kyber768:", kyber768.Scheme().Name())
	fmt.Println("✅ Kyber1024:", kyber1024.Scheme().Name())

	// Test Dilithium imports - using the mode subpackages
	fmt.Println("✅ Dilithium Mode2:", mode2.Scheme().Name())
	fmt.Println("✅ Dilithium Mode3:", mode3.Scheme().Name())
	fmt.Println("✅ Dilithium Mode5:", mode5.Scheme().Name())

	// Test SHAKE functions
	shake128 := sha3.NewShake128()
	shake256 := sha3.NewShake256()

	if shake128 != nil && shake256 != nil {
		fmt.Println("✅ SHAKE128/256 available")
	}

	fmt.Println("\n=== All imports working! ===")
}

/*
go run check_imports.go
*/
