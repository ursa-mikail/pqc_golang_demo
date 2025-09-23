// bist_main.go - Integration file for BIST with your existing PQC application
package main

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"pqc_bist_demo/ciphering"
	"pqc_bist_demo/hashing"
	"pqc_bist_demo/signing"
	"pqc_bist_demo/test_vectors"
	"pqc_bist_demo/util"

	"golang.org/x/crypto/chacha20"
)

func getUserHomeDir() string {
	if h, err := os.UserHomeDir(); err == nil && h != "" {
		return h
	}
	// fallback to current directory if home not available
	return "."
}

// Extended main function that includes BIST
func main() {
	fmt.Println("=== Post-Quantum Cryptography Demo with BIST ===")

	// Check if BIST mode is requested
	if len(os.Args) > 1 && os.Args[1] == "--bist" {
		runBISTMode()
		return
	}

	// Run original demo
	runOriginalDemo()

	// Ask if user wants to run BIST
	fmt.Println("\nWould you like to run the Built-In Self Test (BIST) suite?")
	fmt.Println("This will generate comprehensive test vectors and validate all algorithms.")
	fmt.Print("Run BIST? (y/N): ")

	var response string
	fmt.Scanln(&response)

	if response == "y" || response == "Y" || response == "yes" || response == "Yes" {
		runBISTMode()
	}
}

