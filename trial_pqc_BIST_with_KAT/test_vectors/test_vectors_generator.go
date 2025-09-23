package test_vectors

import (
	"crypto/subtle"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"pqc_bist_demo/ciphering"
	"pqc_bist_demo/signing"
	"pqc_bist_demo/util"
)

// TestVector represents a single test vector with expected values
type TestVector struct {
	ID             string `json:"id"`
	Algorithm      string `json:"algorithm"`
	SecurityLevel  string `json:"security_level"`
	PublicKey      string `json:"public_key,omitempty"`
	PrivateKey     string `json:"private_key,omitempty"`
	Message        string `json:"message,omitempty"`
	Signature      string `json:"signature,omitempty"`
	Ciphertext     string `json:"ciphertext,omitempty"`
	SharedSecret   string `json:"shared_secret,omitempty"`
	ExpectedResult bool   `json:"expected_result"`
	Description    string `json:"description"`
}

// BISTResult represents the result of a Built-In Self Test
type BISTResult struct {
	TestID        string        `json:"test_id"`
	Algorithm     string        `json:"algorithm"`
	TestName      string        `json:"test_name"`
	Passed        bool          `json:"passed"`
	ExecutionTime time.Duration `json:"execution_time"`
	ErrorMessage  string        `json:"error_message,omitempty"`
	Iterations    int           `json:"iterations,omitempty"`
	TestVectors   int           `json:"test_vectors,omitempty"`
}

// BISTSuite contains all BIST tests and results
type BISTSuite struct {
	Results      []BISTResult `json:"results"`
	TestVectors  []TestVector `json:"test_vectors"`
	StartTime    time.Time    `json:"start_time"`
	EndTime      time.Time    `json:"end_time"`
	TotalTests   int          `json:"total_tests"`
	PassedTests  int          `json:"passed_tests"`
	FailedTests  int          `json:"failed_tests"`
	ExitCriteria bool         `json:"exit_criteria_met"`

	Errors []string
}

// NewBISTSuite creates a new BIST suite
func NewBISTSuite() *BISTSuite {
	return &BISTSuite{
		Results:     make([]BISTResult, 0),
		TestVectors: make([]TestVector, 0),
		StartTime:   time.Now(),
	}
}

// AddResult adds a test result to the suite
func (bs *BISTSuite) AddResult(result BISTResult) {
	bs.Results = append(bs.Results, result)
	bs.TotalTests++
	if result.Passed {
		bs.PassedTests++
	} else {
		bs.FailedTests++
	}
}

// AddTestVector adds a test vector to the suite
func (bs *BISTSuite) AddTestVector(vector TestVector) {
	bs.TestVectors = append(bs.TestVectors, vector)
}

// GenerateKEMTestVectors generates comprehensive test vectors for KEM algorithms
func (bs *BISTSuite) GenerateKEMTestVectors() {
	securityLevels := []util.SecurityLevel{
		util.Level128, util.Level192, util.Level256,
	}

	vectorID := 1
	for _, level := range securityLevels {
		algName := ciphering.GetAlgorithmName(level)

		// Generate multiple test vectors per algorithm
		for i := 0; i < 5; i++ {
			// Generate valid keypair
			pubKey, privKey, err := ciphering.GenerateKeyPair(level)
			if err != nil {
				log.Printf("Failed to generate keypair for %s: %v", algName, err)
				continue
			}

			// Generate valid encapsulation
			ciphertext, sharedSecret, err := ciphering.Encapsulate(pubKey)
			if err != nil {
				log.Printf("Failed to encapsulate for %s: %v", algName, err)
				continue
			}

			// Valid test vector
			bs.AddTestVector(TestVector{
				ID:             fmt.Sprintf("KEM-%03d", vectorID),
				Algorithm:      algName,
				SecurityLevel:  level.String(),
				PublicKey:      hex.EncodeToString(pubKey),
				PrivateKey:     hex.EncodeToString(privKey),
				Ciphertext:     hex.EncodeToString(ciphertext),
				SharedSecret:   hex.EncodeToString(sharedSecret),
				ExpectedResult: true,
				Description:    fmt.Sprintf("Valid %s KEM operation", algName),
			})
			vectorID++

			// Invalid test vector (corrupted ciphertext)
			if len(ciphertext) > 0 {
				corruptedCiphertext := make([]byte, len(ciphertext))
				copy(corruptedCiphertext, ciphertext)
				corruptedCiphertext[0] ^= 1 // Flip one bit

				bs.AddTestVector(TestVector{
					ID:             fmt.Sprintf("KEM-%03d", vectorID),
					Algorithm:      algName,
					SecurityLevel:  level.String(),
					PublicKey:      hex.EncodeToString(pubKey),
					PrivateKey:     hex.EncodeToString(privKey),
					Ciphertext:     hex.EncodeToString(corruptedCiphertext),
					SharedSecret:   "", // No expected shared secret for invalid input
					ExpectedResult: false,
					Description:    fmt.Sprintf("Invalid %s KEM operation - corrupted ciphertext", algName),
				})
				vectorID++
			}
		}
	}
}

