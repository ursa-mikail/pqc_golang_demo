# PQCLab — Post-Quantum Cryptography Tutorial

Interactive tutorial covering **Confidentiality · Integrity · Non-Repudiation** using NIST-standardised post-quantum algorithms, with a full migration pipeline via the Trust Agent Mesh.

```
┌─────────────────────────────────────────────────────────────────┐
│          PQC Migration Framework  –  FIPS 202 / 203 / 204 / 205 │
│  Legacy → Joints → Trust-Agent Mesh → PQC 2nd Paths → Pure PQC  │
└─────────────────────────────────────────────────────────────────┘
```

## Algorithm Coverage

| Standard  | Algorithm  | Purpose          | Library                        |
|-----------|------------|------------------|--------------------------------|
| FIPS 202  | SHA-3/SHAKE | Hash functions  | `golang.org/x/crypto/sha3`     |
| FIPS 203  | ML-KEM     | Key encapsulation| `cloudflare/circl` (Kyber)     |
| FIPS 204  | ML-DSA     | Digital signatures| `cloudflare/circl` (Dilithium) |
| FIPS 205  | SLH-DSA    | Hash-based sigs  | SHA-2/SHAKE backed (circl)     |

---

## Quick Start — Docker (recommended)

### Prerequisites
- Docker ≥ 24
- Docker Compose ≥ 2.20

```bash
# Clone / extract the project
cd pqc-tutorial

# Build and start (first build ~2–3 min)
docker compose up --build

# Open in browser
open http://localhost:3000

# Stop
docker compose down

# Full cleanup (removes images too)
docker compose down --rmi all --volumes --remove-orphans
```

The backend is also exposed directly at `http://localhost:8080/api/`.

---

## Local Development (no Docker)

### Backend — Go

```bash
cd backend

# Fetch dependencies + vendor (requires internet on first run)
go mod tidy          # fetches cloudflare/circl and golang.org/x/crypto
go mod vendor        # writes vendor/ directory

# Run (live reload)
go run .

# OR using the vendor directory (fully offline)
go run -mod=vendor .

# Build binary
go build -mod=vendor -o pqc-api .
./pqc-api            # listens on :8080
```

#### Environment
| Variable | Default | Description |
|----------|---------|-------------|
| `PORT`   | `8080`  | HTTP listen port |

### Frontend — TypeScript + React + Vite

```bash
cd frontend

npm install
npm run dev          # http://localhost:5173 — proxies /api → localhost:8080
npm run build        # production build → dist/
```

> The Vite dev server proxies `/api/*` to `http://localhost:8080` automatically.

---

## API Reference

All endpoints return JSON. CORS is open (`*`).

| Method | Path              | Body                                 | Description                    |
|--------|-------------------|--------------------------------------|--------------------------------|
| GET    | `/api/health`     | —                                    | Health check                   |
| GET    | `/api/algorithms` | —                                    | List all supported algorithms  |
| POST   | `/api/kem/demo`   | `{variant:"512\|768\|1024", message}`| ML-KEM encap/decap round-trip  |
| POST   | `/api/sign/demo`  | `{variant:"44\|65\|87", message}`    | ML-DSA sign/verify + tamper    |
| POST   | `/api/slhdsa/demo`| `{variant:"slh-dsa-sha2-128s…", message}` | SLH-DSA sign/verify + tamper |
| POST   | `/api/hash/demo`  | `{variant:"sha3-256…", data}`        | SHA-3/SHAKE hash + avalanche   |
| POST   | `/api/mesh/attest`| —                                    | Full mesh attestation + prune  |

### Example: ML-KEM-768

```bash
curl -X POST http://localhost:8080/api/kem/demo \
  -H 'Content-Type: application/json' \
  -d '{"variant":"768","message":"establish secure channel"}'
```

Response:
```json
{
  "algorithm": "ML-KEM-768",
  "standard": "FIPS 203",
  "securityCategory": 3,
  "verifiedMatch": true,
  "ciphertextBytes": 1088,
  "sharedSecretLen": 32,
  "steps": [...]
}
```

