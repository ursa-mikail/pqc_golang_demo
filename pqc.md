| Family | Typical Use | Strengths | Weaknesses | Industry Adoption Direction | NIST Standard Algorithm(s) |
|--------|-------------|-----------|------------|----------------------------|----------------------------|
| Lattice-Based | Encryption, Signatures, KEM. TLS, VPN, PKI | Efficient, versatile, NIST winners (Kyber, Dilithium) | Larger keys vs RSA/ECC | Primary standard (broad deployment) | Kyber (KEM), Dilithium (signatures), Falcon (signatures) |
| Code-Based | Encryption, KEM. Long-term encryption, high-security infra | Decades of trust, very secure | Very large public keys | Safe fallback for niche, high-assurance | Classic McEliece (KEM) |
| Multivariate | Signatures. N/A (research only) | Fast signatures, algebraic hardness | Many schemes broken, large keys | Research niche. Not adopted currently | None standardized (Rainbow broken) |
| Hash-Based | Signatures. Compliance, archives | Extremely robust, minimal assumptions | Large signatures, slower | Conservative backup (Backup for conservative use cases) | SPHINCS+ (signatures) |