// GenerateSignatureTestVectors generates comprehensive test vectors for signature algorithms
func (bs *BISTSuite) GenerateSignatureTestVectors() {
	securityLevels := []util.SecurityLevel{
		util.Level128, util.Level192, util.Level256,
	}

	testMessages := [][]byte{
		[]byte(""),
		[]byte("a"),
		[]byte("Hello, Post-Quantum World!"),
		[]byte("The quick brown fox jumps over the lazy dog"),
		make([]byte, 1000), // Large message
	}

	// Fill large message with pattern
	for i := range testMessages[4] {
		testMessages[4][i] = byte(i % 256)
	}

	vectorID := 1
	for _, level := range securityLevels {
		algName := signing.GetAlgorithmName(level)

		// Generate keypair for this security level
		pubKey, privKey, err := signing.GenerateKeyPair(level)
		if err != nil {
			log.Printf("Failed to generate keypair for %s: %v", algName, err)
			continue
		}

		for msgIdx, message := range testMessages {
			// Generate valid signature
			signature, err := signing.Sign(privKey, message)
			if err != nil {
				log.Printf("Failed to sign for %s: %v", algName, err)
				continue
			}

			// Valid test vector
			bs.AddTestVector(TestVector{
				ID:             fmt.Sprintf("SIG-%03d", vectorID),
				Algorithm:      algName,
				SecurityLevel:  level.String(),
				PublicKey:      hex.EncodeToString(pubKey),
				PrivateKey:     hex.EncodeToString(privKey),
				Message:        hex.EncodeToString(message),
				Signature:      hex.EncodeToString(signature),
				ExpectedResult: true,
				Description:    fmt.Sprintf("Valid %s signature for message type %d", algName, msgIdx),
			})
			vectorID++

			// Invalid test vector (corrupted signature)
			if len(signature) > 10 {
				corruptedSignature := make([]byte, len(signature))
				copy(corruptedSignature, signature)
				corruptedSignature[len(signature)/2] ^= 1 // Flip bit in middle

				bs.AddTestVector(TestVector{
					ID:             fmt.Sprintf("SIG-%03d", vectorID),
					Algorithm:      algName,
					SecurityLevel:  level.String(),
					PublicKey:      hex.EncodeToString(pubKey),
					PrivateKey:     hex.EncodeToString(privKey),
					Message:        hex.EncodeToString(message),
					Signature:      hex.EncodeToString(corruptedSignature),
					ExpectedResult: false,
					Description:    fmt.Sprintf("Invalid %s signature - corrupted signature", algName),
				})
				vectorID++
			}

			// Invalid test vector (wrong message)
			if len(message) > 0 {
				wrongMessage := make([]byte, len(message))
				copy(wrongMessage, message)
				wrongMessage[0] ^= 1

				bs.AddTestVector(TestVector{
					ID:             fmt.Sprintf("SIG-%03d", vectorID),
					Algorithm:      algName,
					SecurityLevel:  level.String(),
					PublicKey:      hex.EncodeToString(pubKey),
					PrivateKey:     hex.EncodeToString(privKey),
					Message:        hex.EncodeToString(wrongMessage),
					Signature:      hex.EncodeToString(signature),
					ExpectedResult: false,
					Description:    fmt.Sprintf("Invalid %s signature - wrong message", algName),
				})
				vectorID++
			}
		}
	}
}