// runOriginalDemo runs your existing demo code
func runOriginalDemo() {
	// Original demo code from your main.go
	util.Kyber768_unit_demo()
	util.PQC_unit_demo()

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

		// Demonstrate KEM (Key Encapsulation) with ChaCha20
		if err := demoKEMWithChaCha20(level); err != nil {
			log.Printf("KEMWithChaCha20 demo failed: %v", err)
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

// runBISTMode runs the comprehensive BIST suite
func runBISTMode() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("STARTING POST-QUANTUM CRYPTOGRAPHY BUILT-IN SELF TEST (BIST)")
	fmt.Println(strings.Repeat("=", 80))

	startTime := time.Now()

	// Create BIST suite
	suite := test_vectors.NewBISTSuite()

	// Run comprehensive BIST
	suite.RunComprehensiveBIST()

	// Print detailed report
	suite.PrintComprehensiveReport()

	// Save test vectors and report
	if err := suite.SaveTestVectors("pqc_test_vectors.json"); err != nil {
		log.Printf("Warning: Could not save test vectors: %v", err)
	} else {
		log.Printf("Saved BIST test_vectors to: pqc_test_vectors.json")
	}

	if err := suite.SaveBISTReport("pqc_bist_report.json"); err != nil {
		log.Printf("Warning: Could not save BIST report: %v", err)
	} else {
		log.Printf("Saved BIST report to: pqc_bist_report.json")
	}

	totalTime := time.Since(startTime)
	fmt.Printf("\nTotal BIST Execution Time: %v\n", totalTime)

	// Exit with appropriate code
	exitCode := suite.GetExitCode()
	if exitCode == 0 {
		fmt.Println("✅ BIST PASSED - All exit criteria met")
		fmt.Println("System is ready for production use")
	} else {
		fmt.Println("❌ BIST FAILED - Exit criteria not met")
		fmt.Println("System requires attention before production use")
	}

	if exitCode != 0 {
		fmt.Println("❌ BIST FAILED - Exit criteria not met")
		for _, err := range suite.GetErrors() {
			fmt.Printf("  - %s\n", err)
		}
	}

	// Demonstrate test vector validation
	demonstrateTestVectorValidation(suite)

	os.Exit(exitCode)
}

// demonstrateTestVectorValidation shows how to validate individual test vectors
func demonstrateTestVectorValidation(suite *test_vectors.BISTSuite) {
	fmt.Println("\n" + strings.Repeat("-", 60))
	fmt.Println("TEST VECTOR VALIDATION DEMONSTRATION")
	fmt.Println(strings.Repeat("-", 60))

	if len(suite.TestVectors) == 0 {
		fmt.Println("No test vectors available for validation")
		return
	}

	// Validate first few test vectors as examples
	maxToValidate := min(5, len(suite.TestVectors))

	for i := 0; i < maxToValidate; i++ {
		tv := suite.TestVectors[i]
		fmt.Printf("Validating Test Vector %s (%s): ", tv.ID, tv.Algorithm)

		err := test_vectors.ValidateTestVector(tv)
		if err != nil {
			fmt.Printf("❌ FAILED - %v\n", err)
		} else {
			fmt.Printf("✅ PASSED\n")
		}
	}

	fmt.Printf("Validated %d/%d test vectors\n", maxToValidate, len(suite.TestVectors))
	fmt.Println(strings.Repeat("-", 60))
}

// Enhanced demo functions with BIST integration
func demoKEMWithBIST(level util.SecurityLevel) error {
	fmt.Printf("🔡 Enhanced KEM Demo with BIST validation:\n")

	// Generate keypair
	pubKey, privKey, err := ciphering.GenerateKeyPair(level)
	if err != nil {
		return fmt.Errorf("keypair generation failed: %w", err)
	}

	// Validate key sizes match expected values
	expectedPubSize, expectedPrivSize, expectedCtSize, expectedSsSize := ciphering.GetKeySizes(level)

	if len(pubKey) != expectedPubSize {
		return fmt.Errorf("public key size mismatch: got %d, expected %d", len(pubKey), expectedPubSize)
	}
	if len(privKey) != expectedPrivSize {
		return fmt.Errorf("private key size mismatch: got %d, expected %d", len(privKey), expectedPrivSize)
	}

	// Encapsulate
	ciphertext, sharedSecret, err := ciphering.Encapsulate(pubKey)
	if err != nil {
		return fmt.Errorf("encapsulation failed: %w", err)
	}

	// Validate ciphertext and shared secret sizes
	if len(ciphertext) != expectedCtSize {
		return fmt.Errorf("ciphertext size mismatch: got %d, expected %d", len(ciphertext), expectedCtSize)
	}
	if len(sharedSecret) != expectedSsSize {
		return fmt.Errorf("shared secret size mismatch: got %d, expected %d", len(sharedSecret), expectedSsSize)
	}

	// Decapsulate
	recoveredSecret, err := ciphering.Decapsulate(privKey, ciphertext)
	if err != nil {
		return fmt.Errorf("decapsulation failed: %w", err)
	}

	// Verify using constant-time comparison
	if subtle.ConstantTimeCompare(sharedSecret, recoveredSecret) != 1 {
		return fmt.Errorf("shared secrets do not match")
	}

	// Display results with validation status
	fmt.Printf("  Algorithm: %s ✅\n", ciphering.GetAlgorithmName(level))
	fmt.Printf("  Public Key: %d bytes ✅\n", len(pubKey))
	fmt.Printf("  Private Key: %d bytes ✅\n", len(privKey))
	fmt.Printf("  Ciphertext: %d bytes ✅\n", len(ciphertext))
	fmt.Printf("  Shared Secret: %d bytes ✅\n", len(sharedSecret))
	fmt.Printf("  BIST Validation: PASSED ✅\n")

	return nil
}

func demoSigningWithBIST(level util.SecurityLevel) error {
	fmt.Printf("✏️  Enhanced Signature Demo with BIST validation:\n")

	message := []byte("Hello, Post-Quantum World with BIST!")

	// Generate keypair
	pubKey, privKey, err := signing.GenerateKeyPair(level)
	if err != nil {
		return fmt.Errorf("keypair generation failed: %w", err)
	}

	// Validate key sizes
	expectedPubSize, expectedPrivSize, expectedSigSize := signing.GetKeySizes(level)

	if len(pubKey) != expectedPubSize {
		return fmt.Errorf("public key size mismatch: got %d, expected %d", len(pubKey), expectedPubSize)
	}
	if len(privKey) != expectedPrivSize {
		return fmt.Errorf("private key size mismatch: got %d, expected %d", len(privKey), expectedPrivSize)
	}

	// Sign message
	signature, err := signing.Sign(privKey, message)
	if err != nil {
		return fmt.Errorf("signing failed: %w", err)
	}

	// Validate signature size
	if len(signature) != expectedSigSize {
		return fmt.Errorf("signature size mismatch: got %d, expected %d", len(signature), expectedSigSize)
	}

	// Verify signature
	valid, err := signing.Verify(pubKey, message, signature)
	if err != nil {
		return fmt.Errorf("verification failed: %w", err)
	}

	if !valid {
		return fmt.Errorf("signature verification returned false")
	}

	// Test with modified message (should fail)
	modifiedMessage := make([]byte, len(message))
	copy(modifiedMessage, message)
	if len(modifiedMessage) > 0 {
		modifiedMessage[0] ^= 1
	}

	invalidValid, err := signing.Verify(pubKey, modifiedMessage, signature)
	if err != nil {
		return fmt.Errorf("modified message verification error: %w", err)
	}

	if invalidValid {
		return fmt.Errorf("modified message incorrectly verified as valid")
	}

	// Display results
	fmt.Printf("  Algorithm: %s ✅\n", signing.GetAlgorithmName(level))
	fmt.Printf("  Public Key: %d bytes ✅\n", len(pubKey))
	fmt.Printf("  Private Key: %d bytes ✅\n", len(privKey))
	fmt.Printf("  Signature: %d bytes ✅\n", len(signature))
	fmt.Printf("  Message: %q ✅\n", message)
	fmt.Printf("  Valid Signature: %v ✅\n", valid)
	fmt.Printf("  Invalid Detection: %v ✅\n", !invalidValid)
	fmt.Printf("  BIST Validation: PASSED ✅\n")

	return nil
}

// Performance monitoring functions
func measureKEMPerformance(level util.SecurityLevel, iterations int) {
	fmt.Printf("📊 KEM Performance Analysis (%s, %d iterations):\n",
		ciphering.GetAlgorithmName(level), iterations)

	var totalKeyGen, totalEncap, totalDecap time.Duration

	for i := 0; i < iterations; i++ {
		// Key generation
		start := time.Now()
		pubKey, privKey, err := ciphering.GenerateKeyPair(level)
		if err != nil {
			log.Printf("Key generation failed: %v", err)
			continue
		}
		totalKeyGen += time.Since(start)

		// Encapsulation
		start = time.Now()
		ct, ss, err := ciphering.Encapsulate(pubKey)
		if err != nil {
			log.Printf("Encapsulation failed: %v", err)
			continue
		}
		totalEncap += time.Since(start)

		// Decapsulation
		start = time.Now()
		_, err = ciphering.Decapsulate(privKey, ct)
		if err != nil {
			log.Printf("Decapsulation failed: %v", err)
			continue
		}
		totalDecap += time.Since(start)

		_ = ss // Avoid unused variable warning
	}

	avgKeyGen := totalKeyGen / time.Duration(iterations)
	avgEncap := totalEncap / time.Duration(iterations)
	avgDecap := totalDecap / time.Duration(iterations)

	fmt.Printf("  Avg Key Generation: %v\n", avgKeyGen)
	fmt.Printf("  Avg Encapsulation:  %v\n", avgEncap)
	fmt.Printf("  Avg Decapsulation:  %v\n", avgDecap)
	fmt.Printf("  Total Operations:   %v\n", avgKeyGen+avgEncap+avgDecap)
}

func measureSignaturePerformance(level util.SecurityLevel, iterations int) {
	fmt.Printf("📊 Signature Performance Analysis (%s, %d iterations):\n",
		signing.GetAlgorithmName(level), iterations)

	var totalKeyGen, totalSign, totalVerify time.Duration
	message := []byte("Performance test message for BIST")

	for i := 0; i < iterations; i++ {
		// Key generation
		start := time.Now()
		pubKey, privKey, err := signing.GenerateKeyPair(level)
		if err != nil {
			log.Printf("Key generation failed: %v", err)
			continue
		}
		totalKeyGen += time.Since(start)

		// Signing
		start = time.Now()
		signature, err := signing.Sign(privKey, message)
		if err != nil {
			log.Printf("Signing failed: %v", err)
			continue
		}
		totalSign += time.Since(start)

		// Verification
		start = time.Now()
		_, err = signing.Verify(pubKey, message, signature)
		if err != nil {
			log.Printf("Verification failed: %v", err)
			continue
		}
		totalVerify += time.Since(start)
	}

	// Averages
	avgKeyGen := totalKeyGen / time.Duration(iterations)
	avgSign := totalSign / time.Duration(iterations)
	avgVerify := totalVerify / time.Duration(iterations)

	// Print to console
	fmt.Printf("   Avg KeyGen:   %v\n", avgKeyGen)
	fmt.Printf("   Avg Signing:  %v\n", avgSign)
	fmt.Printf("   Avg Verify:   %v\n", avgVerify)

	fmt.Printf("   🔑 KeyGen   : %v (avg)\n", avgKeyGen)
	fmt.Printf("   ✍️  Sign     : %v (avg)\n", avgSign)
	fmt.Printf("   ✅ Verify   : %v (avg)\n", avgVerify)
	fmt.Printf("   ⏱️  Total    : %v (sum of averages)\n", avgKeyGen+avgSign+avgVerify)

	// Prepare JSON report
	report := map[string]interface{}{
		"algorithm":  signing.GetAlgorithmName(level),
		"iterations": iterations,
		"average": map[string]string{
			"keygen": avgKeyGen.String(),
			"sign":   avgSign.String(),
			"verify": avgVerify.String(),
		},
	}

	// Serialize to JSON
	reportBytes, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal report: %v", err)
	}

	// Define output path
	outputPath := "/Users/chanfamily/Progamming/go/trial_pqc_BIST/pqc_bist_report.json"

	// Write file
	if err := os.WriteFile(outputPath, reportBytes, 0644); err != nil {
		log.Fatalf("Failed to write report: %v", err)
	}

	log.Printf("Saved BIST report to: %s", outputPath)
	fmt.Printf("📁 Report size: %d bytes\n", len(reportBytes))

}

