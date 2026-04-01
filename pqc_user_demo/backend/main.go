package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// ─────────────────────────────────────────────────────────────────────────────
// HTTP API — PQC Tutorial Backend
// ─────────────────────────────────────────────────────────────────────────────

func main() {
	mux := http.NewServeMux()

	// CORS middleware wrapper
	handler := corsMiddleware(mux)

	mux.HandleFunc("/api/health", handleHealth)
	mux.HandleFunc("/api/kem/demo", handleKEMDemo)
	mux.HandleFunc("/api/sign/demo", handleSignDemo)
	mux.HandleFunc("/api/hash/demo", handleHashDemo)
	mux.HandleFunc("/api/slhdsa/demo", handleSLHDSADemo)
	mux.HandleFunc("/api/mesh/attest", handleMeshAttest)
	mux.HandleFunc("/api/algorithms", handleAlgorithms)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("PQC Tutorial API listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}

// ─── CORS ────────────────────────────────────────────────────────────────────

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// ─── /api/health ─────────────────────────────────────────────────────────────

func handleHealth(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"status": "ok", "timestamp": time.Now().UTC().Format(time.RFC3339)})
}

// ─── /api/algorithms ─────────────────────────────────────────────────────────

type AlgorithmInfo struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Standard     string `json:"standard"`
	Category     int    `json:"category"`
	ParameterSet string `json:"parameterSet"`
	Purpose      string `json:"purpose"`
	Description  string `json:"description"`
}

func handleAlgorithms(w http.ResponseWriter, r *http.Request) {
	algos := []AlgorithmInfo{
		// KEM
		{ID: "ml-kem-512", Name: "ML-KEM-512", Standard: "FIPS 203", Category: 1, ParameterSet: "k=2", Purpose: "confidentiality", Description: "Module Lattice KEM — fastest, equivalent to AES-128 security"},
		{ID: "ml-kem-768", Name: "ML-KEM-768", Standard: "FIPS 203", Category: 3, ParameterSet: "k=3", Purpose: "confidentiality", Description: "Module Lattice KEM — recommended, equivalent to AES-192 security"},
		{ID: "ml-kem-1024", Name: "ML-KEM-1024", Standard: "FIPS 203", Category: 5, ParameterSet: "k=4", Purpose: "confidentiality", Description: "Module Lattice KEM — maximum security, equivalent to AES-256"},
		// DSA
		{ID: "ml-dsa-44", Name: "ML-DSA-44", Standard: "FIPS 204", Category: 2, ParameterSet: "k=4,l=4", Purpose: "integrity", Description: "Module Lattice DSA — smaller signatures, faster verification"},
		{ID: "ml-dsa-65", Name: "ML-DSA-65", Standard: "FIPS 204", Category: 3, ParameterSet: "k=6,l=5", Purpose: "integrity", Description: "Module Lattice DSA — balanced security/performance (NIST Cat 3)"},
		{ID: "ml-dsa-87", Name: "ML-DSA-87", Standard: "FIPS 204", Category: 5, ParameterSet: "k=8,l=7", Purpose: "integrity", Description: "Module Lattice DSA — highest security category"},
		// SLH-DSA
		{ID: "slh-dsa-sha2-128s", Name: "SLH-DSA-SHA2-128s", Standard: "FIPS 205", Category: 1, ParameterSet: "n=16,h=63,d=7", Purpose: "non-repudiation", Description: "Stateless hash-based — small sig, slow sign, fast verify (SHA-2)"},
		{ID: "slh-dsa-sha2-128f", Name: "SLH-DSA-SHA2-128f", Standard: "FIPS 205", Category: 1, ParameterSet: "n=16,h=66,d=22", Purpose: "non-repudiation", Description: "Stateless hash-based — fast signing, larger signature (SHA-2)"},
		{ID: "slh-dsa-shake-128s", Name: "SLH-DSA-SHAKE-128s", Standard: "FIPS 205", Category: 1, ParameterSet: "n=16,h=63,d=7", Purpose: "non-repudiation", Description: "Stateless hash-based — small sig, slow sign (SHAKE-256 PRF)"},
		{ID: "slh-dsa-shake-128f", Name: "SLH-DSA-SHAKE-128f", Standard: "FIPS 205", Category: 1, ParameterSet: "n=16,h=66,d=22", Purpose: "non-repudiation", Description: "Stateless hash-based — fast signing (SHAKE-256 PRF)"},
		{ID: "slh-dsa-sha2-256s", Name: "SLH-DSA-SHA2-256s", Standard: "FIPS 205", Category: 5, ParameterSet: "n=32,h=64,d=8", Purpose: "non-repudiation", Description: "Stateless hash-based — maximum security, conservative"},
		// Hash
		{ID: "sha3-256", Name: "SHA3-256", Standard: "FIPS 202", Category: 1, ParameterSet: "256-bit", Purpose: "integrity", Description: "Sponge-based hash — quantum-resistant output collision resistance"},
		{ID: "sha3-512", Name: "SHA3-512", Standard: "FIPS 202", Category: 5, ParameterSet: "512-bit", Purpose: "integrity", Description: "Sponge-based hash — 256-bit quantum security level"},
		{ID: "shake256", Name: "SHAKE256", Standard: "FIPS 202", Category: 5, ParameterSet: "512-bit", Purpose: "integrity", Description: "Extendable Output Function — variable-length, general purpose"},
	}
	writeJSON(w, algos)
}

