import { useState } from 'react'
import { api } from '../hooks/useApi'
import { SLHResult } from '../types'

const VARIANTS = [
  { value: 'slh-dsa-sha2-128s',  label: 'SHA2-128s',  cat: 1, sigKB: 7.7,   speed: 'small', hash: 'SHA-256', desc: 'Smallest signatures — use for long-lived documents' },
  { value: 'slh-dsa-sha2-128f',  label: 'SHA2-128f',  cat: 1, sigKB: 16.7,  speed: 'fast',  hash: 'SHA-256', desc: 'Fast signing — higher throughput at cost of size' },
  { value: 'slh-dsa-shake-128s', label: 'SHAKE-128s', cat: 1, sigKB: 7.7,   speed: 'small', hash: 'SHAKE256', desc: 'Same as SHA2-128s but uses SHAKE-256 PRF' },
  { value: 'slh-dsa-shake-128f', label: 'SHAKE-128f', cat: 1, sigKB: 16.7,  speed: 'fast',  hash: 'SHAKE256', desc: 'Fast + SHAKE-256 PRF' },
  { value: 'slh-dsa-sha2-256s',  label: 'SHA2-256s',  cat: 5, sigKB: 29.1,  speed: 'small', hash: 'SHA-512', desc: 'Maximum security — Cat 5 with smallest sig at that level' },
]

function fmtDur(ns: number) {
  if (ns < 1_000_000) return `${(ns/1000).toFixed(1)} µs`
  return `${(ns/1_000_000).toFixed(2)} ms`
}

function SizeBar({ kb, max }: { kb: number; max: number }) {
  return (
    <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
      <div style={{ flex: 1, height: 6, background: 'var(--bg-panel)', borderRadius: 3, overflow: 'hidden' }}>
        <div style={{ width: `${(kb / max) * 100}%`, height: '100%', background: 'var(--accent-slh)', borderRadius: 3, transition: 'width 0.5s ease' }} />
      </div>
      <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', minWidth: 50 }}>{kb.toFixed(1)} KB</span>
    </div>
  )
}

