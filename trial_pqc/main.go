package main

import (
	"fmt"
	"log"
	"strings"

	"trial_pqc/ciphering"
	"trial_pqc/hashing"
	"trial_pqc/signing"
	"trial_pqc/util"
)

func main() {
	fmt.Println("=== Kyber768 Key Encapsulation Example ===")
	util.Kyber768_unit_demo()

	fmt.Println("=== Post-Quantum Cryptography Demo ===")

	// Demo all security levels
	securityLevels := []util.SecurityLevel{
		util.Level128, // AES-128 equivalent
		util.Level192, // AES-192 equivalent
		util.Level256, // AES-256 equivalent
	}

	for _, level := range securityLevels {
		fmt.Printf("🔒 Security Level: %s\n", level.String())
		fmt.Println(strings.Repeat("-", 50))

		// Demonstrate KEM (Key Encapsulation)
		if err := demoKEM(level); err != nil {
			log.Printf("KEM demo failed: %v", err)
		}

		// Demonstrate Digital Signatures
		if err := demoSigning(level); err != nil {
			log.Printf("Signing demo failed: %v", err)
		}

		// Demonstrate Hashing
		if err := demoHashing(level); err != nil {
			log.Printf("Hashing demo failed: %v", err)
		}

		fmt.Println()
	}
}

/*
Expected output (sizes for Kyber768):
- Public key size: 1184 bytes
- Private key size: 2400 bytes
- Ciphertext size: 1088 bytes
- Shared secret size: 32 bytes

Usage:
% go mod init trial_pqc
% go get github.com/cloudflare/circl
% go run main.go
*/

func demoKEM(level util.SecurityLevel) error {
	fmt.Println("📡 Key Encapsulation Mechanism (KEM):")

	// Generate keypair
	pubKey, privKey, err := ciphering.GenerateKeyPair(level)
	if err != nil {
		return fmt.Errorf("keypair generation failed: %w", err)
	}

	// Encapsulate
	ciphertext, sharedSecret, err := ciphering.Encapsulate(pubKey)
	if err != nil {
		return fmt.Errorf("encapsulation failed: %w", err)
	}

	// Decapsulate
	recoveredSecret, err := ciphering.Decapsulate(privKey, ciphertext)
	if err != nil {
		return fmt.Errorf("decapsulation failed: %w", err)
	}

	// Verify
	match := util.SecureCompare(sharedSecret, recoveredSecret)
	fmt.Printf("  Algorithm: %s\n", ciphering.GetAlgorithmName(level))
	fmt.Printf("  Public Key: %d bytes\n", len(pubKey))
	fmt.Printf("  Private Key: %d bytes\n", len(privKey))
	fmt.Printf("  Ciphertext: %d bytes\n", len(ciphertext))
	fmt.Printf("  Shared Secret: %d bytes\n", len(sharedSecret))
	fmt.Printf("  Secrets Match: %v ✅\n", match)

	return nil
}

func demoSigning(level util.SecurityLevel) error {
	fmt.Println("✍️  Digital Signatures:")

	message := []byte("Hello, Post-Quantum World!")

	// Generate keypair
	pubKey, privKey, err := signing.GenerateKeyPair(level)
	if err != nil {
		return fmt.Errorf("keypair generation failed: %w", err)
	}

	// Sign message
	signature, err := signing.Sign(privKey, message)
	if err != nil {
		return fmt.Errorf("signing failed: %w", err)
	}

	// Verify signature
	valid, err := signing.Verify(pubKey, message, signature)
	if err != nil {
		return fmt.Errorf("verification failed: %w", err)
	}

	fmt.Printf("  Algorithm: %s\n", signing.GetAlgorithmName(level))
	fmt.Printf("  Public Key: %d bytes\n", len(pubKey))
	fmt.Printf("  Private Key: %d bytes\n", len(privKey))
	fmt.Printf("  Signature: %d bytes\n", len(signature))
	fmt.Printf("  Message: %q\n", message)
	fmt.Printf("  Valid: %v ✅\n", valid)

	return nil
}

func demoHashing(level util.SecurityLevel) error {
	fmt.Println("🏷️  Post-Quantum Hashing:")

	data := []byte("Data to be hashed with post-quantum security")

	// Hash the data
	hash, err := hashing.Hash(data, level)
	if err != nil {
		return fmt.Errorf("hashing failed: %w", err)
	}

	// Get algorithm info
	alg := hashing.GetAlgorithm(level)

	fmt.Printf("  Algorithm: %s\n", alg)
	fmt.Printf("  Input: %d bytes\n", len(data))
	fmt.Printf("  Hash: %d bytes\n", len(hash))
	fmt.Printf("  Hash (hex): %x\n", hash[:min(len(hash), 16)]) // Show first 16 bytes

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