// ─── /api/kem/demo ───────────────────────────────────────────────────────────

type KEMRequest struct {
	Variant string `json:"variant"` // "512" | "768" | "1024"
	Message string `json:"message"`
}

type KEMResponse struct {
	Algorithm       string        `json:"algorithm"`
	Standard        string        `json:"standard"`
	SecurityCat     int           `json:"securityCategory"`
	PublicKeyHex    string        `json:"publicKeyHex"`
	CiphertextHex   string        `json:"ciphertextHex"`
	SharedSecretHex string        `json:"sharedSecretHex"`
	VerifiedMatch   bool          `json:"verifiedMatch"`
	CiphertextBytes int           `json:"ciphertextBytes"`
	SharedSecretLen int           `json:"sharedSecretLen"`
	Duration        time.Duration `json:"durationNs"`
	Steps           []string      `json:"steps"`
}

func handleKEMDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var req KEMRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var variant MLKEMVariant
	switch req.Variant {
	case "512":
		variant = MLKEM512
	case "768":
		variant = MLKEM768
	case "1024":
		variant = MLKEM1024
	default:
		variant = MLKEM768
	}

	start := time.Now()
	provider := NewMLKEMProvider(variant)

	ct, ss, err := provider.Encapsulate([]byte(req.Message))
	if err != nil {
		writeError(w, fmt.Sprintf("encapsulate: %v", err), http.StatusInternalServerError)
		return
	}

	ss2, err := provider.Decapsulate(ct)
	if err != nil {
		writeError(w, fmt.Sprintf("decapsulate: %v", err), http.StatusInternalServerError)
		return
	}

	match := string(ss) == string(ss2)
	dur := time.Since(start)

	// Get public key bytes for display
	pkBytes := provider.ECDHPublicBytesLen()

	writeJSON(w, KEMResponse{
		Algorithm:       provider.Name(),
		Standard:        provider.FIPSStandard(),
		SecurityCat:     provider.SecurityCategory(),
		PublicKeyHex:    hex.EncodeToString(ss[:min(8, len(ss))]) + "… (" + fmt.Sprintf("%d", pkBytes) + "B)",
		CiphertextHex:   hex.EncodeToString(ct[:min(16, len(ct))]) + "…",
		SharedSecretHex: hex.EncodeToString(ss),
		VerifiedMatch:   match,
		CiphertextBytes: len(ct),
		SharedSecretLen: len(ss),
		Duration:        dur,
		Steps: []string{
			fmt.Sprintf("1. Generated ML-KEM-%s key pair using crypto/rand", req.Variant),
			fmt.Sprintf("2. Alice encapsulates: produces %d-byte ciphertext + shared secret", len(ct)),
			fmt.Sprintf("3. Bob decapsulates: recovers same %d-byte shared secret from ciphertext", len(ss2)),
			fmt.Sprintf("4. Shared secrets match: %v → confidential channel established", match),
			fmt.Sprintf("5. Quantum safe: lattice problems remain hard even for Shor's algorithm"),
		},
	})
}

