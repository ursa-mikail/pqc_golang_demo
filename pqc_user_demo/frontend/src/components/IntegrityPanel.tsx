import { useState } from 'react'
import { api } from '../hooks/useApi'
import { SignResult, HashResult } from '../types'

const DSA_VARIANTS = [
  { value: '44', label: 'ML-DSA-44', cat: 2, pk: '1312B', sig: '2420B', desc: 'Dilithium2 · Category 2 · faster verification, smaller signatures' },
  { value: '65', label: 'ML-DSA-65', cat: 3, pk: '1952B', sig: '3309B', desc: 'Dilithium3 · Category 3 · NIST recommended balance' },
  { value: '87', label: 'ML-DSA-87', cat: 5, pk: '2592B', sig: '4627B', desc: 'Dilithium5 · Category 5 · highest security' },
]

const HASH_VARIANTS = [
  { value: 'sha3-256', label: 'SHA3-256', bits: 256, cat: 1, desc: 'Standard sponge hash, 128-bit quantum security' },
  { value: 'sha3-384', label: 'SHA3-384', bits: 384, cat: 3, desc: '192-bit quantum security' },
  { value: 'sha3-512', label: 'SHA3-512', bits: 512, cat: 5, desc: '256-bit quantum security — maximum' },
  { value: 'shake128', label: 'SHAKE128', bits: 256, cat: 1, desc: 'XOF — variable output, rate=168' },
  { value: 'shake256', label: 'SHAKE256', bits: 512, cat: 5, desc: 'XOF — variable output, rate=136' },
]

function fmtDur(ns: number) {
  if (ns < 1_000_000) return `${(ns/1000).toFixed(1)} µs`
  return `${(ns/1_000_000).toFixed(2)} ms`
}

