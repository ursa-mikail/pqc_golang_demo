import { TabId } from '../types'

interface Props { onNavigate: (tab: TabId) => void }

const concepts = [
  {
    id: 'confidentiality' as TabId,
    title: 'Confidentiality',
    subtitle: 'ML-KEM · FIPS 203',
    color: '#00d4ff',
    glow: 'var(--glow-kem)',
    icon: '🔐',
    threat: 'Shor\'s algorithm breaks RSA/ECDH in polynomial time on a quantum computer.',
    solution: 'ML-KEM (Module Lattice KEM) derives shared secrets from lattice problems — hard for both classical and quantum computers.',
    params: ['ML-KEM-512 · Cat 1', 'ML-KEM-768 · Cat 3', 'ML-KEM-1024 · Cat 5'],
    cta: 'Try KEM Demo →',
  },
  {
    id: 'integrity' as TabId,
    title: 'Integrity',
    subtitle: 'ML-DSA / SHA-3 · FIPS 204 / 202',
    color: '#39ff7a',
    glow: 'var(--glow-sign)',
    icon: '✅',
    threat: 'Grover\'s algorithm halves the effective bit-security of symmetric/hash algorithms; Shor\'s breaks ECDSA.',
    solution: 'ML-DSA (Dilithium) lattice signatures + SHA-3 sponge hashing provide quantum-resistant integrity guarantees.',
    params: ['ML-DSA-44/65/87', 'SHA3-256/384/512', 'SHAKE128 / SHAKE256'],
    cta: 'Try Sign + Hash Demo →',
  },
  {
    id: 'non-repudiation' as TabId,
    title: 'Non-Repudiation',
    subtitle: 'SLH-DSA · FIPS 205',
    color: '#ff9500',
    glow: 'var(--glow-slh)',
    icon: '📜',
    threat: 'If a signer\'s private key is compromised, they could deny signing. Stateful hash-based schemes require counter management.',
    solution: 'SLH-DSA is stateless — each signature uses fresh randomness internally. Security relies only on hash functions.',
    params: ['SHA2-128s/f, 192s/f, 256s/f', 'SHAKE-128s/f, 192s/f, 256s/f'],
    cta: 'Try SLH-DSA Demo →',
  },
  {
    id: 'mesh' as TabId,
    title: 'Trust Agent Mesh',
    subtitle: 'Migration Pipeline',
    color: '#c855ff',
    glow: 'var(--glow-hash)',
    icon: '⬡',
    threat: 'You can\'t flip a switch and go PQC overnight — existing systems depend on classical crypto.',
    solution: 'The trust-agent mesh runs legacy and PQC paths in parallel, attests both, then prunes legacy once all dual-path checks pass.',
    params: ['Dual-path attestation', 'Prune gate', 'Pure-PQC state'],
    cta: 'Run Mesh Attest →',
  },
]

const fipsTable = [
  { std: 'FIPS 202', algo: 'SHA-3 / SHAKE', purpose: 'Hash functions', basis: 'Keccak sponge', cat: '1–5' },
  { std: 'FIPS 203', algo: 'ML-KEM', purpose: 'Key encapsulation', basis: 'Module lattice (MLWE)', cat: '1, 3, 5' },
  { std: 'FIPS 204', algo: 'ML-DSA', purpose: 'Digital signatures', basis: 'Module lattice (MSIS/MLWE)', cat: '2, 3, 5' },
  { std: 'FIPS 205', algo: 'SLH-DSA', purpose: 'Digital signatures', basis: 'Hash functions (XMSS+FORS)', cat: '1, 3, 5' },
]