// ECDHPublicBytesLen returns the expected public key byte length for the variant
func (m *MLKEMProvider) ECDHPublicBytesLen() int {
	return mlkemParams[m.variant].pkBytes
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ─── /api/sign/demo ──────────────────────────────────────────────────────────

type SignRequest struct {
	Variant string `json:"variant"` // "44" | "65" | "87"
	Message string `json:"message"`
}

type SignResponse struct {
	Algorithm       string        `json:"algorithm"`
	Standard        string        `json:"standard"`
	SecurityCat     int           `json:"securityCategory"`
	MessageHex      string        `json:"messageHex"`
	SignatureHex    string        `json:"signatureHex"`
	SignatureBytes  int           `json:"signatureBytes"`
	Verified        bool          `json:"verified"`
	TamperedVerify  bool          `json:"tamperedVerify"`
	Duration        time.Duration `json:"durationNs"`
	Steps           []string      `json:"steps"`
}

func handleSignDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var req SignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	var variant MLDSAVariant
	switch req.Variant {
	case "44":
		variant = MLDSA44
	case "65":
		variant = MLDSA65
	case "87":
		variant = MLDSA87
	default:
		variant = MLDSA65
	}

	msg := []byte(req.Message)
	if len(msg) == 0 {
		msg = []byte("Hello, post-quantum world!")
	}

	start := time.Now()
	provider := NewMLDSAProvider(variant)

	sig, err := provider.Sign(msg)
	if err != nil {
		writeError(w, fmt.Sprintf("sign: %v", err), http.StatusInternalServerError)
		return
	}

	verified, _ := provider.Verify(msg, sig)

	// Tamper test
	tampered := make([]byte, len(msg))
	copy(tampered, msg)
	if len(tampered) > 0 {
		tampered[0] ^= 0x01
	}
	tamperedVerify, _ := provider.Verify(tampered, sig)
	dur := time.Since(start)

	writeJSON(w, SignResponse{
		Algorithm:      provider.Name(),
		Standard:       provider.FIPSStandard(),
		SecurityCat:    provider.SecurityCategory(),
		MessageHex:     hex.EncodeToString(msg),
		SignatureHex:   hex.EncodeToString(sig[:min(24, len(sig))]) + "…",
		SignatureBytes: len(sig),
		Verified:       verified,
		TamperedVerify: tamperedVerify,
		Duration:       dur,
		Steps: []string{
			fmt.Sprintf("1. Generated ML-DSA-%s key pair (Dilithium lattice)", req.Variant),
			fmt.Sprintf("2. Signed %d-byte message → %d-byte signature", len(msg), len(sig)),
			fmt.Sprintf("3. Verified original message: %v", verified),
			fmt.Sprintf("4. Verified tampered message (1 bit flip): %v ← must be false", tamperedVerify),
			"5. Integrity protected: any message alteration invalidates signature",
		},
	})
}

// ─── /api/slhdsa/demo ────────────────────────────────────────────────────────

type SLHRequest struct {
	Variant string `json:"variant"`
	Message string `json:"message"`
}

type SLHResponse struct {
	Algorithm      string        `json:"algorithm"`
	Standard       string        `json:"standard"`
	SecurityCat    int           `json:"securityCategory"`
	HashFamily     string        `json:"hashFamily"`
	SignatureBytes int           `json:"signatureBytes"`
	Verified       bool          `json:"verified"`
	TamperedVerify bool          `json:"tamperedVerify"`
	Duration       time.Duration `json:"durationNs"`
	Steps          []string      `json:"steps"`
}

func handleSLHDSADemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var req SLHRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	variantMap := map[string]SLHDSAVariant{
		"slh-dsa-sha2-128s":  SLHDSA_SHA2_128s,
		"slh-dsa-sha2-128f":  SLHDSA_SHA2_128f,
		"slh-dsa-sha2-256s":  SLHDSA_SHA2_256s,
		"slh-dsa-shake-128s": SLHDSA_SHAKE_128s,
		"slh-dsa-shake-128f": SLHDSA_SHAKE_128f,
	}

	variant, ok := variantMap[req.Variant]
	if !ok {
		variant = SLHDSA_SHA2_128s
	}

	msg := []byte(req.Message)
	if len(msg) == 0 {
		msg = []byte("Non-repudiation demo message")
	}

	start := time.Now()
	provider := NewSLHDSAProvider(variant)

	sig, err := provider.Sign(msg)
	if err != nil {
		writeError(w, fmt.Sprintf("sign: %v", err), http.StatusInternalServerError)
		return
	}

	verified, _ := provider.Verify(msg, sig)

	tampered := make([]byte, len(msg))
	copy(tampered, msg)
	if len(tampered) > 0 {
		tampered[0] ^= 0xFF
	}
	tamperedVerify, _ := provider.Verify(tampered, sig)
	dur := time.Since(start)

	p := slhdsaParamSets[variant]
	hashFamily := p.hashFamily

	writeJSON(w, SLHResponse{
		Algorithm:      string(variant),
		Standard:       "FIPS 205",
		SecurityCat:    provider.SecurityCategory(),
		HashFamily:     hashFamily,
		SignatureBytes: len(sig),
		Verified:       verified,
		TamperedVerify: tamperedVerify,
		Duration:       dur,
		Steps: []string{
			fmt.Sprintf("1. Generated SLH-DSA key pair (stateless hash tree, %s PRF)", hashFamily),
			fmt.Sprintf("2. Signed %d-byte message → %d-byte signature", len(msg), len(sig)),
			"3. No signing state stored — signing the same message produces a different sig each time",
			fmt.Sprintf("4. Verified original: %v | Tampered: %v", verified, tamperedVerify),
			"5. Non-repudiation: provably bound to key pair, cannot deny signing",
		},
	})
}

// ─── /api/hash/demo ──────────────────────────────────────────────────────────

type HashRequest struct {
	Variant string `json:"variant"` // "sha3-256" | "sha3-512" | "shake256"
	Data    string `json:"data"`
}

type HashResponse struct {
	Algorithm  string        `json:"algorithm"`
	Standard   string        `json:"standard"`
	SecurityCat int          `json:"securityCategory"`
	DigestHex  string        `json:"digestHex"`
	DigestLen  int           `json:"digestLen"`
	AvalancheDemo []string   `json:"avalancheDemo"`
	Duration   time.Duration `json:"durationNs"`
	Steps      []string      `json:"steps"`
}