// RunKEMBIST executes BIST tests for KEM algorithms using test vectors
func (bs *BISTSuite) RunKEMBIST() {
	kemVectors := make([]TestVector, 0)
	for _, tv := range bs.TestVectors {
		if strings.HasPrefix(tv.ID, "KEM-") {
			kemVectors = append(kemVectors, tv)
		}
	}

	if len(kemVectors) == 0 {
		bs.AddResult(BISTResult{
			TestID:       "KEM-BIST-001",
			Algorithm:    "KEM",
			TestName:     "Test Vector Generation",
			Passed:       false,
			ErrorMessage: "No KEM test vectors generated",
		})
		return
	}

	start := time.Now()
	passedVectors := 0
	failedVectors := 0

	for idx, tv := range kemVectors {
		// Decode test vector data
		_, err := hex.DecodeString(tv.PublicKey) // _:= pubKey
		if err != nil {
			failedVectors++
			bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Failed to decode PublicKey: %v", idx, err))

			continue
		}

		privKey, err := hex.DecodeString(tv.PrivateKey)
		if err != nil {
			failedVectors++
			bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Failed to decode PrivateKey: %v", idx, err))

			continue
		}

		ciphertext, err := hex.DecodeString(tv.Ciphertext)
		if err != nil {
			failedVectors++
			bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Failed to decode Ciphertext: %v", idx, err))

			continue
		}

		// Attempt decapsulation
		recoveredSecret, err := ciphering.Decapsulate(privKey, ciphertext)

		if tv.ExpectedResult {
			// Should succeed
			if err != nil {
				failedVectors++
				bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Decapsulation failed unexpectedly: %v", idx, err))

				continue
			}

			expectedSecret, err := hex.DecodeString(tv.SharedSecret)
			if err != nil {
				failedVectors++
				bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Shared secret mismatch", idx))

				continue
			}

			// Compare shared secrets
			if subtle.ConstantTimeCompare(recoveredSecret, expectedSecret) == 1 {
				passedVectors++
			} else {
				failedVectors++
			}
		} else {
			// Should fail or produce different result
			if err != nil {
				passedVectors++ // Expected to fail
			} else {
				// Check if result is different from expected
				if tv.SharedSecret != "" {
					expectedSecret, _ := hex.DecodeString(tv.SharedSecret)
					if subtle.ConstantTimeCompare(recoveredSecret, expectedSecret) != 1 {
						passedVectors++ // Expected to be different
					} else {
						failedVectors++ // Unexpected match
						bs.Errors = append(bs.Errors, fmt.Sprintf("Test vector %d: Unexpected match for invalid input", idx))

					}
				} else {
					passedVectors++ // Any result is acceptable for invalid input
				}
			}
		}
	}

	if len(bs.Errors) > 0 {
		fmt.Println("🔴 Exit Criteria Errors:")
		for _, e := range bs.Errors {
			fmt.Printf("  - %s\n", e)
		}
	}

	// Record results
	bs.AddResult(BISTResult{
		TestID:        "KEM-BIST-001",
		Algorithm:     "KEM",
		TestName:      "Test Vector Validation",
		Passed:        failedVectors == 0,
		ExecutionTime: time.Since(start),
		TestVectors:   len(kemVectors),
		ErrorMessage: func() string {
			if failedVectors > 0 {
				return fmt.Sprintf("%d/%d test vectors failed", failedVectors, len(kemVectors))
			}
			return ""
		}(),
	})

	// Individual algorithm tests
	algorithms := map[string]util.SecurityLevel{
		"Kyber512":  util.Level128,
		"Kyber768":  util.Level192,
		"Kyber1024": util.Level256,
	}

	for algName, level := range algorithms {
		bs.runKEMAlgorithmBIST(algName, level)
	}
}

// runKEMAlgorithmBIST runs BIST for a specific KEM algorithm
func (bs *BISTSuite) runKEMAlgorithmBIST(algorithm string, level util.SecurityLevel) {
	start := time.Now()
	iterations := 50
	allPassed := true
	errorMsg := ""

	for i := 0; i < iterations && allPassed; i++ {
		// Generate keypair
		pubKey, privKey, err := ciphering.GenerateKeyPair(level)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Key generation failed: %v", err)
			break
		}

		// Encapsulate
		ciphertext, sharedSecret1, err := ciphering.Encapsulate(pubKey)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Encapsulation failed: %v", err)
			break
		}

		// Decapsulate
		sharedSecret2, err := ciphering.Decapsulate(privKey, ciphertext)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Decapsulation failed: %v", err)
			break
		}

		// Verify shared secrets match
		if subtle.ConstantTimeCompare(sharedSecret1, sharedSecret2) != 1 {
			allPassed = false
			errorMsg = "Shared secrets do not match"
			break
		}

		// Verify key and ciphertext sizes
		expectedPubSize, expectedPrivSize, expectedCtSize, expectedSsSize := ciphering.GetKeySizes(level)
		if len(pubKey) != expectedPubSize || len(privKey) != expectedPrivSize ||
			len(ciphertext) != expectedCtSize || len(sharedSecret1) != expectedSsSize {
			allPassed = false
			errorMsg = "Unexpected key or ciphertext sizes"
			break
		}
	}

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("KEM-%s-BIST", algorithm),
		Algorithm:     algorithm,
		TestName:      "Algorithm Stress Test",
		Passed:        allPassed,
		ExecutionTime: time.Since(start),
		Iterations:    iterations,
		ErrorMessage:  errorMsg,
	})
}