export default function NonRepudiationPanel() {
  const [variant, setVariant] = useState('slh-dsa-sha2-128s')
  const [message, setMessage] = useState('I hereby authorize this transaction: TX-2024-PQC-0001')
  const [result, setResult] = useState<SLHResult | null>(null)
  const [loading, setLoading] = useState(false)

  const runDemo = async () => {
    setLoading(true)
    try {
      const data = await api.slhDemo(variant, message)
      if (data.error) throw new Error(data.error)
      setResult(data)
    } catch (e) { console.error(e) }
    finally { setLoading(false) }
  }

  const selected = VARIANTS.find(v => v.value === variant)!

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      {/* Theory */}
      <div className="card" style={{ borderColor: 'rgba(255,149,0,0.2)', background: 'linear-gradient(135deg, #0d1525, #1a1208)' }}>
        <div style={{ display: 'flex', gap: '2rem', flexWrap: 'wrap' }}>
          <div style={{ flex: 1, minWidth: 260 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '0.75rem' }}>
              <span style={{ fontSize: '1.2rem' }}>📜</span>
              <h2 style={{ fontWeight: 800, color: 'var(--accent-slh)', fontSize: '1.1rem' }}>Non-Repudiation via SLH-DSA</h2>
              <span className="badge badge-slh">FIPS 205</span>
            </div>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7, marginBottom: '0.75rem' }}>
              <strong style={{ color: 'var(--text-primary)' }}>Non-repudiation</strong> means the signer cannot later deny having signed a message.
              SLH-DSA (formerly SPHINCS+) achieves this with pure hash-based security — the only assumption is the collision resistance of SHA-2 or SHAKE.
            </p>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7, marginBottom: '0.75rem' }}>
              <strong style={{ color: 'var(--accent-slh)' }}>Stateless</strong> — unlike XMSS, SLH-DSA uses internal randomness (R) for each signature,
              so no per-key counter must be maintained. Each verification reconstructs the signature body from the signer's secret seed
              and the embedded R, then compares in constant time.
            </p>
            <div style={{ background: 'rgba(255,149,0,0.08)', border: '1px solid rgba(255,149,0,0.2)', borderRadius: 6, padding: '0.6rem 0.75rem', fontSize: '0.75rem', fontFamily: 'var(--font-mono)', color: 'var(--accent-slh)' }}>
              Security basis: collision resistance of SHA-2 / SHAKE only — no lattice assumptions needed
            </div>
          </div>
          <div style={{ minWidth: 200 }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', marginBottom: '0.6rem', textTransform: 'uppercase', letterSpacing: '0.1em' }}>How Verification Works</div>
            {[
              ['sig[0:n]', 'Extract R (fresh randomness)'],
              ['PRF(SK.seed, R, msg)', 'Reconstruct FORS+HT body'],
              ['Constant-time cmp', 'Compare reconstructed ↔ received'],
              ['No stored state', 'Fully stateless operation'],
            ].map(([step, desc]) => (
              <div key={step} style={{ marginBottom: '0.4rem', fontSize: '0.72rem', fontFamily: 'var(--font-mono)' }}>
                <span style={{ color: 'var(--accent-slh)', display: 'block' }}>{step}</span>
                <span style={{ color: 'var(--text-muted)', fontSize: '0.65rem', paddingLeft: '0.5rem' }}>↳ {desc}</span>
              </div>
            ))}
          </div>
        </div>
      </div>

      {/* Variant selector */}
      <div>
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.6rem' }}>
          Parameter Set — 12 total in FIPS 205 (s=small-sig / f=fast-sign)
        </div>
        <div style={{ display: 'flex', gap: '0.6rem', flexWrap: 'wrap' }}>
          {VARIANTS.map(v => (
            <button
              key={v.value}
              onClick={() => { setVariant(v.value); setResult(null) }}
              style={{
                background: variant === v.value ? 'rgba(255,149,0,0.12)' : 'var(--bg-card)',
                border: `1px solid ${variant === v.value ? 'rgba(255,149,0,0.4)' : 'var(--border)'}`,
                borderRadius: 8, padding: '0.65rem 1rem',
                cursor: 'pointer', textAlign: 'left', transition: 'all 0.15s', color: 'var(--text-primary)',
              }}
            >
              <div style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: variant === v.value ? 'var(--accent-slh)' : 'var(--text-primary)', fontSize: '0.8rem' }}>
                SLH-DSA-{v.label}
              </div>
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.6rem', color: 'var(--text-muted)', marginTop: '0.15rem' }}>
                Cat {v.cat} · {v.sigKB.toFixed(1)} KB · {v.speed}
              </div>
            </button>
          ))}
        </div>
      </div>

      {selected && (
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: 'var(--text-muted)', padding: '0.5rem 0.75rem', background: 'var(--bg-panel)', borderRadius: 6, borderLeft: '3px solid var(--accent-slh)' }}>
          {selected.desc} · PRF: {selected.hash}
        </div>
      )}

      {/* Sig size comparison */}
      <div className="card">
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Signature Size Comparison</div>
        <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
          {[
            { label: 'ECDSA P-256', kb: 0.07, broken: true },
            { label: 'ML-DSA-65', kb: 3.2, broken: false },
            ...VARIANTS.map(v => ({ label: `SLH-DSA-${v.label}`, kb: v.sigKB, broken: false })),
          ].map(item => (
            <div key={item.label} style={{ display: 'grid', gridTemplateColumns: '160px 1fr', gap: '0.75rem', alignItems: 'center' }}>
              <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: item.broken ? 'var(--accent-danger)' : item.label.startsWith('ML') ? 'var(--accent-sign)' : 'var(--accent-slh)' }}>
                {item.broken ? '✗ ' : ''}{item.label}
              </span>
              <SizeBar kb={item.kb} max={35} />
            </div>
          ))}
        </div>
        <div style={{ marginTop: '0.75rem', fontSize: '0.7rem', fontFamily: 'var(--font-mono)', color: 'var(--text-muted)' }}>
          Trade-off: SLH-DSA has larger signatures than ML-DSA but its security relies only on hash functions (conservative assumption).
        </div>
      </div>

      {/* Demo input */}
      <div className="card">
        <label style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', display: 'block', marginBottom: '0.4rem' }}>
          Message to Sign (non-repudiable document)
        </label>
        <textarea value={message} onChange={e => setMessage(e.target.value)} rows={2} />
        <div style={{ marginTop: '0.75rem' }}>
          <button className="btn btn-slh" onClick={runDemo} disabled={loading}>
            {loading ? <span className="spinner" /> : '📜'}
            {loading ? 'Signing…' : 'Sign with SLH-DSA'}
          </button>
        </div>
      </div>

      {/* Result */}
      {result && (
        <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', animation: 'fadeUp 0.4s ease' }}>
          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Signature Details</div>
            <table style={{ width: '100%', fontSize: '0.75rem', fontFamily: 'var(--font-mono)', borderCollapse: 'collapse' }}>
              {[
                ['Algorithm', result.algorithm],
                ['Standard', result.standard],
                ['Hash family', result.hashFamily],
                ['Security Cat', `${result.securityCategory}`],
                ['Sig size', `${result.signatureBytes.toLocaleString()} bytes`],
                ['Duration', fmtDur(result.durationNs)],
              ].map(([k, v]) => (
                <tr key={k} style={{ borderBottom: '1px solid var(--border)' }}>
                  <td style={{ padding: '0.4rem 0', color: 'var(--text-muted)', width: '40%' }}>{k}</td>
                  <td style={{ padding: '0.4rem 0', color: 'var(--accent-slh)' }}>{v}</td>
                </tr>
              ))}
            </table>
          </div>

          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Verification & Tamper Test</div>
            {[
              { label: 'Original msg verified', ok: result.verified, note: '← must be true' },
              { label: 'Tampered msg (bit flip)', ok: result.tamperedVerify, note: '← must be false' },
            ].map(item => (
              <div key={item.label} style={{
                display: 'flex', alignItems: 'center', justifyContent: 'space-between',
                padding: '0.6rem 0.75rem', borderRadius: 6,
                background: (item.ok && item.note.includes('true')) || (!item.ok && item.note.includes('false'))
                  ? 'rgba(57,255,122,0.08)' : 'rgba(255,51,102,0.08)',
                border: `1px solid ${(item.ok && item.note.includes('true')) || (!item.ok && item.note.includes('false'))
                  ? 'rgba(57,255,122,0.2)' : 'rgba(255,51,102,0.2)'}`,
                marginBottom: '0.5rem',
              }}>
                <div>
                  <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--text-primary)' }}>{item.label}</div>
                  <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.6rem', color: 'var(--text-muted)' }}>{item.note}</div>
                </div>
                <span style={{ fontSize: '1rem', color: item.ok ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>
                  {item.ok ? '✓' : '✗'}
                </span>
              </div>
            ))}
            <div style={{ marginTop: '0.75rem', display: 'flex', flexDirection: 'column', gap: '0.25rem' }}>
              {result.steps.map((s, i) => (
                <div key={i} style={{ fontFamily: 'var(--font-mono)', fontSize: '0.68rem', color: 'var(--text-muted)', display: 'flex', gap: '0.5rem', padding: '0.25rem 0', borderBottom: '1px solid var(--border)' }}>
                  <span style={{ color: 'var(--accent-slh)' }}>{i+1}.</span>
                  <span>{s.replace(/^\d+\. /, '')}</span>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
