package signing

import (
	"testing"

	"trial_pqc/util"
)

func TestGenerateKeyPair(t *testing.T) {
	levels := []util.SecurityLevel{util.Level128, util.Level192, util.Level256}

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			pubKey, privKey, err := GenerateKeyPair(level)
			if err != nil {
				t.Fatalf("GenerateKeyPair failed: %v", err)
			}

			// Check key sizes match expected values
			expectedPubSize, expectedPrivSize, _ := GetKeySizes(level)

			if len(pubKey) != expectedPubSize {
				t.Errorf("Public key size = %d, want %d", len(pubKey), expectedPubSize)
			}

			if len(privKey) != expectedPrivSize {
				t.Errorf("Private key size = %d, want %d", len(privKey), expectedPrivSize)
			}

			// Keys should not be empty
			if len(pubKey) == 0 || len(privKey) == 0 {
				t.Error("Generated keys should not be empty")
			}
		})
	}
}

func TestSignVerify(t *testing.T) {
	levels := []util.SecurityLevel{util.Level128, util.Level192, util.Level256}
	message := []byte("Hello, Post-Quantum World!")

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			// Generate keypair
			pubKey, privKey, err := GenerateKeyPair(level)
			if err != nil {
				t.Fatalf("GenerateKeyPair failed: %v", err)
			}

			// Sign message
			signature, err := Sign(privKey, message)
			if err != nil {
				t.Fatalf("Sign failed: %v", err)
			}

			// Check signature size
			_, _, expectedSigSize := GetKeySizes(level)
			if len(signature) != expectedSigSize {
				t.Errorf("Signature size = %d, want %d", len(signature), expectedSigSize)
			}

			// Verify signature
			valid, err := Verify(pubKey, message, signature)
			if err != nil {
				t.Fatalf("Verify failed: %v", err)
			}

			if !valid {
				t.Error("Signature verification failed")
			}
		})
	}
}

func TestVerifyWrongMessage(t *testing.T) {
	level := util.Level192
	message := []byte("Original message")
	wrongMessage := []byte("Wrong message")

	// Generate keypair and sign
	pubKey, privKey, err := GenerateKeyPair(level)
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	signature, err := Sign(privKey, message)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	// Verify with wrong message should fail
	valid, err := Verify(pubKey, wrongMessage, signature)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if valid {
		t.Error("Signature verification should have failed with wrong message")
	}
}

func TestVerifyWrongSignature(t *testing.T) {
	level := util.Level192
	message := []byte("Test message")

	// Generate keypair
	pubKey, privKey, err := GenerateKeyPair(level)
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	// Create valid signature
	signature, err := Sign(privKey, message)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	// Corrupt signature
	corruptSignature := make([]byte, len(signature))
	copy(corruptSignature, signature)
	corruptSignature[0] ^= 0xFF // Flip bits

	// Verify corrupted signature should fail
	valid, err := Verify(pubKey, message, corruptSignature)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if valid {
		t.Error("Signature verification should have failed with corrupted signature")
	}
}

func TestGetAlgorithmName(t *testing.T) {
	tests := []struct {
		level    util.SecurityLevel
		expected string
	}{
		{util.Level128, "Dilithium2"},
		{util.Level192, "Dilithium3"},
		{util.Level256, "Dilithium5"},
	}

	for _, test := range tests {
		result := GetAlgorithmName(test.level)
		if result != test.expected {
			t.Errorf("GetAlgorithmName(%v) = %q, want %q",
				test.level, result, test.expected)
		}
	}
}

func TestGetKeySizes(t *testing.T) {
	tests := []struct {
		level                      util.SecurityLevel
		pubSize, privSize, sigSize int
	}{
		{util.Level128, 1312, 2528, 2420}, // Dilithium2
		{util.Level192, 1952, 4000, 3293}, // Dilithium3
		{util.Level256, 2592, 4864, 4595}, // Dilithium5
	}

	for _, test := range tests {
		pubSize, privSize, sigSize := GetKeySizes(test.level)

		if pubSize != test.pubSize {
			t.Errorf("GetKeySizes(%v) pub size = %d, want %d",
				test.level, pubSize, test.pubSize)
		}
		if privSize != test.privSize {
			t.Errorf("GetKeySizes(%v) priv size = %d, want %d",
				test.level, privSize, test.privSize)
		}
		if sigSize != test.sigSize {
			t.Errorf("GetKeySizes(%v) sig size = %d, want %d",
				test.level, sigSize, test.sigSize)
		}
	}
}

func TestInvalidKeys(t *testing.T) {
	message := []byte("test")

	// Test with invalid private key
	_, err := Sign([]byte("invalid"), message)
	if err == nil {
		t.Error("Expected error with invalid private key")
	}

	// Test with invalid public key
	_, err = Verify([]byte("invalid"), message, []byte("signature"))
	if err == nil {
		t.Error("Expected error with invalid public key")
	}
}

func TestEmptyMessage(t *testing.T) {
	level := util.Level192
	message := []byte{}

	// Generate keypair
	pubKey, privKey, err := GenerateKeyPair(level)
	if err != nil {
		t.Fatalf("GenerateKeyPair failed: %v", err)
	}

	// Sign empty message
	signature, err := Sign(privKey, message)
	if err != nil {
		t.Fatalf("Sign failed: %v", err)
	}

	// Verify empty message
	valid, err := Verify(pubKey, message, signature)
	if err != nil {
		t.Fatalf("Verify failed: %v", err)
	}

	if !valid {
		t.Error("Signature verification failed for empty message")
	}
}

// Benchmark key generation
func BenchmarkGenerateKeyPair(b *testing.B) {
	levels := []util.SecurityLevel{util.Level128, util.Level192, util.Level256}

	for _, level := range levels {
		b.Run(GetAlgorithmName(level), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _, err := GenerateKeyPair(level)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Benchmark signing
func BenchmarkSign(b *testing.B) {
	message := []byte("Benchmark message for signing performance")
	_, privKey, err := GenerateKeyPair(util.Level192)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Sign(privKey, message)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark verification
func BenchmarkVerify(b *testing.B) {
	message := []byte("Benchmark message for verification performance")
	pubKey, privKey, err := GenerateKeyPair(util.Level192)
	if err != nil {
		b.Fatal(err)
	}

	signature, err := Sign(privKey, message)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Verify(pubKey, message, signature)
		if err != nil {
			b.Fatal(err)
		}
	}
}