// RunSignatureBIST executes BIST tests for signature algorithms using test vectors
func (bs *BISTSuite) RunSignatureBIST() {
	sigVectors := make([]TestVector, 0)
	for _, tv := range bs.TestVectors {
		if strings.HasPrefix(tv.ID, "SIG-") {
			sigVectors = append(sigVectors, tv)
		}
	}

	if len(sigVectors) == 0 {
		bs.AddResult(BISTResult{
			TestID:       "SIG-BIST-001",
			Algorithm:    "SIGNATURE",
			TestName:     "Test Vector Generation",
			Passed:       false,
			ErrorMessage: "No signature test vectors generated",
		})
		return
	}

	start := time.Now()
	passedVectors := 0
	failedVectors := 0

	for _, tv := range sigVectors {
		// Decode test vector data
		pubKey, err := hex.DecodeString(tv.PublicKey)
		if err != nil {
			failedVectors++
			continue
		}

		message, err := hex.DecodeString(tv.Message)
		if err != nil {
			failedVectors++
			continue
		}

		signature, err := hex.DecodeString(tv.Signature)
		if err != nil {
			failedVectors++
			continue
		}

		// Attempt verification
		valid, err := signing.Verify(pubKey, message, signature)
		if err != nil {
			failedVectors++
			continue
		}

		if valid == tv.ExpectedResult {
			passedVectors++
		} else {
			failedVectors++
		}
	}

	// Record results
	bs.AddResult(BISTResult{
		TestID:        "SIG-BIST-001",
		Algorithm:     "SIGNATURE",
		TestName:      "Test Vector Validation",
		Passed:        failedVectors == 0,
		ExecutionTime: time.Since(start),
		TestVectors:   len(sigVectors),
		ErrorMessage: func() string {
			if failedVectors > 0 {
				return fmt.Sprintf("%d/%d test vectors failed", failedVectors, len(sigVectors))
			}
			return ""
		}(),
	})

	// Individual algorithm tests
	algorithms := map[string]util.SecurityLevel{
		"Dilithium2": util.Level128,
		"Dilithium3": util.Level192,
		"Dilithium5": util.Level256,
	}

	for algName, level := range algorithms {
		bs.runSignatureAlgorithmBIST(algName, level)
	}
}

// runSignatureAlgorithmBIST runs BIST for a specific signature algorithm
func (bs *BISTSuite) runSignatureAlgorithmBIST(algorithm string, level util.SecurityLevel) {
	start := time.Now()
	iterations := 25
	allPassed := true
	errorMsg := ""

	testMessages := [][]byte{
		[]byte(""),
		[]byte("test message"),
		[]byte("Hello, Post-Quantum Cryptography!"),
		make([]byte, 1024),
	}

	// Fill large message
	for i := range testMessages[3] {
		testMessages[3][i] = byte(i % 256)
	}

	for i := 0; i < iterations && allPassed; i++ {
		// Generate keypair
		pubKey, privKey, err := signing.GenerateKeyPair(level)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Key generation failed: %v", err)
			break
		}

		// Test with different message types
		for _, message := range testMessages {
			// Sign message
			signature, err := signing.Sign(privKey, message)
			if err != nil {
				allPassed = false
				errorMsg = fmt.Sprintf("Signing failed: %v", err)
				break
			}

			// Verify signature
			valid, err := signing.Verify(pubKey, message, signature)
			if err != nil {
				allPassed = false
				errorMsg = fmt.Sprintf("Verification failed: %v", err)
				break
			}

			if !valid {
				allPassed = false
				errorMsg = "Valid signature verification failed"
				break
			}

			// Test with wrong message (should fail)
			if len(message) > 0 {
				wrongMessage := make([]byte, len(message))
				copy(wrongMessage, message)
				wrongMessage[0] ^= 1

				wrongValid, err := signing.Verify(pubKey, wrongMessage, signature)
				if err != nil {
					allPassed = false
					errorMsg = fmt.Sprintf("Wrong message verification error: %v", err)
					break
				}

				if wrongValid {
					allPassed = false
					errorMsg = "Wrong message incorrectly verified"
					break
				}
			}
		}

		if !allPassed {
			break
		}

		// Verify key and signature sizes
		expectedPubSize, expectedPrivSize, expectedSigSize := signing.GetKeySizes(level)
		if len(pubKey) != expectedPubSize || len(privKey) != expectedPrivSize {
			allPassed = false
			errorMsg = "Unexpected key sizes"
			break
		}

		// Sign a test message to check signature size
		testSig, _ := signing.Sign(privKey, []byte("size test"))
		if len(testSig) != expectedSigSize {
			allPassed = false
			errorMsg = "Unexpected signature size"
			break
		}
	}

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("SIG-%s-BIST", algorithm),
		Algorithm:     algorithm,
		TestName:      "Algorithm Stress Test",
		Passed:        allPassed,
		ExecutionTime: time.Since(start),
		Iterations:    iterations,
		ErrorMessage:  errorMsg,
	})
}

