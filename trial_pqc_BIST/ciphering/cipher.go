package ciphering

import (
	"fmt"

	"github.com/cloudflare/circl/kem"
	"github.com/cloudflare/circl/kem/kyber/kyber1024"
	"github.com/cloudflare/circl/kem/kyber/kyber512"
	"github.com/cloudflare/circl/kem/kyber/kyber768"

	"pqc_bist_demo/util"
)

// getScheme returns the appropriate Kyber scheme based on security level
func getScheme(level util.SecurityLevel) kem.Scheme {
	switch level {
	case util.Level128:
		return kyber512.Scheme()
	case util.Level192:
		return kyber768.Scheme()
	case util.Level256:
		return kyber1024.Scheme()
	default:
		return kyber768.Scheme() // Default to Level 3
	}
}

// GetAlgorithmName returns the algorithm name for the given security level
func GetAlgorithmName(level util.SecurityLevel) string {
	switch level {
	case util.Level128:
		return "Kyber512"
	case util.Level192:
		return "Kyber768"
	case util.Level256:
		return "Kyber1024"
	default:
		return "Kyber768"
	}
}

// GenerateKeyPair generates a new key pair for the specified security level
func GenerateKeyPair(level util.SecurityLevel) (publicKey []byte, privateKey []byte, err error) {
	scheme := getScheme(level)

	pubKey, privKey, err := scheme.GenerateKeyPair()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to generate keypair: %w", err)
	}

	// Marshal keys to byte slices
	publicKey, err = pubKey.MarshalBinary()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	privateKey, err = privKey.MarshalBinary()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to marshal private key: %w", err)
	}

	return publicKey, privateKey, nil
}

// Encapsulate creates a shared secret and ciphertext using the public key
func Encapsulate(publicKey []byte) (ciphertext []byte, sharedSecret []byte, err error) {
	// We need to determine which scheme was used based on key size
	level := detectSecurityLevel(len(publicKey))
	scheme := getScheme(level)

	// Unmarshal the public key
	pubKey, err := scheme.UnmarshalBinaryPublicKey(publicKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	// Encapsulate
	ct, ss, err := scheme.Encapsulate(pubKey)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to encapsulate: %w", err)
	}

	return ct, ss, nil
}

// Decapsulate recovers the shared secret using the private key and ciphertext
func Decapsulate(privateKey []byte, ciphertext []byte) (sharedSecret []byte, err error) {
	// Detect security level based on private key size
	level := detectSecurityLevelFromPrivateKey(len(privateKey))
	scheme := getScheme(level)

	// Unmarshal the private key
	privKey, err := scheme.UnmarshalBinaryPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	// Decapsulate
	ss, err := scheme.Decapsulate(privKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("failed to decapsulate: %w", err)
	}

	return ss, nil
}

// detectSecurityLevel determines security level based on public key size
func detectSecurityLevel(pubKeySize int) util.SecurityLevel {
	switch pubKeySize {
	case 800: // Kyber512
		return util.Level128
	case 1184: // Kyber768
		return util.Level192
	case 1568: // Kyber1024
		return util.Level256
	default:
		return util.Level192 // Default
	}
}

// detectSecurityLevelFromPrivateKey determines security level based on private key size
func detectSecurityLevelFromPrivateKey(privKeySize int) util.SecurityLevel {
	switch privKeySize {
	case 1632: // Kyber512
		return util.Level128
	case 2400: // Kyber768
		return util.Level192
	case 3168: // Kyber1024
		return util.Level256
	default:
		return util.Level192 // Default
	}
}

// GetKeySizes returns the expected key and ciphertext sizes for a security level
func GetKeySizes(level util.SecurityLevel) (pubKeySize, privKeySize, ciphertextSize, sharedSecretSize int) {
	switch level {
	case util.Level128: // Kyber512
		return 800, 1632, 768, 32
	case util.Level192: // Kyber768
		return 1184, 2400, 1088, 32
	case util.Level256: // Kyber1024
		return 1568, 3168, 1568, 32
	default:
		return 1184, 2400, 1088, 32 // Default to Level 3
	}
}