func handleHashDemo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "POST required", http.StatusMethodNotAllowed)
		return
	}

	var req HashRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	variantMap := map[string]HashVariant{
		"sha3-256": SHA3_256,
		"sha3-384": SHA3_384,
		"sha3-512": SHA3_512,
		"shake128": SHAKE128,
		"shake256": SHAKE256,
	}
	variant, ok := variantMap[req.Variant]
	if !ok {
		variant = SHA3_256
	}

	data := []byte(req.Data)
	if len(data) == 0 {
		data = []byte("post-quantum integrity check")
	}

	start := time.Now()
	provider := NewPQCHashProvider(variant)

	digest, err := provider.Hash(data)
	if err != nil {
		writeError(w, fmt.Sprintf("hash: %v", err), http.StatusInternalServerError)
		return
	}

	// Avalanche demo: one bit flip → completely different hash
	flipped := make([]byte, len(data))
	copy(flipped, data)
	if len(flipped) > 0 {
		flipped[0] ^= 0x01
	}
	digest2, _ := provider.Hash(flipped)
	dur := time.Since(start)

	writeJSON(w, HashResponse{
		Algorithm:   provider.Name(),
		Standard:    provider.FIPSStandard(),
		SecurityCat: provider.SecurityCategory(),
		DigestHex:   hex.EncodeToString(digest),
		DigestLen:   len(digest),
		AvalancheDemo: []string{
			"Original: " + hex.EncodeToString(digest[:8]) + "…",
			"1-bit flip: " + hex.EncodeToString(digest2[:8]) + "…",
		},
		Duration: dur,
		Steps: []string{
			fmt.Sprintf("1. Input: %q (%d bytes)", string(data), len(data)),
			fmt.Sprintf("2. SHA-3 sponge absorbs all input blocks"),
			fmt.Sprintf("3. Squeeze phase produces %d-byte digest", len(digest)),
			"4. Avalanche effect: 1-bit change → ~50%% of output bits change",
			"5. Quantum safe: Grover's gives √N speedup — halves bit-security, still safe at 256+",
		},
	})
}

// ─── /api/mesh/attest ────────────────────────────────────────────────────────

type MeshAttestResponse struct {
	Timestamp  string            `json:"timestamp"`
	AllPassed  bool              `json:"allPassed"`
	Results    []MeshPathResult  `json:"results"`
	Pruned     bool              `json:"pruned"`
	PrunedAt   string            `json:"prunedAt,omitempty"`
}

type MeshPathResult struct {
	Provider   string `json:"provider"`
	Standard   string `json:"standard"`
	Category   int    `json:"category"`
	LegacyOK   bool   `json:"legacyOk"`
	PQCPathOK  bool   `json:"pqcPathOk"`
	Compatible bool   `json:"compatible"`
	Notes      string `json:"notes"`
}

func handleMeshAttest(w http.ResponseWriter, r *http.Request) {
	legacy := NewLegacyCryptoLayer(LegacyConfig{
		RSAKeyBits: 2048, ECDHCurve: "P-256", ECDSACurve: "P-384",
		AESKeyBits: 256, AESLegacyGCM: true,
	})

	kemJoint := NewKEMJoint(legacy)
	signJoint := NewSignJoint(legacy)
	hashJoint := NewHashJoint(legacy)

	mesh := NewTrustAgentMesh(TrustAgentMeshConfig{
		KEMJoint: kemJoint, SignJoint: signJoint, HashJoint: hashJoint,
		Attesters: []AttestationMode{DualPathAttestation},
	})

	mesh.RegisterPQCPath(NewMLKEMProvider(MLKEM768))
	mesh.RegisterPQCPath(NewMLDSAProvider(MLDSA65))
	mesh.RegisterPQCPath(NewSLHDSAProvider(SLHDSA_SHA2_128s))

	report, err := mesh.AttestAll()
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var results []MeshPathResult
	for _, res := range report.Results {
		results = append(results, MeshPathResult{
			Provider:   res.ProviderName,
			Standard:   res.Standard,
			Category:   res.Category,
			LegacyOK:   res.LegacyOK,
			PQCPathOK:  res.PQCPathOK,
			Compatible: res.Compatible,
			Notes:      res.Notes,
		})
	}

	gate := NewPruneGate(mesh, report)
	pruned, _ := gate.Evaluate()

	resp := MeshAttestResponse{
		Timestamp: report.Timestamp.Format(time.RFC3339),
		AllPassed: report.AllPassed,
		Results:   results,
		Pruned:    pruned != nil,
	}
	if pruned != nil {
		resp.PrunedAt = pruned.At.Format(time.RFC3339)
	}

	writeJSON(w, resp)
}
