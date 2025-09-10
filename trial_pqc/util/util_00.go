package util

import (
	"fmt"
	"log"

	"github.com/cloudflare/circl/kem/kyber/kyber768"
)

func Kyber768_unit_demo() {
	fmt.Println("=== Kyber768 Key Encapsulation Example ===")

	// Initialize Kyber768 KEM
	scheme := kyber768.Scheme()

	// Generate keypair
	publicKey, privateKey, err := scheme.GenerateKeyPair()
	if err != nil {
		log.Fatal("Failed to generate keypair:", err)
	}

	// Get byte representations using MarshalBinary
	pubKeyBytes, err := publicKey.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to marshal public key:", err)
	}

	privKeyBytes, err := privateKey.MarshalBinary()
	if err != nil {
		log.Fatal("Failed to marshal private key:", err)
	}

	fmt.Printf("Public key size: %d bytes\n", len(pubKeyBytes))
	fmt.Printf("Private key size: %d bytes\n", len(privKeyBytes))

	// Encapsulation (sender side)
	ct, ss, err := scheme.Encapsulate(publicKey)
	if err != nil {
		log.Fatal("Failed to encapsulate:", err)
	}

	fmt.Printf("Ciphertext size: %d bytes\n", len(ct))
	fmt.Printf("Shared secret size: %d bytes\n", len(ss))
	fmt.Printf("Shared secret: %x\n", ss[:16]) // Show first 16 bytes

	// Decapsulation (receiver side)
	ss2, err := scheme.Decapsulate(privateKey, ct)
	if err != nil {
		log.Fatal("Failed to decapsulate:", err)
	}

	// Verify shared secrets match
	if string(ss) == string(ss2) {
		fmt.Println("✅ Shared secrets match!")
	} else {
		fmt.Println("❌ Shared secrets don't match!")
	}
}

/*
Expected output (sizes for Kyber768):
- Public key size: 1184 bytes
- Private key size: 2400 bytes
- Ciphertext size: 1088 bytes
- Shared secret size: 32 bytes
*/