/*
% go mod init pqc_bist_demo
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

func demoKEMWithChaCha20(level util.SecurityLevel) error {
	fmt.Println("🔐 KEM + ChaCha20 Demo:")

	// 1. Generate keypair
	pubKey, privKey, err := ciphering.GenerateKeyPair(level)
	if err != nil {
		return fmt.Errorf("keypair generation failed: %w", err)
	}

	// 2. Encapsulate (sender)
	ciphertext, sharedSecretSender, err := ciphering.Encapsulate(pubKey)
	if err != nil {
		return fmt.Errorf("encapsulation failed: %w", err)
	}

	// 3. Decapsulate (receiver)
	sharedSecretReceiver, err := ciphering.Decapsulate(privKey, ciphertext)
	if err != nil {
		return fmt.Errorf("decapsulation failed: %w", err)
	}

	// Verify shared secrets match
	if !util.SecureCompare(sharedSecretSender, sharedSecretReceiver) {
		return fmt.Errorf("shared secrets do not match")
	}
	fmt.Printf("  ✅ Shared secret established (%d bytes)\n", len(sharedSecretSender))

	// 4. Use shared secret as ChaCha20 key
	key := sharedSecretSender // already 32 bytes (256 bits)
	nonce := make([]byte, chacha20.NonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return fmt.Errorf("nonce generation failed: %w", err)
	}

	// Encrypt a message
	plaintext := []byte("Hello PQC + ChaCha20 world!")
	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return fmt.Errorf("cipher init failed: %w", err)
	}
	ciphertextMsg := make([]byte, len(plaintext))
	cipher.XORKeyStream(ciphertextMsg, plaintext)

	// Decrypt the message
	cipherDec, _ := chacha20.NewUnauthenticatedCipher(key, nonce)
	recovered := make([]byte, len(ciphertextMsg))
	cipherDec.XORKeyStream(recovered, ciphertextMsg)

	// Show results
	fmt.Printf("  Nonce: %x\n", nonce)
	fmt.Printf("  Plaintext: %s\n", plaintext)
	fmt.Printf("  Ciphertext: %x\n", ciphertextMsg)
	fmt.Printf("  Recovered: %s\n", recovered)

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