// RunComprehensiveBIST runs the complete BIST suite with exit criteria
func (bs *BISTSuite) RunComprehensiveBIST() {
	fmt.Println("Starting Comprehensive BIST Suite...")
	fmt.Println(strings.Repeat("=", 80))

	// Phase 1: Generate test vectors
	fmt.Println("Phase 1: Generating test vectors...")
	bs.GenerateKEMTestVectors()
	bs.GenerateSignatureTestVectors()

	fmt.Printf("Generated %d test vectors total\n", len(bs.TestVectors))

	// Phase 2: Run KEM BIST
	fmt.Println("\nPhase 2: Running KEM BIST...")
	bs.RunKEMBIST()

	// Phase 3: Run Signature BIST
	fmt.Println("\nPhase 3: Running Signature BIST...")
	bs.RunSignatureBIST()

	// Phase 4: Cross-validation tests
	fmt.Println("\nPhase 4: Running cross-validation tests...")
	bs.runCrossValidationTests()

	// Phase 5: Performance regression tests
	fmt.Println("\nPhase 5: Running performance tests...")
	bs.runPerformanceTests()

	// Finalize
	bs.EndTime = time.Now()
	bs.evaluateExitCriteria()
}

// runCrossValidationTests ensures algorithms work correctly together
func (bs *BISTSuite) runCrossValidationTests() {
	start := time.Now()
	allPassed := true
	errorMsg := ""

	// Test that we can use KEM-generated keys with signature operations
	// and vice versa (ensuring no memory corruption or interference)

	securityLevels := []util.SecurityLevel{util.Level128, util.Level192, util.Level256}

	for _, level := range securityLevels {
		// Generate both KEM and signature keypairs
		kemPub, kemPriv, err := ciphering.GenerateKeyPair(level)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("KEM key generation failed: %v", err)
			break
		}

		sigPub, sigPriv, err := signing.GenerateKeyPair(level)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Signature key generation failed: %v", err)
			break
		}

		// Perform operations to ensure no interference
		ct, ss1, err := ciphering.Encapsulate(kemPub)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Encapsulation failed: %v", err)
			break
		}

		message := []byte("Cross validation test message")
		signature, err := signing.Sign(sigPriv, message)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Signing failed: %v", err)
			break
		}

		ss2, err := ciphering.Decapsulate(kemPriv, ct)
		if err != nil {
			allPassed = false
			errorMsg = fmt.Sprintf("Decapsulation failed: %v", err)
			break
		}

		valid, err := signing.Verify(sigPub, message, signature)
		if err != nil || !valid {
			allPassed = false
			errorMsg = fmt.Sprintf("Verification failed: %v", err)
			break
		}

		// Verify shared secrets match
		if subtle.ConstantTimeCompare(ss1, ss2) != 1 {
			allPassed = false
			errorMsg = "Shared secrets do not match in cross-validation"
			break
		}
	}

	bs.AddResult(BISTResult{
		TestID:        "CROSS-VAL-001",
		Algorithm:     "CROSS-VALIDATION",
		TestName:      "Algorithm Interference Test",
		Passed:        allPassed,
		ExecutionTime: time.Since(start),
		ErrorMessage:  errorMsg,
	})
}

// runPerformanceTests ensures algorithms meet performance requirements
func (bs *BISTSuite) runPerformanceTests() {
	// Performance thresholds (adjust based on requirements)
	thresholds := map[string]time.Duration{
		"Kyber512-KeyGen":   50 * time.Millisecond,
		"Kyber512-Encap":    20 * time.Millisecond,
		"Kyber512-Decap":    20 * time.Millisecond,
		"Kyber768-KeyGen":   100 * time.Millisecond,
		"Kyber1024-KeyGen":  150 * time.Millisecond,
		"Dilithium2-KeyGen": 200 * time.Millisecond,
		"Dilithium2-Sign":   100 * time.Millisecond,
		"Dilithium2-Verify": 50 * time.Millisecond,
		"Dilithium3-KeyGen": 300 * time.Millisecond,
		"Dilithium5-KeyGen": 500 * time.Millisecond,
	}

	testOperations := []struct {
		name      string
		level     util.SecurityLevel
		operation string
	}{
		{"Kyber512", util.Level128, "KEM"},
		{"Kyber768", util.Level192, "KEM"},
		{"Kyber1024", util.Level256, "KEM"},
		{"Dilithium2", util.Level128, "SIG"},
		{"Dilithium3", util.Level192, "SIG"},
		{"Dilithium5", util.Level256, "SIG"},
	}

	for _, test := range testOperations {
		if test.operation == "KEM" {
			bs.runKEMPerformanceTest(test.name, test.level, thresholds)
		} else {
			bs.runSignaturePerformanceTest(test.name, test.level, thresholds)
		}
	}
}

