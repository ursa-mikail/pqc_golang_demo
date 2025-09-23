package hashing

import (
	"bytes"
	"testing"

	"pqc_bist_demo/util"
)

func TestGetAlgorithm(t *testing.T) {
	tests := []struct {
		level    util.SecurityLevel
		expected string
	}{
		{util.Level128, "SHAKE128"},
		{util.Level192, "SHAKE256"},
		{util.Level256, "SHA3-256"},
	}

	for _, test := range tests {
		result := GetAlgorithm(test.level)
		if result != test.expected {
			t.Errorf("GetAlgorithm(%v) = %v, want %v", test.level, result, test.expected)
		}
	}
}

func TestHash(t *testing.T) {
	testData := []byte("test data for hashing")

	levels := []util.SecurityLevel{
		util.Level128,
		util.Level192,
		util.Level256,
	}

	for _, level := range levels {
		hash, err := Hash(testData, level)
		if err != nil {
			t.Errorf("Hash failed for level %v: %v", level, err)
			continue
		}

		if len(hash) == 0 {
			t.Errorf("Hash returned empty result for level %v", level)
			continue
		}

		// Test that the same input produces the same hash
		hash2, err := Hash(testData, level)
		if err != nil {
			t.Errorf("Second hash failed for level %v: %v", level, err)
			continue
		}

		if !bytes.Equal(hash, hash2) {
			t.Errorf("Hash is not deterministic for level %v", level)
		}

		// Test that different inputs produce different hashes
		differentData := []byte("different test data")
		hash3, err := Hash(differentData, level)
		if err != nil {
			t.Errorf("Hash with different data failed for level %v: %v", level, err)
			continue
		}

		if bytes.Equal(hash, hash3) {
			t.Errorf("Same hash produced for different inputs at level %v", level)
		}
	}
}

func TestHashWithSize(t *testing.T) {
	testData := []byte("test data for variable size hashing")

	tests := []struct {
		level      util.SecurityLevel
		outputSize int
		shouldFail bool
	}{
		{util.Level128, 16, false},
		{util.Level128, 32, false},
		{util.Level128, 64, false},
		{util.Level192, 16, false},
		{util.Level192, 32, false},
		{util.Level192, 64, false},
		{util.Level256, 32, false}, // SHA3-256 fixed size
		{util.Level256, 64, true},  // SHA3-256 doesn't support variable size
	}

	for _, test := range tests {
		hash, err := HashWithSize(testData, test.level, test.outputSize)

		if test.shouldFail {
			if err == nil {
				t.Errorf("Expected error for level %v with size %d, but got none", test.level, test.outputSize)
			}
			continue
		}

		if err != nil {
			t.Errorf("HashWithSize failed for level %v, size %d: %v", test.level, test.outputSize, err)
			continue
		}

		if len(hash) != test.outputSize {
			t.Errorf("Expected hash size %d, got %d for level %v", test.outputSize, len(hash), test.level)
		}
	}
}

func TestGetDefaultOutputSize(t *testing.T) {
	tests := []struct {
		level    util.SecurityLevel
		expected int
	}{
		{util.Level128, 16},
		{util.Level192, 32},
		{util.Level256, 32},
	}

	for _, test := range tests {
		result := GetDefaultOutputSize(test.level)
		if result != test.expected {
			t.Errorf("GetDefaultOutputSize(%v) = %d, want %d", test.level, result, test.expected)
		}
	}
}

func TestVerifyHash(t *testing.T) {
	hash1 := []byte{1, 2, 3, 4, 5}
	hash2 := []byte{1, 2, 3, 4, 5}
	hash3 := []byte{1, 2, 3, 4, 6}

	if !VerifyHash(hash1, hash2) {
		t.Error("VerifyHash should return true for identical hashes")
	}

	if VerifyHash(hash1, hash3) {
		t.Error("VerifyHash should return false for different hashes")
	}
}

func TestHashChain(t *testing.T) {
	data := [][]byte{
		[]byte("first"),
		[]byte("second"),
		[]byte("third"),
	}

	levels := []util.SecurityLevel{
		util.Level128,
		util.Level192,
		util.Level256,
	}

	for _, level := range levels {
		hash, err := HashChain(data, level)
		if err != nil {
			t.Errorf("HashChain failed for level %v: %v", level, err)
			continue
		}

		if len(hash) == 0 {
			t.Errorf("HashChain returned empty result for level %v", level)
		}

		// Test with empty data
		_, err = HashChain([][]byte{}, level)
		if err == nil {
			t.Errorf("HashChain should fail with empty data")
		}

		// Test that order matters
		reorderedData := [][]byte{
			[]byte("second"),
			[]byte("first"),
			[]byte("third"),
		}

		hash2, err := HashChain(reorderedData, level)
		if err != nil {
			t.Errorf("HashChain failed with reordered data for level %v: %v", level, err)
			continue
		}

		if bytes.Equal(hash, hash2) {
			t.Errorf("HashChain should produce different results for different order")
		}
	}
}

func BenchmarkHash(b *testing.B) {
	testData := []byte("benchmark test data that is reasonably long to get meaningful timing results")
	level := util.Level256

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Hash(testData, level)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSHAKE128(b *testing.B) {
	testData := []byte("benchmark test data that is reasonably long to get meaningful timing results")
	level := util.Level128

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Hash(testData, level)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkSHAKE256(b *testing.B) {
	testData := []byte("benchmark test data that is reasonably long to get meaningful timing results")
	level := util.Level192

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := Hash(testData, level)
		if err != nil {
			b.Fatal(err)
		}
	}
}
