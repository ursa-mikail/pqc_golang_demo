package ciphering

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
			expectedPubSize, expectedPrivSize, _, _ := GetKeySizes(level)

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

func TestEncapsulateDecapsulate(t *testing.T) {
	levels := []util.SecurityLevel{util.Level128, util.Level192, util.Level256}

	for _, level := range levels {
		t.Run(level.String(), func(t *testing.T) {
			// Generate keypair
			pubKey, privKey, err := GenerateKeyPair(level)
			if err != nil {
				t.Fatalf("GenerateKeyPair failed: %v", err)
			}

			// Encapsulate
			ciphertext, sharedSecret1, err := Encapsulate(pubKey)
			if err != nil {
				t.Fatalf("Encapsulate failed: %v", err)
			}

			// Check sizes
			_, _, expectedCtSize, expectedSsSize := GetKeySizes(level)
			if len(ciphertext) != expectedCtSize {
				t.Errorf("Ciphertext size = %d, want %d", len(ciphertext), expectedCtSize)
			}
			if len(sharedSecret1) != expectedSsSize {
				t.Errorf("Shared secret size = %d, want %d", len(sharedSecret1), expectedSsSize)
			}

			// Decapsulate
			sharedSecret2, err := Decapsulate(privKey, ciphertext)
			if err != nil {
				t.Fatalf("Decapsulate failed: %v", err)
			}

			// Shared secrets should match
			if !util.SecureCompare(sharedSecret1, sharedSecret2) {
				t.Error("Shared secrets do not match")
			}
		})
	}
}

func TestGetAlgorithmName(t *testing.T) {
	tests := []struct {
		level    util.SecurityLevel
		expected string
	}{
		{util.Level128, "Kyber512"},
		{util.Level192, "Kyber768"},
		{util.Level256, "Kyber1024"},
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
		level             util.SecurityLevel
		pubSize, privSize int
		ctSize, ssSize    int
	}{
		{util.Level128, 800, 1632, 768, 32},   // Kyber512
		{util.Level192, 1184, 2400, 1088, 32}, // Kyber768
		{util.Level256, 1568, 3168, 1568, 32}, // Kyber1024
	}

	for _, test := range tests {
		pubSize, privSize, ctSize, ssSize := GetKeySizes(test.level)

		if pubSize != test.pubSize {
			t.Errorf("GetKeySizes(%v) pub size = %d, want %d",
				test.level, pubSize, test.pubSize)
		}
		if privSize != test.privSize {
			t.Errorf("GetKeySizes(%v) priv size = %d, want %d",
				test.level, privSize, test.privSize)
		}
		if ctSize != test.ctSize {
			t.Errorf("GetKeySizes(%v) ct size = %d, want %d",
				test.level, ctSize, test.ctSize)
		}
		if ssSize != test.ssSize {
			t.Errorf("GetKeySizes(%v) ss size = %d, want %d",
				test.level, ssSize, test.ssSize)
		}
	}
}

func TestInvalidKeys(t *testing.T) {
	// Test with invalid public key
	_, _, err := Encapsulate([]byte("invalid"))
	if err == nil {
		t.Error("Expected error with invalid public key")
	}

	// Test with invalid private key
	_, err = Decapsulate([]byte("invalid"), []byte("invalid"))
	if err == nil {
		t.Error("Expected error with invalid private key")
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

// Benchmark encapsulation
func BenchmarkEncapsulate(b *testing.B) {
	pubKey, _, err := GenerateKeyPair(util.Level192)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := Encapsulate(pubKey)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark decapsulation
func BenchmarkDecapsulate(b *testing.B) {
	pubKey, privKey, err := GenerateKeyPair(util.Level192)
	if err != nil {
		b.Fatal(err)
	}

	ciphertext, _, err := Encapsulate(pubKey)
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Decapsulate(privKey, ciphertext)
		if err != nil {
			b.Fatal(err)
		}
	}
}