// runKEMPerformanceTest tests KEM algorithm performance
func (bs *BISTSuite) runKEMPerformanceTest(algorithm string, level util.SecurityLevel, thresholds map[string]time.Duration) {
	// Key generation performance
	start := time.Now()
	pubKey, privKey, err := ciphering.GenerateKeyPair(level)
	keyGenTime := time.Since(start)

	keyGenThreshold, exists := thresholds[algorithm+"-KeyGen"]
	passed := err == nil && (!exists || keyGenTime <= keyGenThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-KEYGEN", algorithm),
		Algorithm:     algorithm,
		TestName:      "Key Generation Performance",
		Passed:        passed,
		ExecutionTime: keyGenTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Key generation failed: %v", err)
			}
			if exists && keyGenTime > keyGenThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", keyGenTime, keyGenThreshold)
			}
			return ""
		}(),
	})

	if err != nil {
		return
	}

	// Encapsulation performance
	start = time.Now()
	ct, ss, err := ciphering.Encapsulate(pubKey)
	encapTime := time.Since(start)

	encapThreshold, exists := thresholds[algorithm+"-Encap"]
	passed = err == nil && (!exists || encapTime <= encapThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-ENCAP", algorithm),
		Algorithm:     algorithm,
		TestName:      "Encapsulation Performance",
		Passed:        passed,
		ExecutionTime: encapTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Encapsulation failed: %v", err)
			}
			if exists && encapTime > encapThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", encapTime, encapThreshold)
			}
			return ""
		}(),
	})

	if err != nil {
		return
	}

	// Decapsulation performance
	start = time.Now()
	_, err = ciphering.Decapsulate(privKey, ct)
	decapTime := time.Since(start)

	decapThreshold, exists := thresholds[algorithm+"-Decap"]
	passed = err == nil && (!exists || decapTime <= decapThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-DECAP", algorithm),
		Algorithm:     algorithm,
		TestName:      "Decapsulation Performance",
		Passed:        passed,
		ExecutionTime: decapTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Decapsulation failed: %v", err)
			}
			if exists && decapTime > decapThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", decapTime, decapThreshold)
			}
			return ""
		}(),
	})

	_ = ss // Avoid unused variable warning
}

// runSignaturePerformanceTest tests signature algorithm performance
func (bs *BISTSuite) runSignaturePerformanceTest(algorithm string, level util.SecurityLevel, thresholds map[string]time.Duration) {
	message := []byte("Performance test message")

	// Key generation performance
	start := time.Now()
	pubKey, privKey, err := signing.GenerateKeyPair(level)
	keyGenTime := time.Since(start)

	keyGenThreshold, exists := thresholds[algorithm+"-KeyGen"]
	passed := err == nil && (!exists || keyGenTime <= keyGenThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-KEYGEN", algorithm),
		Algorithm:     algorithm,
		TestName:      "Key Generation Performance",
		Passed:        passed,
		ExecutionTime: keyGenTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Key generation failed: %v", err)
			}
			if exists && keyGenTime > keyGenThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", keyGenTime, keyGenThreshold)
			}
			return ""
		}(),
	})

	if err != nil {
		return
	}

	// Signing performance
	start = time.Now()
	signature, err := signing.Sign(privKey, message)
	signTime := time.Since(start)

	signThreshold, exists := thresholds[algorithm+"-Sign"]
	passed = err == nil && (!exists || signTime <= signThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-SIGN", algorithm),
		Algorithm:     algorithm,
		TestName:      "Signing Performance",
		Passed:        passed,
		ExecutionTime: signTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Signing failed: %v", err)
			}
			if exists && signTime > signThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", signTime, signThreshold)
			}
			return ""
		}(),
	})

	if err != nil {
		return
	}

	// Verification performance
	start = time.Now()
	valid, err := signing.Verify(pubKey, message, signature)
	verifyTime := time.Since(start)

	verifyThreshold, exists := thresholds[algorithm+"-Verify"]
	passed = err == nil && valid && (!exists || verifyTime <= verifyThreshold)

	bs.AddResult(BISTResult{
		TestID:        fmt.Sprintf("PERF-%s-VERIFY", algorithm),
		Algorithm:     algorithm,
		TestName:      "Verification Performance",
		Passed:        passed,
		ExecutionTime: verifyTime,
		ErrorMessage: func() string {
			if err != nil {
				return fmt.Sprintf("Verification failed: %v", err)
			}
			if !valid {
				return "Signature verification returned false"
			}
			if exists && verifyTime > verifyThreshold {
				return fmt.Sprintf("Performance threshold exceeded: %v > %v", verifyTime, verifyThreshold)
			}
			return ""
		}(),
	})
}

