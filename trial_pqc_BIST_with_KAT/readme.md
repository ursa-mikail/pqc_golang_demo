
```
% go run .
=== Post-Quantum Cryptography Demo with BIST ===
=== Kyber768 Key Encapsulation Example ===
Public key size: 1184 bytes
Private key size: 2400 bytes
Ciphertext size: 1088 bytes
Shared secret size: 32 bytes
Shared secret: 3979cfc014a704f91686710255350db9
✅ Shared secrets match!
=== Cloudflare CIRCL Post-Quantum Cryptography Examples ===
Available algorithms in CIRCL v1.6.1:
✓ ML-DSA (standardized Dilithium)
✓ Dilithium (pre-standardization)
✓ Kyber KEM
✗ FALCON (not available)
✗ SPHINCS+/SLH-DSA (not in v1.6.1, only in main branch)
✗ HQC (not available)

=== Example 1: ML-DSA-44 (FIPS 204) ===
ML-DSA-44 parameters:
  Public key size:  1312 bytes
  Private key size: 2560 bytes
  Signature size:   2420 bytes
Generating ML-DSA key pair...
Signing message: Hello from ML-DSA (FIPS 204 standardized Dilithium)!
Generated signature (2420 bytes)
✓ ML-DSA signature verification PASSED

=== Post-Quantum Cryptography Demo ===
🔒 Security Level: Level 1 (~AES-128)
--------------------------------------------------
📡 Key Encapsulation Mechanism (KEM):
  Algorithm: Kyber512
  Public Key: 800 bytes
  Private Key: 1632 bytes
  Ciphertext: 768 bytes
  Shared Secret: 32 bytes
  Secrets Match: true ✅
🔐 KEM + ChaCha20 Demo:
  ✅ Shared secret established (32 bytes)
  Nonce: 3a608e8a747188c576205b3c
  Plaintext: Hello PQC + ChaCha20 world!
  Ciphertext: 300b4d62b0ab4d9af20901c283f61ea29244fef8b831ae0675eb27
  Recovered: Hello PQC + ChaCha20 world!
✍️  Digital Signatures:
  Algorithm: Dilithium2
  Public Key: 1312 bytes
  Private Key: 2528 bytes
  Signature: 2420 bytes
  Message: "Hello, Post-Quantum World!"
  Valid: true ✅
🏷️  Post-Quantum Hashing:
  Algorithm: SHAKE128
  Input: 44 bytes
  Hash: 16 bytes
  Hash (hex): d6c079ceda4db23e61a8d994409b2c25

🔒 Security Level: Level 3 (~AES-192)
--------------------------------------------------
📡 Key Encapsulation Mechanism (KEM):
  Algorithm: Kyber768
  Public Key: 1184 bytes
  Private Key: 2400 bytes
  Ciphertext: 1088 bytes
  Shared Secret: 32 bytes
  Secrets Match: true ✅
🔐 KEM + ChaCha20 Demo:
  ✅ Shared secret established (32 bytes)
  Nonce: 1a1736ab48eb0ba97f883fd3
  Plaintext: Hello PQC + ChaCha20 world!
  Ciphertext: 2cf51b0ad316ec51b401407931be0e5d578edbd73ed3a5bd05ba15
  Recovered: Hello PQC + ChaCha20 world!
✍️  Digital Signatures:
  Algorithm: Dilithium3
  Public Key: 1952 bytes
  Private Key: 4000 bytes
  Signature: 3293 bytes
  Message: "Hello, Post-Quantum World!"
  Valid: true ✅
🏷️  Post-Quantum Hashing:
  Algorithm: SHAKE256
  Input: 44 bytes
  Hash: 32 bytes
  Hash (hex): 54ee4daf3a1fd61712dbaa3c0d4b270a

🔒 Security Level: Level 5 (~AES-256)
--------------------------------------------------
📡 Key Encapsulation Mechanism (KEM):
  Algorithm: Kyber1024
  Public Key: 1568 bytes
  Private Key: 3168 bytes
  Ciphertext: 1568 bytes
  Shared Secret: 32 bytes
  Secrets Match: true ✅
🔐 KEM + ChaCha20 Demo:
  ✅ Shared secret established (32 bytes)
  Nonce: 484b0132189828bdb70eb549
  Plaintext: Hello PQC + ChaCha20 world!
  Ciphertext: ff813902e3857dfc374a455956a9f285aa468d28c7834e62f20368
  Recovered: Hello PQC + ChaCha20 world!
✍️  Digital Signatures:
  Algorithm: Dilithium5
  Public Key: 2592 bytes
  Private Key: 4864 bytes
  Signature: 4595 bytes
  Message: "Hello, Post-Quantum World!"
  Valid: true ✅
🏷️  Post-Quantum Hashing:
  Algorithm: SHA3-256
  Input: 44 bytes
  Hash: 32 bytes
  Hash (hex): 38b2c17fce6eb59c8a4aed9bc48e5747


Would you like to run the Built-In Self Test (BIST) suite?
This will generate comprehensive test vectors and validate all algorithms.
Run BIST? (y/N): y

================================================================================
STARTING POST-QUANTUM CRYPTOGRAPHY BUILT-IN SELF TEST (BIST)
================================================================================
Starting Comprehensive BIST Suite...
================================================================================
Phase 1: Generating test vectors...
Generated 75 test vectors total

Phase 2: Running KEM BIST...

Phase 3: Running Signature BIST...

Phase 4: Running cross-validation tests...

Phase 5: Running performance tests...

Exit Criteria Evaluation:
  Critical Tests Passed: true
  Overall Pass Rate: 100.0% (required: 95%)
  Test Vectors Generated: 75 (KEM: 30, SIG: 45)
  Vector Requirements Met: true
  EXIT CRITERIA MET: true

================================================================================
POST-QUANTUM CRYPTOGRAPHY BIST COMPREHENSIVE REPORT
================================================================================
Execution Time: 326.821084ms
Total Tests: 27
Passed: 27
Failed: 0
Success Rate: 100.0%
Test Vectors: 75
Exit Criteria Met: true

DETAILED RESULTS:
--------------------------------------------------------------------------------

KEM Algorithms:
  KEM             Test Vector Validation              [PASS] 3.894875ms (30 vectors)
  Kyber512        Algorithm Stress Test               [PASS] 8.973167ms (50 iter)
  Kyber768        Algorithm Stress Test               [PASS] 12.20375ms (50 iter)
  Kyber1024       Algorithm Stress Test               [PASS] 15.480792ms (50 iter)

Signature Algorithms:
  SIGNATURE       Test Vector Validation              [PASS]  9.156ms (45 vectors)
  Dilithium3      Algorithm Stress Test               [PASS] 84.328208ms (25 iter)
  Dilithium5      Algorithm Stress Test               [PASS] 111.423375ms (25 iter)
  Dilithium2      Algorithm Stress Test               [PASS] 47.41875ms (25 iter)

Performance Tests:
  Kyber512        Key Generation Performance          [PASS] 24.667µs
  Kyber512        Encapsulation Performance           [PASS] 30.209µs
  Kyber512        Decapsulation Performance           [PASS]   44.5µs
  Kyber768        Key Generation Performance          [PASS] 51.292µs
  Kyber768        Encapsulation Performance           [PASS]   54.5µs
  Kyber768        Decapsulation Performance           [PASS] 53.458µs
  Kyber1024       Key Generation Performance          [PASS] 67.583µs
  Kyber1024       Encapsulation Performance           [PASS] 74.334µs
  Kyber1024       Decapsulation Performance           [PASS] 77.875µs
  Dilithium2      Key Generation Performance          [PASS] 92.083µs
  Dilithium2      Signing Performance                 [PASS] 205.083µs
  Dilithium2      Verification Performance            [PASS] 75.958µs
  Dilithium3      Key Generation Performance          [PASS] 153.167µs
  Dilithium3      Signing Performance                 [PASS] 182.292µs
  Dilithium3      Verification Performance            [PASS] 122.25µs
  Dilithium5      Key Generation Performance          [PASS] 220.917µs
  Dilithium5      Signing Performance                 [PASS] 284.875µs
  Dilithium5      Verification Performance            [PASS] 202.833µs

Cross-Validation:
  CROSS-VALIDATION Algorithm Interference Test         [PASS] 2.760041ms

TEST VECTOR SUMMARY:
----------------------------------------
KEM Test Vectors: 30
Signature Test Vectors: 45
Valid Cases: 30
Invalid Cases: 45
Total: 75
================================================================================
✅ Saved test vectors to: pqc_test_vectors.json
2025/09/23 14:17:58 Saved BIST test_vectors to: pqc_test_vectors.json
✅ Saved BIST report to: pqc_bist_report.json (1158603 bytes)
2025/09/23 14:17:58 Saved BIST report to: pqc_bist_report.json

Total BIST Execution Time: 338.686458ms
✅ BIST PASSED - All exit criteria met
System is ready for production use

------------------------------------------------------------
TEST VECTOR VALIDATION DEMONSTRATION
------------------------------------------------------------
Validating Test Vector KEM-001 (Kyber512): ✅ PASSED
Validating Test Vector KEM-002 (Kyber512): ✅ PASSED
Validating Test Vector KEM-003 (Kyber512): ✅ PASSED
Validating Test Vector KEM-004 (Kyber512): ✅ PASSED
Validating Test Vector KEM-005 (Kyber512): ✅ PASSED
Validated 5/75 test vectors
------------------------------------------------------------
```

