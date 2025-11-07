package main

import (
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/kem/xwing"
)

func main() {
	fmt.Println("=== X-Wing KEM Breakdown ===\n")

	// Generate seeds
	keySeed := make([]byte, 32)
	rand.Read(keySeed)

	encapSeed := make([]byte, 64)
	rand.Read(encapSeed)

	fmt.Printf("1. KEY GENERATION\n")
	fmt.Printf("   Key derivation seed: %x (%d bytes)\n", keySeed, len(keySeed))

	// Derive key pair
	sk, pk := xwing.DeriveKeyPairPacked(keySeed)

	fmt.Printf("\n2. KEY PAIR STRUCTURE\n")
	fmt.Printf("   Public Key (pk):\n")
	fmt.Printf("     - Full size: %d bytes\n", len(pk))
	fmt.Printf("     - First 50 bytes: %.100x\n", pk[:min(50, len(pk))])
	fmt.Printf("     - Composition:\n")
	fmt.Printf("       • X25519 public key: 32 bytes\n")
	fmt.Printf("       • ML-KEM-768 public key: 1,184 bytes\n")
	fmt.Printf("       • Total: 1,216 bytes\n")

	fmt.Printf("\n   Private Key (sk):\n")
	fmt.Printf("     - Packed/seed form: %x (%d bytes)\n", sk, len(sk))
	fmt.Printf("     - This is just the original 32-byte seed (z)\n")
	fmt.Printf("     - In memory, expanded to ~2,464 bytes:\n")
	fmt.Printf("       • ML-KEM-768 private key: 2,400 bytes\n")
	fmt.Printf("       • X25519 private key: 32 bytes\n")
	fmt.Printf("       • Precomputed X25519 public key: 32 bytes\n")

	fmt.Printf("\n3. ENCAPSULATION\n")
	fmt.Printf("   Encapsulation seed: %x (%d bytes)\n", encapSeed, len(encapSeed))

	// Encapsulate
	ss, ct, _ := xwing.Encapsulate(pk, encapSeed)

	fmt.Printf("\n   Ciphertext (ct):\n")
	fmt.Printf("     - Full size: %d bytes\n", len(ct))
	fmt.Printf("     - First 50 bytes: %.100x\n", ct[:min(50, len(ct))])
	fmt.Printf("     - Composition:\n")
	fmt.Printf("       • ML-KEM-768 ciphertext: 1,088 bytes\n")
	fmt.Printf("       • X25519 encapsulated key: 32 bytes\n")
	fmt.Printf("       • Total: 1,120 bytes\n")

	fmt.Printf("\n4. SHARED SECRET\n")
	fmt.Printf("   Generated shared secret: %x (%d bytes)\n", ss, len(ss))

	// Decapsulate
	ss2 := xwing.Decapsulate(ct, sk)

	fmt.Printf("   Decapsulated shared secret: %x (%d bytes)\n", ss2, len(ss2))

	// Verify they match
	if string(ss) == string(ss2) {
		fmt.Printf("\n✅ SUCCESS: Shared secrets match!\n")
	} else {
		fmt.Printf("\n❌ ERROR: Shared secrets don't match!\n")
	}

	fmt.Printf("\n5. SIZE SUMMARY\n")
	fmt.Printf("   +----------------------+------------+---------------------+\n")
	fmt.Printf("   | Component            | Size       | Notes               |\n")
	fmt.Printf("   +----------------------+------------+---------------------+\n")
	fmt.Printf("   | Public Key (pk)      | %5d bytes | X25519 + ML-KEM-768 |\n", len(pk))
	fmt.Printf("   | Ciphertext (ct)      | %5d bytes | ML-KEM + X25519     |\n", len(ct))
	fmt.Printf("   | Private Key (packed) | %5d bytes | Seed only           |\n", len(sk))
	fmt.Printf("   | Private Key (expanded)| ~2,464 bytes | In memory         |\n")
	fmt.Printf("   | Shared Secret        | %5d bytes |                     |\n", len(ss))
	fmt.Printf("   +----------------------+------------+---------------------+\n")
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
% go run main.go
=== X-Wing KEM Breakdown ===

1. KEY GENERATION
   Key derivation seed: bd0b90b6983cd7aebed31b6938dbbd66415fc5cb554672a8698e222247f3afd0 (32 bytes)

2. KEY PAIR STRUCTURE
   Public Key (pk):
     - Full size: 1216 bytes
     - First 50 bytes: f00a82b6cc46558972fab81bb9b47f266b48402a558a4121e3f8b0567709db304a0d304d3af82efed0adf5d63e392a716526
     - Composition:
       • X25519 public key: 32 bytes
       • ML-KEM-768 public key: 1,184 bytes
       • Total: 1,216 bytes

   Private Key (sk):
     - Packed/seed form: bd0b90b6983cd7aebed31b6938dbbd66415fc5cb554672a8698e222247f3afd0 (32 bytes)
     - This is just the original 32-byte seed (z)
     - In memory, expanded to ~2,464 bytes:
       • ML-KEM-768 private key: 2,400 bytes
       • X25519 private key: 32 bytes
       • Precomputed X25519 public key: 32 bytes

3. ENCAPSULATION
   Encapsulation seed: c5642236d14a4c1658630e1df8fd18bbd6b24c84518b0b23f1be70a8d937c09f871b0c91011d9147eadbebe43307c5a23d3dee7b097371cba61666aeb2903b34 (64 bytes)

   Ciphertext (ct):
     - Full size: 1120 bytes
     - First 50 bytes: be5b0cb5f55de1d873a472ec9b981a80dd6925628080740ac13dcbcea2a80641479597456f0909050a13f622d14025ef993f
     - Composition:
       • ML-KEM-768 ciphertext: 1,088 bytes
       • X25519 encapsulated key: 32 bytes
       • Total: 1,120 bytes

4. SHARED SECRET
   Generated shared secret: 65d72e9e13c59a5085feb1dd10630fb753a34624a2b524be82ccde4626de4688 (32 bytes)
   Decapsulated shared secret: 65d72e9e13c59a5085feb1dd10630fb753a34624a2b524be82ccde4626de4688 (32 bytes)

✅ SUCCESS: Shared secrets match!

5. SIZE SUMMARY
   +----------------------+------------+---------------------+
   | Component            | Size       | Notes               |
   +----------------------+------------+---------------------+
   | Public Key (pk)      |  1216 bytes | X25519 + ML-KEM-768 |
   | Ciphertext (ct)      |  1120 bytes | ML-KEM + X25519     |
   | Private Key (packed) |    32 bytes | Seed only           |
   | Private Key (expanded)| ~2,464 bytes | In memory         |
   | Shared Secret        |    32 bytes |                     |
   +----------------------+------------+---------------------+
*/
