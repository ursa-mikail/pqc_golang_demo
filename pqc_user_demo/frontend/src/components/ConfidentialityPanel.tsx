import { useState } from 'react'
import { api } from '../hooks/useApi'
import { KEMResult } from '../types'

const VARIANTS = [
  { value: '512',  label: 'ML-KEM-512',  cat: 1, pk: '800B',  ct: '768B',  ss: '32B',  desc: 'Security Category 1 — comparable to AES-128. Fastest, smallest keys.' },
  { value: '768',  label: 'ML-KEM-768',  cat: 3, pk: '1184B', ct: '1088B', ss: '32B',  desc: 'Security Category 3 — comparable to AES-192. NIST recommended.' },
  { value: '1024', label: 'ML-KEM-1024', cat: 5, pk: '1568B', ct: '1568B', ss: '32B',  desc: 'Security Category 5 — comparable to AES-256. Maximum security.' },
]

function fmtDur(ns: number) {
  if (ns < 1_000_000) return `${(ns/1000).toFixed(1)} µs`
  return `${(ns/1_000_000).toFixed(2)} ms`
}

export default function ConfidentialityPanel() {
  const [variant, setVariant] = useState('768')
  const [message, setMessage] = useState('Hello from Alice — establish secure channel!')
  const [result, setResult] = useState<KEMResult | null>(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState('')

  const selected = VARIANTS.find(v => v.value === variant)!

  const runDemo = async () => {
    setLoading(true)
    setError('')
    try {
      const data = await api.kemDemo(variant, message)
      if (data.error) throw new Error(data.error)
      setResult(data)
    } catch (e: any) {
      setError(e.message)
    } finally {
      setLoading(false)
    }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      {/* Theory */}
      <div className="card" style={{ borderColor: 'rgba(0,212,255,0.2)', background: 'linear-gradient(135deg, #0d1525, #0d1a2e)' }}>
        <div style={{ display: 'flex', gap: '2rem', flexWrap: 'wrap' }}>
          <div style={{ flex: 1, minWidth: 240 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '0.75rem' }}>
              <span style={{ fontSize: '1.2rem' }}>🔐</span>
              <h2 style={{ fontWeight: 800, color: 'var(--accent-kem)', fontSize: '1.1rem' }}>Confidentiality via ML-KEM</h2>
              <span className="badge badge-kem">FIPS 203</span>
            </div>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7, marginBottom: '0.75rem' }}>
              <strong style={{ color: 'var(--text-primary)' }}>Key Encapsulation Mechanism (KEM)</strong> — Alice uses Bob's public key
              to encapsulate a random shared secret into a ciphertext. Bob decapsulates using his private key.
              Neither party sees the other's private key. The shared secret seeds a symmetric cipher (e.g. AES-GCM).
            </p>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
              ML-KEM security rests on the <strong style={{ color: 'var(--accent-kem)' }}>Module Learning With Errors (MLWE)</strong> problem —
              believed hard for both classical and quantum computers. Backed by cloudflare/circl's constant-time Kyber implementation.
            </p>
          </div>
          <div style={{ flex: 0, minWidth: 200 }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', marginBottom: '0.5rem', textTransform: 'uppercase', letterSpacing: '0.1em' }}>KEM Flow</div>
            {[
              ['Alice', 'pk, sk = KEM.KeyGen()'],
              ['Alice→Bob', 'Sends pk'],
              ['Bob', 'ct, ss = KEM.Enc(pk)'],
              ['Bob→Alice', 'Sends ct'],
              ['Alice', 'ss = KEM.Dec(sk, ct)'],
              ['Both', 'Use ss to key AES-GCM'],
            ].map(([who, op]) => (
              <div key={who} style={{ display: 'flex', gap: '0.5rem', marginBottom: '0.3rem', fontSize: '0.72rem', fontFamily: 'var(--font-mono)' }}>
                <span style={{ color: 'var(--accent-kem)', minWidth: 72, flexShrink: 0 }}>{who}</span>
                <span style={{ color: 'var(--text-secondary)' }}>{op}</span>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Algo selector */}
      <div style={{ display: 'flex', gap: '0.75rem', flexWrap: 'wrap' }}>
        {VARIANTS.map(v => (
          <button
            key={v.value}
            onClick={() => { setVariant(v.value); setResult(null) }}
            style={{
              background: variant === v.value ? 'rgba(0,212,255,0.15)' : 'var(--bg-card)',
              border: `1px solid ${variant === v.value ? 'rgba(0,212,255,0.5)' : 'var(--border)'}`,
              borderRadius: 8,
              padding: '0.75rem 1.2rem',
              cursor: 'pointer',
              textAlign: 'left',
              transition: 'all 0.15s',
              color: 'var(--text-primary)',
              minWidth: 200,
            }}
          >
            <div style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: variant === v.value ? 'var(--accent-kem)' : 'var(--text-primary)', fontSize: '0.85rem' }}>{v.label}</div>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', marginTop: '0.2rem' }}>
              Cat {v.cat} · pk={v.pk} ct={v.ct} ss={v.ss}
            </div>
          </button>
        ))}
      </div>

      {selected && (
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--text-muted)', padding: '0.5rem 0.75rem', background: 'var(--bg-panel)', borderRadius: 6, borderLeft: '3px solid var(--accent-kem)' }}>
          {selected.desc}
        </div>
      )}

      {/* Input */}
      <div className="card">
        <div style={{ marginBottom: '0.75rem' }}>
          <label style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', display: 'block', marginBottom: '0.4rem' }}>
            Payload (simulates data to establish shared context)
          </label>
          <input
            type="text"
            value={message}
            onChange={e => setMessage(e.target.value)}
            placeholder="Enter any message…"
          />
        </div>
        <button className="btn btn-kem" onClick={runDemo} disabled={loading}>
          {loading ? <span className="spinner" /> : '▶'}
          {loading ? 'Generating key pair…' : 'Run ML-KEM Demo'}
        </button>
        {error && <div style={{ marginTop: '0.5rem', color: 'var(--accent-danger)', fontFamily: 'var(--font-mono)', fontSize: '0.75rem' }}>✗ {error}</div>}
      </div>

      {/* Results */}
      {result && (
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', animation: 'fadeUp 0.4s ease' }}>
          {/* Key material */}
          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Key Material</div>
            <table style={{ width: '100%', fontSize: '0.75rem', fontFamily: 'var(--font-mono)', borderCollapse: 'collapse' }}>
              {[
                ['Algorithm', result.algorithm],
                ['Standard', result.standard],
                ['Security Cat', `${result.securityCategory}`],
                ['CT size', `${result.ciphertextBytes} bytes`],
                ['SS length', `${result.sharedSecretLen} bytes`],
                ['Duration', fmtDur(result.durationNs)],
              ].map(([k, v]) => (
                <tr key={k} style={{ borderBottom: '1px solid var(--border)' }}>
                  <td style={{ padding: '0.4rem 0', color: 'var(--text-muted)', width: '45%' }}>{k}</td>
                  <td style={{ padding: '0.4rem 0', color: 'var(--accent-kem)' }}>{v}</td>
                </tr>
              ))}
            </table>
          </div>

          {/* Verification */}
          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Shared Secret Verification</div>
            <div style={{
              display: 'flex', alignItems: 'center', gap: '0.6rem',
              padding: '0.75rem', borderRadius: 8,
              background: result.verifiedMatch ? 'rgba(57,255,122,0.1)' : 'rgba(255,51,102,0.1)',
              border: `1px solid ${result.verifiedMatch ? 'rgba(57,255,122,0.3)' : 'rgba(255,51,102,0.3)'}`,
              marginBottom: '0.75rem',
            }}>
              <span className={`status-dot ${result.verifiedMatch ? 'ok' : 'err'}`} />
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.8rem', fontWeight: 700, color: result.verifiedMatch ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>
                {result.verifiedMatch ? 'Shared secrets MATCH ✓' : 'Mismatch ✗'}
              </span>
            </div>
            <div style={{ marginBottom: '0.5rem', fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase' }}>Ciphertext (truncated)</div>
            <div className="hex-display" style={{ marginBottom: '0.75rem' }}>{result.ciphertextHex}</div>
            <div style={{ marginBottom: '0.5rem', fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase' }}>Shared Secret</div>
            <div className="hex-display" style={{ wordBreak: 'break-all' }}>{result.sharedSecretHex}</div>
          </div>

          {/* Steps */}
          <div className="card" style={{ gridColumn: '1 / -1' }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Step-by-Step Trace</div>
            <div style={{ display: 'flex', flexDirection: 'column', gap: '0.4rem' }}>
              {result.steps.map((step, i) => (
                <div key={i} style={{
                  display: 'flex', alignItems: 'flex-start', gap: '0.75rem',
                  padding: '0.6rem 0.75rem', borderRadius: 6,
                  background: 'var(--bg-panel)',
                  animation: `slideIn 0.3s ease ${i * 0.08}s both`,
                }}>
                  <span style={{ color: 'var(--accent-kem)', fontFamily: 'var(--font-mono)', fontSize: '0.7rem', minWidth: 20 }}>{i + 1}.</span>
                  <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--text-secondary)' }}>{step.replace(/^\d+\. /, '')}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}

      {/* Size comparison */}
      <div className="card">
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '1rem' }}>
          Key Size Comparison: Classical vs Post-Quantum
        </div>
        <div style={{ display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(200px, 1fr))', gap: '0.75rem', fontSize: '0.75rem', fontFamily: 'var(--font-mono)' }}>
          {[
            { algo: 'RSA-2048', pk: '256B', broken: true },
            { algo: 'ECDH P-256', pk: '64B', broken: true },
            { algo: 'ML-KEM-512', pk: '800B', broken: false },
            { algo: 'ML-KEM-768', pk: '1184B', broken: false },
            { algo: 'ML-KEM-1024', pk: '1568B', broken: false },
          ].map(a => (
            <div key={a.algo} style={{
              padding: '0.6rem 0.75rem', borderRadius: 6,
              background: a.broken ? 'rgba(255,51,102,0.08)' : 'rgba(0,212,255,0.08)',
              border: `1px solid ${a.broken ? 'rgba(255,51,102,0.2)' : 'rgba(0,212,255,0.2)'}`,
            }}>
              <div style={{ fontWeight: 700, color: a.broken ? 'var(--accent-danger)' : 'var(--accent-kem)', marginBottom: '0.2rem' }}>{a.algo}</div>
              <div style={{ color: 'var(--text-muted)' }}>pub key: {a.pk}</div>
              <div style={{ color: a.broken ? 'var(--accent-danger)' : 'var(--accent-sign)', fontSize: '0.65rem', marginTop: '0.2rem' }}>
                {a.broken ? '✗ Broken by Shor' : '✓ Quantum-safe'}
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  )
}