### Example: ML-DSA-65 Sign + Verify

```bash
curl -X POST http://localhost:8080/api/sign/demo \
  -H 'Content-Type: application/json' \
  -d '{"variant":"65","message":"integrity-protected message"}'
```

### Example: Mesh Attestation

```bash
curl -X POST http://localhost:8080/api/mesh/attest | jq .
```

---

## Vendor Directory

The `backend/vendor/` directory is pre-populated with:

- `github.com/cloudflare/circl@v1.3.7` — Kyber (ML-KEM) + Dilithium (ML-DSA)
- `golang.org/x/crypto/sha3` — SHA-3 / SHAKE (sourced from circl's internal sha3)
- `golang.org/x/sys/cpu` — CPU feature detection stub (AVX2 disabled → falls back to pure-Go)

This means the backend builds **fully offline** with `go run -mod=vendor .`.

---

## Security Properties Demonstrated

### Confidentiality (ML-KEM)
- Establishes shared secret without transmitting it
- Security: Module Learning With Errors (MLWE) — hard for quantum computers
- Parameter sets: ML-KEM-512/768/1024 (NIST Categories 1/3/5)

### Integrity (ML-DSA + SHA-3)
- Any bit flip in the signed message → verification failure
- Avalanche effect: 1 input bit change → ~50% output bits change in SHA-3
- ML-DSA security: Module Short Integer Solution (MSIS) problem

### Non-Repudiation (SLH-DSA)
- Stateless hash-based signatures — no per-key counter
- Security relies only on hash function collision resistance
- 12 parameter sets (SHA2/SHAKE × 4 security levels × small/fast)
- Each signing uses fresh internal randomness `R`

### Migration (Trust Agent Mesh)
- Dual-path attestation: legacy + PQC run simultaneously
- Prune gate: removes legacy only after `AllPassed == true`
- Zero-downtime migration pattern

---

## Architecture

```
pqc-tutorial/
├── backend/
│   ├── main.go              # HTTP API server (replaces original CLI main)
│   ├── pqc_algorithms.go    # ML-KEM, ML-DSA, SLH-DSA, SHA-3 providers
│   ├── joints.go            # Cryptographic joints (legacy ↔ PQC bridges)
│   ├── legacy.go            # Legacy crypto layer (RSA/ECDH/ECDSA/AES)
│   ├── trust_agent_mesh.go  # TA-KEM / TA-Sign / TA-Hash + attestation
│   ├── prune_gate.go        # Prune gate + PurePQCState
│   ├── go.mod / go.sum
│   ├── vendor/              # All deps vendored (offline builds)
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── App.tsx
│   │   ├── components/
│   │   │   ├── OverviewPanel.tsx        # Landing + concept cards
│   │   │   ├── ConfidentialityPanel.tsx # ML-KEM demo
│   │   │   ├── IntegrityPanel.tsx       # ML-DSA + SHA-3 demo
│   │   │   ├── NonRepudiationPanel.tsx  # SLH-DSA demo
│   │   │   └── MeshPanel.tsx            # Trust mesh + attestation
│   │   ├── hooks/useApi.ts
│   │   └── types/index.ts
│   ├── nginx.conf
│   └── Dockerfile
└── docker-compose.yml
```

---

## References

- [FIPS 202 — SHA-3 Standard](https://csrc.nist.gov/publications/detail/fips/202/final)
- [FIPS 203 — ML-KEM](https://csrc.nist.gov/publications/detail/fips/203/final)
- [FIPS 204 — ML-DSA](https://csrc.nist.gov/publications/detail/fips/204/final)
- [FIPS 205 — SLH-DSA](https://csrc.nist.gov/publications/detail/fips/205/final)
- [cloudflare/circl](https://github.com/cloudflare/circl)
- [golang.org/x/crypto](https://pkg.go.dev/golang.org/x/crypto)
