package signing

import (
	"fmt"

	"github.com/cloudflare/circl/sign"
	"github.com/cloudflare/circl/sign/dilithium/mode2"
	"github.com/cloudflare/circl/sign/dilithium/mode3"
	"github.com/cloudflare/circl/sign/dilithium/mode5"

	"pqc_bist_demo/util"
)

// getScheme returns the appropriate Dilithium scheme based on security level
func getScheme(level util.SecurityLevel) sign.Scheme {
	switch level {
	case util.Level128:
		return mode2.Scheme()
	case util.Level192:
		return mode3.Scheme()
	case util.Level256:
		return mode5.Scheme()
	default:
		return mode3.Scheme() // Default to Level 3
	}
}

// GetAlgorithmName returns the algorithm name for the given security level
func GetAlgorithmName(level util.SecurityLevel) string {
	switch level {
	case util.Level128:
		return "Dilithium2"
	case util.Level192:
		return "Dilithium3"
	case util.Level256:
		return "Dilithium5"
	default:
		return "Dilithium3"
	}
}

// GenerateKeyPair generates a new signing key pair for the specified security level
func GenerateKeyPair(level util.SecurityLevel) (publicKey []byte, privateKey []byte, err error) {
	scheme := getScheme(level)

	pubKey, privKey, err := scheme.GenerateKey()
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

// Sign creates a digital signature for the given message
func Sign(privateKey []byte, message []byte) (signature []byte, err error) {
	// Detect security level based on private key size
	level := detectSecurityLevelFromPrivateKey(len(privateKey))
	scheme := getScheme(level)

	// Unmarshal the private key
	privKey, err := scheme.UnmarshalBinaryPrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	// Sign the message
	sig := scheme.Sign(privKey, message, nil) // nil for no randomizer

	return sig, nil
}

// Verify checks if a signature is valid for the given message and public key
func Verify(publicKey []byte, message []byte, signature []byte) (bool, error) {
	// Detect security level based on public key size
	level := detectSecurityLevel(len(publicKey))
	scheme := getScheme(level)

	// Unmarshal the public key
	pubKey, err := scheme.UnmarshalBinaryPublicKey(publicKey)
	if err != nil {
		return false, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	// Verify the signature
	valid := scheme.Verify(pubKey, message, signature, nil) // nil for no randomizer

	return valid, nil
}

// detectSecurityLevel determines security level based on public key size
func detectSecurityLevel(pubKeySize int) util.SecurityLevel {
	switch pubKeySize {
	case 1312: // Dilithium2
		return util.Level128
	case 1952: // Dilithium3
		return util.Level192
	case 2592: // Dilithium5
		return util.Level256
	default:
		return util.Level192 // Default
	}
}

// detectSecurityLevelFromPrivateKey determines security level based on private key size
func detectSecurityLevelFromPrivateKey(privKeySize int) util.SecurityLevel {
	switch privKeySize {
	case 2528: // Dilithium2
		return util.Level128
	case 4000: // Dilithium3
		return util.Level192
	case 4864: // Dilithium5
		return util.Level256
	default:
		return util.Level192 // Default
	}
}

// GetKeySizes returns the expected key and signature sizes for a security level
func GetKeySizes(level util.SecurityLevel) (pubKeySize, privKeySize, signatureSize int) {
	switch level {
	case util.Level128: // Dilithium2
		return 1312, 2528, 2420
	case util.Level192: // Dilithium3
		return 1952, 4000, 3293
	case util.Level256: // Dilithium5
		return 2592, 4864, 4595
	default:
		return 1952, 4000, 3293 // Default to Level 3
	}
}
