package util

import (
	"fmt"
	"log"

	"github.com/cloudflare/circl/kem/kyber/kyber768"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"
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

func PQC_unit_demo() {
	fmt.Println("=== Cloudflare CIRCL Post-Quantum Cryptography Examples ===")
	fmt.Println("Available algorithms in CIRCL v1.6.1:")
	fmt.Println("✓ ML-DSA (standardized Dilithium)")
	fmt.Println("✓ Dilithium (pre-standardization)")
	fmt.Println("✓ Kyber KEM")
	fmt.Println("✗ FALCON (not available)")
	fmt.Println("✗ SPHINCS+/SLH-DSA (not in v1.6.1, only in main branch)")
	fmt.Println("✗ HQC (not available)")
	fmt.Println()

	// Example 1: ML-DSA (standardized version of Dilithium)
	fmt.Println("=== Example 1: ML-DSA-44 (FIPS 204) ===")
	mldsaExample()
	fmt.Println()
}

func mldsaExample() {
	// Create a new ML-DSA scheme instance
	scheme := mldsa44.Scheme()

	fmt.Printf("ML-DSA-44 parameters:\n")
	fmt.Printf("  Public key size:  %d bytes\n", scheme.PublicKeySize())
	fmt.Printf("  Private key size: %d bytes\n", scheme.PrivateKeySize())
	fmt.Printf("  Signature size:   %d bytes\n", scheme.SignatureSize())

	// Generate key pair
	fmt.Println("Generating ML-DSA key pair...")
	publicKey, privateKey, err := scheme.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate ML-DSA key pair: %v", err)
	}

	// Message to sign
	message := []byte("Hello from ML-DSA (FIPS 204 standardized Dilithium)!")

	// Sign the message
	fmt.Printf("Signing message: %s\n", string(message))
	signature := scheme.Sign(privateKey, message, nil)
	fmt.Printf("Generated signature (%d bytes)\n", len(signature))

	// Verify the signature
	valid := scheme.Verify(publicKey, message, signature, nil)
	if valid {
		fmt.Println("✓ ML-DSA signature verification PASSED")
	} else {
		fmt.Println("✗ ML-DSA signature verification FAILED")
	}
}