export default function OverviewPanel({ onNavigate }: Props) {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '2.5rem' }}>
      {/* Hero */}
      <div style={{
        background: 'linear-gradient(135deg, #0d1525 0%, #111c30 50%, #0d1525 100%)',
        border: '1px solid var(--border)',
        borderRadius: 16,
        padding: '3rem 2.5rem',
        position: 'relative',
        overflow: 'hidden',
      }}>
        {/* Background glow blobs */}
        <div style={{
          position: 'absolute', top: -80, right: -80,
          width: 300, height: 300, borderRadius: '50%',
          background: 'radial-gradient(circle, rgba(0,212,255,0.08) 0%, transparent 70%)',
          pointerEvents: 'none',
        }} />
        <div style={{
          position: 'absolute', bottom: -60, left: 100,
          width: 200, height: 200, borderRadius: '50%',
          background: 'radial-gradient(circle, rgba(200,85,255,0.07) 0%, transparent 70%)',
          pointerEvents: 'none',
        }} />

        <div style={{ maxWidth: 680, position: 'relative' }}>
          <div style={{
            fontFamily: 'var(--font-mono)', fontSize: '0.7rem',
            color: 'var(--accent-kem)', letterSpacing: '0.2em',
            marginBottom: '0.8rem',
          }}>
            // NIST POST-QUANTUM CRYPTOGRAPHY STANDARDS 2024
          </div>
          <h1 style={{
            fontFamily: 'var(--font-sans)', fontWeight: 800,
            fontSize: 'clamp(1.8rem, 4vw, 2.8rem)',
            lineHeight: 1.1,
            letterSpacing: '-0.02em',
            marginBottom: '1rem',
          }}>
            Quantum computers will break<br />
            <span style={{ color: 'var(--accent-kem)' }}>RSA, ECDH, and ECDSA.</span>
            <br />Here's what replaces them.
          </h1>
          <p style={{ color: 'var(--text-secondary)', fontSize: '1rem', lineHeight: 1.7, marginBottom: '1.5rem' }}>
            This interactive lab runs <strong style={{ color: 'var(--text-primary)' }}>real FIPS-standardized algorithms</strong> on
            a Go backend (cloudflare/circl · golang.org/x/crypto). Explore how ML-KEM, ML-DSA, SLH-DSA, and SHA-3 protect
            confidentiality, integrity, and non-repudiation in a post-quantum world.
          </p>
          <div style={{ display: 'flex', gap: '0.75rem', flexWrap: 'wrap' }}>
            <button className="btn btn-kem" onClick={() => onNavigate('confidentiality')}>
              Start with Confidentiality →
            </button>
            <button className="btn btn-ghost" onClick={() => onNavigate('mesh')}>
              View Migration Pipeline
            </button>
          </div>
        </div>
      </div>

      {/* FIPS Standards table */}
      <section>
        <div style={{ marginBottom: '1rem', display: 'flex', alignItems: 'center', gap: '0.75rem' }}>
          <h2 style={{ fontWeight: 700, fontSize: '1rem' }}>NIST PQC Standards at a Glance</h2>
          <span className="badge badge-cat">2024 Final</span>
        </div>
        <div style={{ overflowX: 'auto' }}>
          <table style={{ width: '100%', borderCollapse: 'collapse', fontFamily: 'var(--font-mono)', fontSize: '0.78rem' }}>
            <thead>
              <tr style={{ borderBottom: '1px solid var(--border)' }}>
                {['Standard', 'Algorithm', 'Purpose', 'Mathematical Basis', 'NIST Cat'].map(h => (
                  <th key={h} style={{ textAlign: 'left', padding: '0.6rem 1rem', color: 'var(--text-muted)', fontWeight: 700, letterSpacing: '0.08em', fontSize: '0.65rem', textTransform: 'uppercase' }}>{h}</th>
                ))}
              </tr>
            </thead>
            <tbody>
              {fipsTable.map((row, i) => (
                <tr key={i} style={{ borderBottom: '1px solid var(--border)' }}>
                  <td style={{ padding: '0.7rem 1rem' }}>
                    <span className={`badge ${row.std === 'FIPS 203' ? 'badge-kem' : row.std === 'FIPS 204' ? 'badge-sign' : row.std === 'FIPS 205' ? 'badge-slh' : 'badge-hash'}`}>{row.std}</span>
                  </td>
                  <td style={{ padding: '0.7rem 1rem', fontWeight: 700, color: 'var(--text-primary)' }}>{row.algo}</td>
                  <td style={{ padding: '0.7rem 1rem', color: 'var(--text-secondary)' }}>{row.purpose}</td>
                  <td style={{ padding: '0.7rem 1rem', color: 'var(--text-muted)' }}>{row.basis}</td>
                  <td style={{ padding: '0.7rem 1rem', color: 'var(--text-secondary)' }}>{row.cat}</td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>
      </section>

      {/* Concept cards */}
      <section>
        <h2 style={{ fontWeight: 700, fontSize: '1rem', marginBottom: '1rem' }}>Security Properties — Interactive Demos</h2>
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(280px, 1fr))', gap: '1rem' }}>
          {concepts.map(c => (
            <div
              key={c.id}
              className="card"
              style={{ cursor: 'pointer', transition: 'all 0.2s', borderColor: 'var(--border)' }}
              onClick={() => onNavigate(c.id)}
              onMouseEnter={e => {
                const el = e.currentTarget as HTMLDivElement
                el.style.borderColor = c.color
                el.style.boxShadow = c.glow
              }}
              onMouseLeave={e => {
                const el = e.currentTarget as HTMLDivElement
                el.style.borderColor = 'var(--border)'
                el.style.boxShadow = 'none'
              }}
            >
              <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start', marginBottom: '0.75rem' }}>
                <div>
                  <div style={{ fontSize: '1.5rem', marginBottom: '0.3rem' }}>{c.icon}</div>
                  <h3 style={{ fontWeight: 800, color: c.color, fontSize: '1.1rem' }}>{c.title}</h3>
                  <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', marginTop: '0.2rem' }}>{c.subtitle}</div>
                </div>
              </div>

              <div style={{ marginBottom: '0.75rem' }}>
                <div style={{ fontSize: '0.7rem', color: 'var(--accent-danger)', fontFamily: 'var(--font-mono)', marginBottom: '0.3rem', textTransform: 'uppercase', letterSpacing: '0.08em' }}>⚠ Quantum Threat</div>
                <p style={{ fontSize: '0.78rem', color: 'var(--text-secondary)', lineHeight: 1.5 }}>{c.threat}</p>
              </div>

              <div style={{ marginBottom: '1rem' }}>
                <div style={{ fontSize: '0.7rem', color: c.color, fontFamily: 'var(--font-mono)', marginBottom: '0.3rem', textTransform: 'uppercase', letterSpacing: '0.08em' }}>✓ PQC Solution</div>
                <p style={{ fontSize: '0.78rem', color: 'var(--text-secondary)', lineHeight: 1.5 }}>{c.solution}</p>
              </div>

              <div style={{ display: 'flex', flexWrap: 'wrap', gap: '0.3rem', marginBottom: '1rem' }}>
                {c.params.map(p => (
                  <span key={p} className="badge badge-cat" style={{ fontSize: '0.6rem' }}>{p}</span>
                ))}
              </div>

              <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: c.color, fontWeight: 700 }}>
                {c.cta}
              </div>
            </div>
          ))}
        </div>
      </section>

      {/* Why PQC matters */}
      <section style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem' }}>
        <div className="card">
          <h3 style={{ fontWeight: 700, marginBottom: '0.75rem', fontSize: '0.9rem', color: 'var(--accent-warn)' }}>⚠ The Harvest Now, Decrypt Later Threat</h3>
          <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
            Nation-state adversaries are <strong style={{ color: 'var(--text-primary)' }}>recording encrypted traffic today</strong> to decrypt it once
            a cryptographically-relevant quantum computer (CRQC) exists. Data with a 10+ year
            sensitivity horizon is <strong style={{ color: 'var(--accent-danger)' }}>already at risk</strong>.
          </p>
        </div>
        <div className="card">
          <h3 style={{ fontWeight: 700, marginBottom: '0.75rem', fontSize: '0.9rem', color: 'var(--accent-sign)' }}>✓ Migrate with Confidence</h3>
          <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
            NIST finalized FIPS 203/204/205 in 2024. The algorithms in this lab are <strong style={{ color: 'var(--text-primary)' }}>production-grade</strong> —
            backed by cloudflare/circl (constant-time, formally reviewed) and golang.org/x/crypto.
            Start migrating now with the dual-path trust-mesh pattern.
          </p>
        </div>
      </section>
    </div>
  )
}
