package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/cloudflare/circl/kem/xwing"
)

// 1. Key Exchange Phase (Using X-Wing KEM)
// Alice (Key Generator)
func aliceSetup() ([]byte, []byte, error) {
	// Generate key pair
	keySeed := make([]byte, 32)
	if _, err := rand.Read(keySeed); err != nil {
		return nil, nil, err
	}

	sk, pk := xwing.DeriveKeyPairPacked(keySeed)
	return sk, pk, nil
}

// Bob (Encapsulator)
func bobSendMessage(pk []byte, message string) ([]byte, []byte, []byte, error) {
	// Generate encapsulation seed
	encapSeed := make([]byte, 64)
	if _, err := rand.Read(encapSeed); err != nil {
		return nil, nil, nil, err
	}

	// Encapsulate to get shared secret and ciphertext
	sharedSecret, ct, err := xwing.Encapsulate(pk, encapSeed)
	if err != nil {
		return nil, nil, nil, err
	}

	// Use shared secret for AES-256-GCM
	encryptedMessage, nonce, err := encryptWithAESGCM(sharedSecret, []byte(message))
	if err != nil {
		return nil, nil, nil, err
	}

	return ct, encryptedMessage, nonce, nil
}

// Alice (Decapsulator)
func aliceReceiveMessage(sk, ct, encryptedMessage, nonce []byte) (string, error) {
	// Decapsulate to get shared secret
	sharedSecret := xwing.Decapsulate(ct, sk)

	// Decrypt with AES-256-GCM
	decrypted, err := decryptWithAESGCM(sharedSecret, nonce, encryptedMessage)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// 2. AES-256-GCM Encryption/Decryption Functions
// Derive AES key from KEM shared secret
func deriveAESKey(sharedSecret []byte) []byte {
	// In practice, you might want to use HKDF or similar
	// For simplicity, we'll use first 32 bytes of shared secret
	// Real implementations should use proper key derivation
	aesKey := make([]byte, 32)
	copy(aesKey, sharedSecret[:32])
	return aesKey
}

func encryptWithAESGCM(sharedSecret, plaintext []byte) ([]byte, []byte, error) {
	aesKey := deriveAESKey(sharedSecret)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	// Encrypt and authenticate
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

func decryptWithAESGCM(sharedSecret, nonce, ciphertext []byte) ([]byte, error) {
	aesKey := deriveAESKey(sharedSecret)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Decrypt and verify authentication
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// main flow

func main() {
	fmt.Println("=== Alice and Bob: X-Wing KEM + AES-256-GCM ===\n")

	// Step 1: Alice generates key pair
	fmt.Println("1. KEY EXCHANGE")
	alicePrivateKey, alicePublicKey, err := aliceSetup()
	if err != nil {
		panic(err)
	}
	fmt.Printf("   Alice generated key pair\n")
	fmt.Printf("   Public key size: %d bytes\n", len(alicePublicKey))

	// Step 2: Bob sends encrypted message
	fmt.Println("\n2. ENCRYPTION")
	message := "Hello Alice! This is a secret message."
	fmt.Printf("   Original message: %q\n", message)

	ciphertext, encryptedMessage, nonce, err := bobSendMessage(alicePublicKey, message)
	if err != nil {
		panic(err)
	}
	fmt.Printf("   Bob encrypted message using Alice's public key\n")
	fmt.Printf("   KEM ciphertext: %d bytes\n", len(ciphertext))
	fmt.Printf("   AES-GCM encrypted message: %d bytes\n", len(encryptedMessage))

	// Step 3: Alice decrypts the message
	fmt.Println("\n3. DECRYPTION")
	decrypted, err := aliceReceiveMessage(alicePrivateKey, ciphertext, encryptedMessage, nonce)
	if err != nil {
		panic(err)
	}
	fmt.Printf("   Alice decrypted message: %q\n", decrypted)

	// Verification
	if decrypted == message {
		fmt.Println("\n✅ SUCCESS: Secure communication established!")
	} else {
		fmt.Println("\n❌ ERROR: Decryption failed!")
	}

	fmt.Println("\n4. SECURITY PROPERTIES")
	fmt.Println("   ✅ Post-Quantum Security: X-Wing KEM provides PQ security")
	fmt.Println("   ✅ Forward Secrecy: Each session uses new ephemeral shared secret")
	fmt.Println("   ✅ Authentication: AES-GCM provides integrity and authenticity")
	fmt.Println("   ✅ Confidentiality: AES-256 provides strong encryption")
}

/*
% go run main.go
=== Alice and Bob: X-Wing KEM + AES-256-GCM ===

1. KEY EXCHANGE
   Alice generated key pair
   Public key size: 1216 bytes

2. ENCRYPTION
   Original message: "Hello Alice! This is a secret message."
   Bob encrypted message using Alice's public key
   KEM ciphertext: 1120 bytes
   AES-GCM encrypted message: 54 bytes

3. DECRYPTION
   Alice decrypted message: "Hello Alice! This is a secret message."

✅ SUCCESS: Secure communication established!

4. SECURITY PROPERTIES
   ✅ Post-Quantum Security: X-Wing KEM provides PQ security
   ✅ Forward Secrecy: Each session uses new ephemeral shared secret
   ✅ Authentication: AES-GCM provides integrity and authenticity
   ✅ Confidentiality: AES-256 provides strong encryption

 # Real-World Considerations

 // For production use, consider:
func productionKeyDerivation(sharedSecret []byte) []byte {
	// Use proper KDF like HKDF
	// salt := make([]byte, 16)
	// rand.Read(salt)
	// return hkdf.Extract(sha256.New, sharedSecret, salt)
	return sharedSecret[:32] // Simplified for example
}

// Additional security measures:
// - Add proper error handling
// - Use constant-time comparisons
// - Implement key rotation
// - Add additional authentication data (AAD) in GCM
// - Use separate keys for different directions
*/