// evaluateExitCriteria determines if the BIST suite has met exit criteria
func (bs *BISTSuite) evaluateExitCriteria() {
	// Exit criteria:
	// 1. All critical tests must pass
	// 2. At least 95% overall pass rate
	// 3. All test vectors must be validated
	// 4. No performance regressions beyond thresholds
	// 5. All algorithms must have generated and validated test vectors

	criticalTests := []string{
		"KEM-BIST-001",
		"SIG-BIST-001",
		"CROSS-VAL-001",
	}

	criticalTestsPassed := true
	for _, testID := range criticalTests {
		found := false
		for _, result := range bs.Results {
			if result.TestID == testID {
				found = true
				if !result.Passed {
					criticalTestsPassed = false
				}
				break
			}
		}
		if !found {
			criticalTestsPassed = false
		}
	}

	// Calculate overall pass rate
	passRate := 0.0
	if bs.TotalTests > 0 {
		passRate = float64(bs.PassedTests) / float64(bs.TotalTests)
	}

	// Check test vector requirements
	kemVectorCount := 0
	sigVectorCount := 0
	for _, tv := range bs.TestVectors {
		if strings.HasPrefix(tv.ID, "KEM-") {
			kemVectorCount++
		} else if strings.HasPrefix(tv.ID, "SIG-") {
			sigVectorCount++
		}
	}

	// Minimum requirements
	minKEMVectors := 30 // 5 vectors per algorithm * 3 algorithms * 2 (valid/invalid)
	minSigVectors := 45 // 3 algorithms * 5 messages * 3 test types

	vectorRequirementsMet := kemVectorCount >= minKEMVectors && sigVectorCount >= minSigVectors

	// Overall exit criteria evaluation
	bs.ExitCriteria = criticalTestsPassed &&
		passRate >= 0.95 &&
		vectorRequirementsMet &&
		len(bs.TestVectors) > 0

	fmt.Printf("\nExit Criteria Evaluation:\n")
	fmt.Printf("  Critical Tests Passed: %v\n", criticalTestsPassed)
	fmt.Printf("  Overall Pass Rate: %.1f%% (required: 95%%)\n", passRate*100)
	fmt.Printf("  Test Vectors Generated: %d (KEM: %d, SIG: %d)\n",
		len(bs.TestVectors), kemVectorCount, sigVectorCount)
	fmt.Printf("  Vector Requirements Met: %v\n", vectorRequirementsMet)
	fmt.Printf("  EXIT CRITERIA MET: %v\n", bs.ExitCriteria)
}