``` % go run .
 Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 5e23e0152a005e369c02e7bad0e3f93c
  Got:      5e23e0152a005e369c02e7bad0e3f93c
✅ PASSED

Running test: KEM-006 (Invalid Kyber512 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber512
  Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-007 (Valid Kyber512 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber512
  Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 4ee748876833244875584eef0da04c7a
  Got:      4ee748876833244875584eef0da04c7a
✅ PASSED

Running test: KEM-008 (Invalid Kyber512 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber512
  Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-009 (Valid Kyber512 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber512
  Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 652f21a8f03a2c15428c16dbfe3a88b1
  Got:      652f21a8f03a2c15428c16dbfe3a88b1
✅ PASSED

Running test: KEM-010 (Invalid Kyber512 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber512
  Public Key Length: 800 bytes
  Private Key Length: 1632 bytes
  Ciphertext Length: 768 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-011 (Valid Kyber768 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 6f21c09208036461ad73acd06ff18644
  Got:      6f21c09208036461ad73acd06ff18644
✅ PASSED

Running test: KEM-012 (Invalid Kyber768 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-013 (Valid Kyber768 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: f4968e4bf98451b5730137c855a47a53
  Got:      f4968e4bf98451b5730137c855a47a53
✅ PASSED

Running test: KEM-014 (Invalid Kyber768 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-015 (Valid Kyber768 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 3d8990153908552b333fadf1b42bc445
  Got:      3d8990153908552b333fadf1b42bc445
✅ PASSED

Running test: KEM-016 (Invalid Kyber768 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-017 (Valid Kyber768 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 4ff15a40d2410ed65475ae2b94ec35a3
  Got:      4ff15a40d2410ed65475ae2b94ec35a3
✅ PASSED

Running test: KEM-018 (Invalid Kyber768 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-019 (Valid Kyber768 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 6189328a838e2d59e051cdda8683f5fe
  Got:      6189328a838e2d59e051cdda8683f5fe
✅ PASSED

Running test: KEM-020 (Invalid Kyber768 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber768
  Public Key Length: 1184 bytes
  Private Key Length: 2400 bytes
  Ciphertext Length: 1088 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-021 (Valid Kyber1024 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: c7d25743225ab5083afc5bca9e7e7b66
  Got:      c7d25743225ab5083afc5bca9e7e7b66
✅ PASSED

Running test: KEM-022 (Invalid Kyber1024 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-023 (Valid Kyber1024 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 8f3e93cc779978a08dc96684d6622c06
  Got:      8f3e93cc779978a08dc96684d6622c06
✅ PASSED

Running test: KEM-024 (Invalid Kyber1024 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-025 (Valid Kyber1024 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 717268cf96ebb187579b5bf03aa8d908
  Got:      717268cf96ebb187579b5bf03aa8d908
✅ PASSED

Running test: KEM-026 (Invalid Kyber1024 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-027 (Valid Kyber1024 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: 5da8a3acee8ba5fce3b18f1ec9f9fdbc
  Got:      5da8a3acee8ba5fce3b18f1ec9f9fdbc
✅ PASSED

Running test: KEM-028 (Invalid Kyber1024 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: KEM-029 (Valid Kyber1024 KEM operation)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 32 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets match: ✅
  Expected: bf1071ed7a661a6f3a4517c15d7a2322
  Got:      bf1071ed7a661a6f3a4517c15d7a2322
✅ PASSED

Running test: KEM-030 (Invalid Kyber1024 KEM operation - corrupted ciphertext)
----------------------------------------
🔐 KEM Test: Kyber1024
  Public Key Length: 1568 bytes
  Private Key Length: 3168 bytes
  Ciphertext Length: 1568 bytes
  Expected Shared Secret Length: 0 bytes
  Recovered Shared Secret Length: 32 bytes
  Shared secrets correctly do not match: ✅
✅ PASSED

Running test: SIG-001 (Valid Dilithium2 signature for message type 0)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-002 (Invalid Dilithium2 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-003 (Invalid Dilithium2 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium2
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-004 (Valid Dilithium2 signature for message type 1)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 1
  Message Length: 1 bytes
  Is Wrong Message Test: false
  Message (hex): 61
  Message (text): a
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-005 (Invalid Dilithium2 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-006 (Invalid Dilithium2 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium2
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-007 (Valid Dilithium2 signature for message type 2)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 2
  Message Length: 26 bytes
  Is Wrong Message Test: false
  Message (hex): 48656c6c6f2c20506f73742d5175616e74756d20576f726c6421
  Message (text): Hello, Post-Quantum World!
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-008 (Invalid Dilithium2 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-009 (Invalid Dilithium2 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium2
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-010 (Valid Dilithium2 signature for message type 3)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 3
  Message Length: 43 bytes
  Is Wrong Message Test: false
  Message (hex): 54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67
  Message (text): The quick brown fox jumps over the lazy dog
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-011 (Invalid Dilithium2 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-012 (Invalid Dilithium2 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium2
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-013 (Valid Dilithium2 signature for message type 4)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 4
  Message Length: 1000 bytes
  Is Wrong Message Test: false
  Message (hex, first 32 bytes): 000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f...
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-014 (Invalid Dilithium2 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium2
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-015 (Invalid Dilithium2 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium2
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1312 bytes
  Signature Length: 2420 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-016 (Valid Dilithium3 signature for message type 0)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-017 (Invalid Dilithium3 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-018 (Invalid Dilithium3 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium3
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-019 (Valid Dilithium3 signature for message type 1)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 1
  Message Length: 1 bytes
  Is Wrong Message Test: false
  Message (hex): 61
  Message (text): a
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-020 (Invalid Dilithium3 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-021 (Invalid Dilithium3 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium3
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-022 (Valid Dilithium3 signature for message type 2)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 2
  Message Length: 26 bytes
  Is Wrong Message Test: false
  Message (hex): 48656c6c6f2c20506f73742d5175616e74756d20576f726c6421
  Message (text): Hello, Post-Quantum World!
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-023 (Invalid Dilithium3 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-024 (Invalid Dilithium3 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium3
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-025 (Valid Dilithium3 signature for message type 3)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 3
  Message Length: 43 bytes
  Is Wrong Message Test: false
  Message (hex): 54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67
  Message (text): The quick brown fox jumps over the lazy dog
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-026 (Invalid Dilithium3 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-027 (Invalid Dilithium3 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium3
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-028 (Valid Dilithium3 signature for message type 4)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 4
  Message Length: 1000 bytes
  Is Wrong Message Test: false
  Message (hex, first 32 bytes): 000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f...
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-029 (Invalid Dilithium3 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium3
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-030 (Invalid Dilithium3 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium3
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 1952 bytes
  Signature Length: 3293 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-031 (Valid Dilithium5 signature for message type 0)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-032 (Invalid Dilithium5 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-033 (Invalid Dilithium5 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium5
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-034 (Valid Dilithium5 signature for message type 1)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 1
  Message Length: 1 bytes
  Is Wrong Message Test: false
  Message (hex): 61
  Message (text): a
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-035 (Invalid Dilithium5 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-036 (Invalid Dilithium5 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium5
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-037 (Valid Dilithium5 signature for message type 2)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 2
  Message Length: 26 bytes
  Is Wrong Message Test: false
  Message (hex): 48656c6c6f2c20506f73742d5175616e74756d20576f726c6421
  Message (text): Hello, Post-Quantum World!
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-038 (Invalid Dilithium5 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-039 (Invalid Dilithium5 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium5
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-040 (Valid Dilithium5 signature for message type 3)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 3
  Message Length: 43 bytes
  Is Wrong Message Test: false
  Message (hex): 54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67
  Message (text): The quick brown fox jumps over the lazy dog
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-041 (Invalid Dilithium5 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-042 (Invalid Dilithium5 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium5
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-043 (Valid Dilithium5 signature for message type 4)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 4
  Message Length: 1000 bytes
  Is Wrong Message Test: false
  Message (hex, first 32 bytes): 000102030405060708090a0b0c0d0e0f101112131415161718191a1b1c1d1e1f...
  Signature Valid: true
  Signature verification succeeded: ✅
✅ PASSED

Running test: SIG-044 (Invalid Dilithium5 signature - corrupted signature)
----------------------------------------
✍️ Signature Test: Dilithium5
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 0 bytes
  Is Wrong Message Test: false
  Message (hex): 
  Message (text): 
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

Running test: SIG-045 (Invalid Dilithium5 signature - wrong message)
----------------------------------------
✍️ Signature Test: Dilithium5
  [DEBUG] Wrong message test detected
  [DEBUG] Original message type: 0
  [DEBUG] Original message length: 0 bytes
  [DEBUG] Modified message length: 5 bytes
  [DEBUG] Modified message (first 50 bytes): 77726f6e67
  Public Key Length: 2592 bytes
  Signature Length: 4595 bytes
  Message Type: 0
  Message Length: 5 bytes
  Is Wrong Message Test: true
  Message (hex): 77726f6e67
  Message (text): wrong
  Signature Valid: false
  Signature correctly invalid: ✅
✅ PASSED

============================================================
KAT RESULTS SUMMARY
============================================================
Overall Results:
  Total Tests: 75
  Passed: 75 ✅
  Failed: 0 ❌
  Success Rate: 100.0%

Results by Category:
  KEM Tests: 30/30 passed (100.0%)
  Signature Tests: 45/45 passed (100.0%)

Results saved to: kat_results_20250923_141822.json

============================================================🎉 ALL KAT TESTS PASSED!
============================================================
```