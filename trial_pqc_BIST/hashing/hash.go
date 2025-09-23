package hashing

import (
	"fmt"

	"golang.org/x/crypto/sha3"

	"pqc_bist_demo/util"
)

// GetAlgorithm returns the recommended hash algorithm for the security level
func GetAlgorithm(level util.SecurityLevel) string {
	switch level {
	case util.Level128:
		return "SHAKE128"
	case util.Level192:
		return "SHAKE256"
	case util.Level256:
		return "SHA3-256"
	default:
		return "SHAKE256"
	}
}

// Hash computes a hash of the input data using the appropriate algorithm for the security level
func Hash(data []byte, level util.SecurityLevel) ([]byte, error) {
	switch level {
	case util.Level128:
		return hashSHAKE128(data), nil
	case util.Level192:
		return hashSHAKE256(data), nil
	case util.Level256:
		return hashSHA3_256(data), nil
	default:
		return hashSHAKE256(data), nil
	}
}

// HashWithSize computes a hash with a specific output size (useful for SHAKE functions)
func HashWithSize(data []byte, level util.SecurityLevel, outputSize int) ([]byte, error) {
	switch level {
	case util.Level128:
		return hashSHAKE128WithSize(data, outputSize), nil
	case util.Level192:
		return hashSHAKE256WithSize(data, outputSize), nil
	case util.Level256:
		if outputSize != 32 {
			return nil, fmt.Errorf("SHA3-256 has fixed output size of 32 bytes")
		}
		return hashSHA3_256(data), nil
	default:
		return hashSHAKE256WithSize(data, outputSize), nil
	}
}

// GetDefaultOutputSize returns the default output size for each algorithm
func GetDefaultOutputSize(level util.SecurityLevel) int {
	switch level {
	case util.Level128:
		return 16 // 128 bits
	case util.Level192:
		return 32 // 256 bits (SHAKE256 commonly used with 256-bit output)
	case util.Level256:
		return 32 // 256 bits
	default:
		return 32
	}
}

// hashSHAKE128 computes SHAKE128 hash with default 128-bit output
func hashSHAKE128(data []byte) []byte {
	return hashSHAKE128WithSize(data, 16) // 16 bytes = 128 bits
}

// hashSHAKE128WithSize computes SHAKE128 hash with custom output size
func hashSHAKE128WithSize(data []byte, outputSize int) []byte {
	hash := make([]byte, outputSize)
	shake := sha3.NewShake128()
	shake.Write(data)
	shake.Read(hash)
	return hash
}

// hashSHAKE256 computes SHAKE256 hash with default 256-bit output
func hashSHAKE256(data []byte) []byte {
	return hashSHAKE256WithSize(data, 32) // 32 bytes = 256 bits
}

// hashSHAKE256WithSize computes SHAKE256 hash with custom output size
func hashSHAKE256WithSize(data []byte, outputSize int) []byte {
	hash := make([]byte, outputSize)
	shake := sha3.NewShake256()
	shake.Write(data)
	shake.Read(hash)
	return hash
}

// hashSHA3_256 computes SHA3-256 hash (fixed 256-bit output)
func hashSHA3_256(data []byte) []byte {
	hash := sha3.Sum256(data)
	return hash[:]
}

// HMAC-like constructions for post-quantum security
// Note: In practice, you'd want to use proper KMAC (Keccak-based MAC) for PQ security

// VerifyHash performs a secure comparison of two hashes
func VerifyHash(hash1, hash2 []byte) bool {
	return util.SecureCompare(hash1, hash2)
}

// HashChain creates a hash chain (useful for Merkle trees, etc.)
func HashChain(data [][]byte, level util.SecurityLevel) ([]byte, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("no data to hash")
	}

	// Start with the first piece of data
	result, err := Hash(data[0], level)
	if err != nil {
		return nil, err
	}

	// Chain hash the remaining data
	for i := 1; i < len(data); i++ {
		// Concatenate current hash with next data and hash again
		combined := append(result, data[i]...)
		result, err = Hash(combined, level)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