// PrintComprehensiveReport prints detailed BIST results
func (bs *BISTSuite) PrintComprehensiveReport() {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("POST-QUANTUM CRYPTOGRAPHY BIST COMPREHENSIVE REPORT")
	fmt.Println(strings.Repeat("=", 80))

	totalDuration := bs.EndTime.Sub(bs.StartTime)
	fmt.Printf("Execution Time: %v\n", totalDuration)
	fmt.Printf("Total Tests: %d\n", bs.TotalTests)
	fmt.Printf("Passed: %d\n", bs.PassedTests)
	fmt.Printf("Failed: %d\n", bs.FailedTests)
	fmt.Printf("Success Rate: %.1f%%\n", float64(bs.PassedTests)/float64(bs.TotalTests)*100)
	fmt.Printf("Test Vectors: %d\n", len(bs.TestVectors))
	fmt.Printf("Exit Criteria Met: %v\n", bs.ExitCriteria)

	fmt.Println("\nDETAILED RESULTS:")
	fmt.Println(strings.Repeat("-", 80))

	// Group results by category
	categories := map[string][]BISTResult{
		"KEM Algorithms":         {},
		"Signature Algorithms":   {},
		"Performance Tests":      {},
		"Cross-Validation":       {},
		"Test Vector Validation": {},
	}

	for _, result := range bs.Results {
		switch {
		case strings.Contains(result.TestID, "KEM-") && !strings.Contains(result.TestID, "PERF"):
			categories["KEM Algorithms"] = append(categories["KEM Algorithms"], result)
		case strings.Contains(result.TestID, "SIG-") && !strings.Contains(result.TestID, "PERF"):
			categories["Signature Algorithms"] = append(categories["Signature Algorithms"], result)
		case strings.Contains(result.TestID, "PERF-"):
			categories["Performance Tests"] = append(categories["Performance Tests"], result)
		case strings.Contains(result.TestID, "CROSS-"):
			categories["Cross-Validation"] = append(categories["Cross-Validation"], result)
		case strings.Contains(result.TestName, "Test Vector"):
			categories["Test Vector Validation"] = append(categories["Test Vector Validation"], result)
		}
	}

	for category, results := range categories {
		if len(results) == 0 {
			continue
		}

		fmt.Printf("\n%s:\n", category)
		for _, result := range results {
			status := "PASS"
			if !result.Passed {
				status = "FAIL"
			}

			fmt.Printf("  %-15s %-35s [%s] %8v",
				result.Algorithm,
				result.TestName,
				status,
				result.ExecutionTime)

			if result.Iterations > 0 {
				fmt.Printf(" (%d iter)", result.Iterations)
			}
			if result.TestVectors > 0 {
				fmt.Printf(" (%d vectors)", result.TestVectors)
			}

			fmt.Println()

			if !result.Passed && result.ErrorMessage != "" {
				fmt.Printf("    Error: %s\n", result.ErrorMessage)
			}
		}
	}

	// Test vector summary
	fmt.Printf("\nTEST VECTOR SUMMARY:\n")
	fmt.Println(strings.Repeat("-", 40))

	kemCount := 0
	sigCount := 0
	validCount := 0
	invalidCount := 0

	for _, tv := range bs.TestVectors {
		if strings.HasPrefix(tv.ID, "KEM-") {
			kemCount++
		} else if strings.HasPrefix(tv.ID, "SIG-") {
			sigCount++
		}

		if tv.ExpectedResult {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Printf("KEM Test Vectors: %d\n", kemCount)
	fmt.Printf("Signature Test Vectors: %d\n", sigCount)
	fmt.Printf("Valid Cases: %d\n", validCount)
	fmt.Printf("Invalid Cases: %d\n", invalidCount)
	fmt.Printf("Total: %d\n", len(bs.TestVectors))

	fmt.Println(strings.Repeat("=", 80))
}

// SaveTestVectors saves test vectors to JSON file
func (bs *BISTSuite) SaveTestVectors(filename string) error {
	data, err := json.MarshalIndent(bs.TestVectors, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal test vectors: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write test vectors to file: %w", err)
	}

	fmt.Printf("✅ Saved test vectors to: %s\n", filename)
	return nil
}

// SaveBISTReport saves complete BIST report to JSON
func (bs *BISTSuite) SaveBISTReport(filename string) error {
	data, err := json.MarshalIndent(bs, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal BIST report: %w", err)
	}

	// Write to file
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("failed to write BIST report to file: %w", err)
	}

	fmt.Printf("✅ Saved BIST report to: %s (%d bytes)\n", filename, len(data))
	return nil
}

// GetExitCode returns appropriate exit code based on BIST results
func (bs *BISTSuite) GetExitCode() int {
	if bs.ExitCriteria {
		return 0 // Success
	}
	return 1 // Failure
}

func (bs *BISTSuite) GetErrors() []string {
	return bs.Errors
}

// ValidateTestVector validates a single test vector
func ValidateTestVector(tv TestVector) error {
	switch {
	case strings.HasPrefix(tv.ID, "KEM-"):
		return validateKEMTestVector(tv)
	case strings.HasPrefix(tv.ID, "SIG-"):
		return validateSignatureTestVector(tv)
	default:
		return fmt.Errorf("unknown test vector type: %s", tv.ID)
	}
}

// validateKEMTestVector validates a KEM test vector
func validateKEMTestVector(tv TestVector) error {
	// Decode public key
	_, err := hex.DecodeString(tv.PublicKey) // _ := pubKey
	if err != nil {
		return fmt.Errorf("invalid public key hex: %w", err)
	}

	// Decode private key
	privKey, err := hex.DecodeString(tv.PrivateKey)
	if err != nil {
		return fmt.Errorf("invalid private key hex: %w", err)
	}

	// Decode ciphertext
	ciphertext, err := hex.DecodeString(tv.Ciphertext)
	if err != nil {
		return fmt.Errorf("invalid ciphertext hex: %w", err)
	}

	// Attempt decapsulation
	sharedSecret, err := ciphering.Decapsulate(privKey, ciphertext)

	if tv.ExpectedResult {
		if err != nil {
			return fmt.Errorf("expected valid decapsulation but got error: %w", err)
		}

		// Compare with expected shared secret if provided
		if tv.SharedSecret != "" {
			expectedSecret, err := hex.DecodeString(tv.SharedSecret)
			if err != nil {
				return fmt.Errorf("invalid expected shared secret hex: %w", err)
			}

			if subtle.ConstantTimeCompare(sharedSecret, expectedSecret) != 1 {
				return fmt.Errorf("shared secret mismatch")
			}
		}
	} else {
		// For invalid test vectors, we don't necessarily expect an error,
		// just that the result (if any) doesn't match the expected value
		if tv.SharedSecret != "" && err == nil {
			expectedSecret, err := hex.DecodeString(tv.SharedSecret)
			if err == nil && subtle.ConstantTimeCompare(sharedSecret, expectedSecret) == 1 {
				return fmt.Errorf("unexpected shared secret match for invalid test case")
			}
		}
	}

	return nil
}

// validateSignatureTestVector validates a signature test vector
func validateSignatureTestVector(tv TestVector) error {
	// Decode public key
	pubKey, err := hex.DecodeString(tv.PublicKey)
	if err != nil {
		return fmt.Errorf("invalid public key hex: %w", err)
	}

	// Decode message
	message, err := hex.DecodeString(tv.Message)
	if err != nil {
		return fmt.Errorf("invalid message hex: %w", err)
	}

	// Decode signature
	signature, err := hex.DecodeString(tv.Signature)
	if err != nil {
		return fmt.Errorf("invalid signature hex: %w", err)
	}

	// Verify signature
	valid, err := signing.Verify(pubKey, message, signature)
	if err != nil {
		return fmt.Errorf("verification error: %w", err)
	}

	if valid != tv.ExpectedResult {
		return fmt.Errorf("verification result mismatch: got %v, expected %v",
			valid, tv.ExpectedResult)
	}

	return nil
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