export default function IntegrityPanel() {
  const [dsaVariant, setDsaVariant] = useState('65')
  const [hashVariant, setHashVariant] = useState('sha3-256')
  const [message, setMessage] = useState('Integrity-protected message — FIPS 204 / 202')
  const [signResult, setSignResult] = useState<SignResult | null>(null)
  const [hashResult, setHashResult] = useState<HashResult | null>(null)
  const [loadingSign, setLoadingSign] = useState(false)
  const [loadingHash, setLoadingHash] = useState(false)
  const [activeDemo, setActiveDemo] = useState<'sign' | 'hash'>('sign')

  const runSign = async () => {
    setLoadingSign(true)
    try {
      const data = await api.signDemo(dsaVariant, message)
      if (data.error) throw new Error(data.error)
      setSignResult(data)
    } catch (e) { console.error(e) }
    finally { setLoadingSign(false) }
  }

  const runHash = async () => {
    setLoadingHash(true)
    try {
      const data = await api.hashDemo(hashVariant, message)
      if (data.error) throw new Error(data.error)
      setHashResult(data)
    } catch (e) { console.error(e) }
    finally { setLoadingHash(false) }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      {/* Theory */}
      <div className="card" style={{ borderColor: 'rgba(57,255,122,0.2)', background: 'linear-gradient(135deg, #0d1525, #0d1e18)' }}>
        <div style={{ display: 'flex', gap: '2rem', flexWrap: 'wrap' }}>
          <div style={{ flex: 1, minWidth: 240 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '0.75rem' }}>
              <span style={{ fontSize: '1.2rem' }}>✅</span>
              <h2 style={{ fontWeight: 800, color: 'var(--accent-sign)', fontSize: '1.1rem' }}>Integrity: ML-DSA + SHA-3</h2>
              <span className="badge badge-sign">FIPS 204</span>
              <span className="badge badge-cat">FIPS 202</span>
            </div>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7, marginBottom: '0.75rem' }}>
              <strong style={{ color: 'var(--text-primary)' }}>ML-DSA</strong> (formerly Dilithium) is a lattice-based digital signature algorithm.
              Any modification to the signed message produces a verification failure — even a single bit flip.
              Security relies on the hardness of the <strong style={{ color: 'var(--accent-sign)' }}>Module Short Integer Solution (MSIS)</strong> problem.
            </p>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
              <strong style={{ color: 'var(--text-primary)' }}>SHA-3</strong> (Keccak sponge) is quantum-resistant at &gt;256-bit output —
              Grover's algorithm only provides a √N speedup, halving effective security bits.
              SHA3-256 gives 128-bit quantum security; SHA3-512 gives 256-bit.
            </p>
          </div>
          <div style={{ minWidth: 200 }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', marginBottom: '0.5rem', textTransform: 'uppercase', letterSpacing: '0.1em' }}>Tamper Resistance</div>
            <div style={{ fontSize: '0.75rem', fontFamily: 'var(--font-mono)', lineHeight: 1.8, color: 'var(--text-secondary)' }}>
              <div>Original msg → Sig <span style={{ color: 'var(--accent-sign)' }}>✓ VALID</span></div>
              <div style={{ color: 'var(--text-muted)', fontSize: '0.7rem' }}>  (1 bit changed)</div>
              <div>Tampered msg → Sig <span style={{ color: 'var(--accent-danger)' }}>✗ INVALID</span></div>
            </div>
            <div style={{ marginTop: '0.75rem', padding: '0.5rem 0.75rem', background: 'rgba(57,255,122,0.08)', borderRadius: 6, fontSize: '0.7rem', fontFamily: 'var(--font-mono)', color: 'var(--accent-sign)' }}>
              Avalanche: 1 bit flip → ~50% output bits change
            </div>
          </div>
        </div>
      </div>

      {/* Demo switcher */}
      <div style={{ display: 'flex', gap: '0.5rem' }}>
        {(['sign', 'hash'] as const).map(d => (
          <button
            key={d}
            onClick={() => setActiveDemo(d)}
            style={{
              background: activeDemo === d ? 'rgba(57,255,122,0.15)' : 'var(--bg-card)',
              border: `1px solid ${activeDemo === d ? 'rgba(57,255,122,0.4)' : 'var(--border)'}`,
              borderRadius: 8, padding: '0.6rem 1.2rem',
              cursor: 'pointer', color: activeDemo === d ? 'var(--accent-sign)' : 'var(--text-secondary)',
              fontFamily: 'var(--font-mono)', fontWeight: 700, fontSize: '0.8rem',
              transition: 'all 0.15s',
            }}
          >
            {d === 'sign' ? '⬡ ML-DSA Signature Demo' : '# SHA-3 Hash Demo'}
          </button>
        ))}
      </div>

      {/* Shared input */}
      <div className="card">
        <label style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', display: 'block', marginBottom: '0.4rem' }}>
          Message
        </label>
        <textarea
          value={message}
          onChange={e => setMessage(e.target.value)}
          rows={2}
          placeholder="Enter message to sign / hash…"
        />
      </div>

      {/* DSA Demo */}
      {activeDemo === 'sign' && (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
          <div style={{ display: 'flex', gap: '0.75rem', flexWrap: 'wrap' }}>
            {DSA_VARIANTS.map(v => (
              <button
                key={v.value}
                onClick={() => { setDsaVariant(v.value); setSignResult(null) }}
                style={{
                  background: dsaVariant === v.value ? 'rgba(57,255,122,0.12)' : 'var(--bg-card)',
                  border: `1px solid ${dsaVariant === v.value ? 'rgba(57,255,122,0.4)' : 'var(--border)'}`,
                  borderRadius: 8, padding: '0.7rem 1.1rem',
                  cursor: 'pointer', textAlign: 'left', transition: 'all 0.15s', color: 'var(--text-primary)',
                }}
              >
                <div style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: dsaVariant === v.value ? 'var(--accent-sign)' : 'var(--text-primary)', fontSize: '0.82rem' }}>{v.label}</div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.62rem', color: 'var(--text-muted)', marginTop: '0.2rem' }}>Cat {v.cat} · {v.sig} sig</div>
              </button>
            ))}
          </div>
          <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: 'var(--text-muted)', padding: '0.4rem 0.75rem', background: 'var(--bg-panel)', borderRadius: 6, borderLeft: '3px solid var(--accent-sign)' }}>
            {DSA_VARIANTS.find(v => v.value === dsaVariant)?.desc}
          </div>
          <div>
            <button className="btn btn-sign" onClick={runSign} disabled={loadingSign}>
              {loadingSign ? <span className="spinner" /> : '▶'}
              {loadingSign ? 'Signing…' : 'Sign & Verify'}
            </button>
          </div>

          {signResult && (
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', animation: 'fadeUp 0.4s ease' }}>
              <div className="card">
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Signature Info</div>
                <table style={{ width: '100%', fontSize: '0.75rem', fontFamily: 'var(--font-mono)', borderCollapse: 'collapse' }}>
                  {[
                    ['Algorithm', signResult.algorithm],
                    ['Sig bytes', `${signResult.signatureBytes}`],
                    ['Duration', fmtDur(signResult.durationNs)],
                  ].map(([k, v]) => (
                    <tr key={k} style={{ borderBottom: '1px solid var(--border)' }}>
                      <td style={{ padding: '0.4rem 0', color: 'var(--text-muted)', width: '45%' }}>{k}</td>
                      <td style={{ padding: '0.4rem 0', color: 'var(--accent-sign)' }}>{v}</td>
                    </tr>
                  ))}
                </table>
                <div style={{ marginTop: '0.75rem', fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', marginBottom: '0.3rem' }}>Signature (truncated)</div>
                <div className="hex-display">{signResult.signatureHex}</div>
              </div>

              <div className="card">
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>Tamper Test</div>
                {[
                  { label: 'Original message', ok: signResult.verified, note: 'Must be ✓' },
                  { label: 'Tampered message (1-bit flip)', ok: signResult.tamperedVerify, note: 'Must be ✗' },
                ].map(item => (
                  <div key={item.label} style={{
                    display: 'flex', alignItems: 'center', justifyContent: 'space-between',
                    padding: '0.6rem 0.75rem', borderRadius: 6,
                    background: item.ok === !item.note.includes('✗') ? 'rgba(57,255,122,0.08)' : item.ok ? 'rgba(255,51,102,0.08)' : 'rgba(57,255,122,0.08)',
                    border: `1px solid ${(!item.ok && item.note.includes('✗')) || (item.ok && !item.note.includes('✗')) ? 'rgba(57,255,122,0.2)' : 'rgba(255,51,102,0.2)'}`,
                    marginBottom: '0.5rem',
                  }}>
                    <div>
                      <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--text-primary)' }}>{item.label}</div>
                      <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.62rem', color: 'var(--text-muted)' }}>{item.note}</div>
                    </div>
                    <span style={{ fontFamily: 'var(--font-mono)', fontSize: '1rem', color: item.ok ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>
                      {item.ok ? '✓' : '✗'}
                    </span>
                  </div>
                ))}
                <div style={{ marginTop: '0.75rem', display: 'flex', flexDirection: 'column', gap: '0.3rem' }}>
                  {signResult.steps.map((s, i) => (
                    <div key={i} style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', padding: '0.3rem 0', borderBottom: '1px solid var(--border)', display: 'flex', gap: '0.5rem' }}>
                      <span style={{ color: 'var(--accent-sign)' }}>{i+1}.</span>
                      <span>{s.replace(/^\d+\. /, '')}</span>
                    </div>
                  ))}
                </div>
              </div>
            </div>
          )}
        </div>
      )}

      {/* Hash Demo */}
      {activeDemo === 'hash' && (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem' }}>
          <div style={{ display: 'flex', gap: '0.5rem', flexWrap: 'wrap' }}>
            {HASH_VARIANTS.map(v => (
              <button
                key={v.value}
                onClick={() => { setHashVariant(v.value); setHashResult(null) }}
                style={{
                  background: hashVariant === v.value ? 'rgba(200,85,255,0.12)' : 'var(--bg-card)',
                  border: `1px solid ${hashVariant === v.value ? 'rgba(200,85,255,0.4)' : 'var(--border)'}`,
                  borderRadius: 8, padding: '0.6rem 1rem',
                  cursor: 'pointer', textAlign: 'left', transition: 'all 0.15s', color: 'var(--text-primary)',
                }}
              >
                <div style={{ fontFamily: 'var(--font-mono)', fontWeight: 700, color: hashVariant === v.value ? 'var(--accent-hash)' : 'var(--text-primary)', fontSize: '0.8rem' }}>{v.label}</div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.6rem', color: 'var(--text-muted)', marginTop: '0.15rem' }}>Cat {v.cat} · {v.bits}-bit</div>
              </button>
            ))}
          </div>
          <button className="btn btn-hash" onClick={runHash} disabled={loadingHash}>
            {loadingHash ? <span className="spinner" /> : '#'}
            {loadingHash ? 'Hashing…' : 'Compute Hash'}
          </button>

          {hashResult && (
            <div style={{ display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '1rem', animation: 'fadeUp 0.4s ease' }}>
              <div className="card">
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.5rem' }}>Digest</div>
                <div className="hex-display" style={{ marginBottom: '0.75rem', wordBreak: 'break-all' }}>{hashResult.digestHex}</div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: 'var(--text-muted)' }}>
                  {hashResult.digestLen * 8}-bit output · {fmtDur(hashResult.durationNs)}
                </div>
              </div>
              <div className="card">
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.5rem' }}>Avalanche Effect</div>
                {hashResult.avalancheDemo.map((line, i) => (
                  <div key={i} style={{ fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: i === 0 ? 'var(--accent-sign)' : 'var(--accent-danger)', padding: '0.35rem 0', borderBottom: i === 0 ? '1px solid var(--border)' : 'none' }}>
                    {line}
                  </div>
                ))}
                <p style={{ fontSize: '0.72rem', color: 'var(--text-muted)', marginTop: '0.5rem', fontFamily: 'var(--font-mono)' }}>
                  A single bit flip in input produces a completely different digest — SHA-3's avalanche property ensures no correlation.
                </p>
              </div>
            </div>
          )}
        </div>
      )}
    </div>
  )
}
