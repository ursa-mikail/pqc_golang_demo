import { useState } from 'react'
import { api } from '../hooks/useApi'
import { MeshAttestResult } from '../types'

function fmtTime(iso: string) {
  return new Date(iso).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' })
}

const PIPELINE_STEPS = [
  { id: 1, label: 'Legacy Layer',       desc: 'RSA-2048 · ECDH P-256 · ECDSA P-384 · AES-256-GCM', color: '#ff3366', icon: '⚠' },
  { id: 2, label: 'Cryptographic Joints', desc: 'KEM Joint · Sign Joint · Hash Joint — both paths exposed', color: '#ffcc00', icon: '⬡' },
  { id: 3, label: 'Trust Agent Mesh',   desc: 'TA-KEM · TA-Sign · TA-Hash — dual-path attestation', color: '#00d4ff', icon: '◈' },
  { id: 4, label: 'PQC Second Paths',   desc: 'ML-KEM-768 · ML-DSA-65 · SLH-DSA-SHA2-128s', color: '#39ff7a', icon: '✓' },
  { id: 5, label: 'Prune Gate',         desc: 'AllPassed? → DeactivateLegacy() on all joints', color: '#c855ff', icon: '◆' },
  { id: 6, label: 'Pure PQC State',     desc: 'Legacy paths removed — quantum-safe only', color: '#39ff7a', icon: '⬢' },
]

export default function MeshPanel() {
  const [result, setResult] = useState<MeshAttestResult | null>(null)
  const [loading, setLoading] = useState(false)
  const [activeStep, setActiveStep] = useState<number | null>(null)

  const runAttest = async () => {
    setLoading(true)
    setResult(null)
    // Animate through pipeline steps
    for (let i = 1; i <= 6; i++) {
      setActiveStep(i)
      await new Promise(r => setTimeout(r, 300))
    }
    try {
      const data = await api.meshAttest()
      setResult(data)
    } catch (e) { console.error(e) }
    finally {
      setLoading(false)
      setActiveStep(null)
    }
  }

  return (
    <div style={{ display: 'flex', flexDirection: 'column', gap: '1.5rem' }}>
      {/* Theory */}
      <div className="card" style={{ borderColor: 'rgba(200,85,255,0.2)', background: 'linear-gradient(135deg, #0d1525, #160d20)' }}>
        <div style={{ display: 'flex', gap: '2rem', flexWrap: 'wrap' }}>
          <div style={{ flex: 1, minWidth: 260 }}>
            <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem', marginBottom: '0.75rem' }}>
              <span style={{ fontSize: '1.2rem' }}>⬡</span>
              <h2 style={{ fontWeight: 800, color: 'var(--accent-hash)', fontSize: '1.1rem' }}>Trust Agent Mesh — Migration Pipeline</h2>
            </div>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7, marginBottom: '0.75rem' }}>
              You cannot flip a switch and go post-quantum overnight. The trust-agent mesh runs
              <strong style={{ color: 'var(--text-primary)' }}> legacy and PQC paths simultaneously</strong>, attests both,
              then removes the legacy paths only after every dual-path compatibility check passes.
            </p>
            <p style={{ fontSize: '0.82rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
              Each <strong style={{ color: 'var(--accent-hash)' }}>Cryptographic Joint</strong> exposes both paths.
              Each <strong style={{ color: 'var(--accent-kem)' }}>Trust Agent</strong> (TA-KEM, TA-Sign, TA-Hash)
              independently attests its joint. The <strong style={{ color: 'var(--accent-hash)' }}>Prune Gate</strong> fires
              only when <code style={{ background: 'var(--bg-panel)', padding: '0 4px', borderRadius: 3, fontSize: '0.75rem' }}>AllPassed == true</code>.
            </p>
          </div>
          <div style={{ minWidth: 200 }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', marginBottom: '0.5rem', textTransform: 'uppercase', letterSpacing: '0.1em' }}>Attestation Logic</div>
            <div className="code-block">
{`func AttestAll() {
  for each pqcPath {
    res.LegacyOK  = joint.LegacyActive
    res.PQCPathOK = joint.PQCActive
    res.Compatible = both active
    if !compatible {
      report.AllPassed = false
    }
  }
}

func Evaluate() {
  if !report.AllPassed { return }
  joint.DeactivateLegacy() // prune!
}`}
            </div>
          </div>
        </div>
      </div>

      {/* Migration pipeline visualiser */}
      <div className="card">
        <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '1rem' }}>
          Migration Pipeline
        </div>
        <div style={{ display: 'flex', flexDirection: 'column', gap: '0.5rem' }}>
          {PIPELINE_STEPS.map((step, i) => {
            const isActive = activeStep === step.id
            const isDone = result && step.id <= 6
            const isPruned = result?.pruned && step.id === 5
            return (
              <div
                key={step.id}
                style={{
                  display: 'flex',
                  alignItems: 'center',
                  gap: '1rem',
                  padding: '0.75rem 1rem',
                  borderRadius: 8,
                  background: isActive ? `${step.color}18` : isDone ? 'var(--bg-panel)' : 'var(--bg-card)',
                  border: `1px solid ${isActive ? step.color + '55' : isDone ? 'var(--border)' : 'var(--border)'}`,
                  transition: 'all 0.3s',
                  boxShadow: isActive ? `0 0 12px ${step.color}33` : 'none',
                }}
              >
                <div style={{
                  width: 32, height: 32, borderRadius: '50%',
                  background: isActive ? step.color : isDone ? step.color + '22' : 'var(--bg-panel)',
                  border: `2px solid ${step.color}`,
                  display: 'flex', alignItems: 'center', justifyContent: 'center',
                  fontSize: '0.75rem', color: isActive ? '#000' : step.color,
                  fontWeight: 700, flexShrink: 0,
                  transition: 'all 0.3s',
                }}>
                  {isActive ? <span style={{ animation: 'spin 0.7s linear infinite' }}>◌</span> : step.id}
                </div>
                <div style={{ display: 'flex', gap: '0.5rem', alignItems: 'baseline', flex: 1 }}>
                  <span style={{ fontWeight: 700, fontSize: '0.82rem', color: isActive ? step.color : 'var(--text-primary)' }}>{step.label}</span>
                  <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.68rem', color: 'var(--text-muted)' }}>{step.desc}</span>
                </div>
                {isPruned && (
                  <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--accent-sign)', fontWeight: 700 }}>✓ PRUNED</span>
                )}
                {i < PIPELINE_STEPS.length - 1 && (
                  <div style={{
                    position: 'absolute',
                    left: 24, top: '100%',
                    width: 2, height: 8,
                    background: step.color + '44',
                  }} />
                )}
              </div>
            )
          })}
        </div>
      </div>

      {/* Run button */}
      <div>
        <button
          className="btn"
          onClick={runAttest}
          disabled={loading}
          style={{ background: 'var(--accent-hash)', color: '#000', fontSize: '0.85rem', padding: '0.8rem 2rem' }}
        >
          {loading ? <><span className="spinner" /> Running Attestation…</> : '⬡ Run Full Mesh Attestation'}
        </button>
      </div>

      {/* Results */}
      {result && (
        <div style={{ display: 'flex', flexDirection: 'column', gap: '1rem', animation: 'fadeUp 0.4s ease' }}>
          {/* Summary banner */}
          <div style={{
            display: 'flex', alignItems: 'center', gap: '1rem',
            padding: '1rem 1.5rem', borderRadius: 10,
            background: result.allPassed ? 'rgba(57,255,122,0.1)' : 'rgba(255,51,102,0.1)',
            border: `1px solid ${result.allPassed ? 'rgba(57,255,122,0.4)' : 'rgba(255,51,102,0.4)'}`,
          }}>
            <span style={{ fontSize: '1.5rem' }}>{result.allPassed ? '✅' : '❌'}</span>
            <div>
              <div style={{ fontWeight: 800, fontSize: '1rem', color: result.allPassed ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>
                Attestation {result.allPassed ? 'PASSED' : 'FAILED'} — AllPassed={String(result.allPassed)}
              </div>
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--text-muted)', marginTop: '0.2rem' }}>
                {fmtTime(result.timestamp)} · {result.results.length} paths checked
                {result.pruned && result.prunedAt && ` · Legacy pruned at ${fmtTime(result.prunedAt)}`}
              </div>
            </div>
            {result.pruned && (
              <div style={{ marginLeft: 'auto', padding: '0.4rem 0.8rem', background: 'rgba(57,255,122,0.15)', border: '1px solid rgba(57,255,122,0.3)', borderRadius: 6, fontFamily: 'var(--font-mono)', fontSize: '0.72rem', color: 'var(--accent-sign)', fontWeight: 700 }}>
                PURE PQC MODE ⬢
              </div>
            )}
          </div>

          {/* Path results table */}
          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>
              Dual-Path Attestation Results
            </div>
            <div style={{ overflowX: 'auto' }}>
              <table style={{ width: '100%', borderCollapse: 'collapse', fontSize: '0.78rem', fontFamily: 'var(--font-mono)' }}>
                <thead>
                  <tr style={{ borderBottom: '1px solid var(--border)' }}>
                    {['Provider', 'Standard', 'Cat', 'Legacy ✓', 'PQC ✓', 'Compat', 'Notes'].map(h => (
                      <th key={h} style={{ textAlign: 'left', padding: '0.5rem 0.75rem', color: 'var(--text-muted)', fontSize: '0.62rem', textTransform: 'uppercase', letterSpacing: '0.08em', fontWeight: 700 }}>{h}</th>
                    ))}
                  </tr>
                </thead>
                <tbody>
                  {result.results.map((row, i) => (
                    <tr key={i} style={{ borderBottom: '1px solid var(--border)', animation: `slideIn 0.3s ease ${i * 0.07}s both` }}>
                      <td style={{ padding: '0.6rem 0.75rem', fontWeight: 700, color: 'var(--text-primary)' }}>{row.provider}</td>
                      <td style={{ padding: '0.6rem 0.75rem' }}>
                        <span className={`badge ${row.standard === 'FIPS 203' ? 'badge-kem' : row.standard === 'FIPS 204' ? 'badge-sign' : row.standard === 'FIPS 205' ? 'badge-slh' : 'badge-hash'}`}>{row.standard}</span>
                      </td>
                      <td style={{ padding: '0.6rem 0.75rem', color: 'var(--text-secondary)' }}>{row.category}</td>
                      <td style={{ padding: '0.6rem 0.75rem', color: row.legacyOk ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>{row.legacyOk ? '✓' : '✗'}</td>
                      <td style={{ padding: '0.6rem 0.75px', color: row.pqcPathOk ? 'var(--accent-sign)' : 'var(--accent-danger)' }}>{row.pqcPathOk ? '✓' : '✗'}</td>
                      <td style={{ padding: '0.6rem 0.75px' }}>
                        <span style={{ color: row.compatible ? 'var(--accent-sign)' : 'var(--accent-danger)', fontWeight: 700 }}>
                          {row.compatible ? '✓ OK' : '✗ BLOCKED'}
                        </span>
                      </td>
                      <td style={{ padding: '0.6rem 0.75rem', color: 'var(--text-muted)' }}>{row.notes}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>

          {/* Prune gate result */}
          <div style={{
            padding: '1rem 1.5rem', borderRadius: 10,
            background: result.pruned ? 'rgba(200,85,255,0.1)' : 'rgba(255,204,0,0.1)',
            border: `1px solid ${result.pruned ? 'rgba(200,85,255,0.3)' : 'rgba(255,204,0,0.3)'}`,
          }}>
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.5rem' }}>Prune Gate Decision</div>
            {result.pruned ? (
              <div>
                <div style={{ fontWeight: 700, color: 'var(--accent-hash)', marginBottom: '0.4rem' }}>
                  ✓ Gate OPEN — Legacy paths deactivated
                </div>
                <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--text-secondary)', lineHeight: 1.8 }}>
                  <div>KEM joint legacy pruned:  <span style={{ color: 'var(--accent-sign)' }}>true</span></div>
                  <div>Sign joint legacy pruned: <span style={{ color: 'var(--accent-sign)' }}>true</span></div>
                  <div>Hash joint legacy pruned: <span style={{ color: 'var(--accent-sign)' }}>true</span></div>
                  {result.prunedAt && <div style={{ marginTop: '0.3rem', color: 'var(--text-muted)' }}>Pruned at: {result.prunedAt}</div>}
                </div>
                <div style={{ marginTop: '0.75rem', padding: '0.6rem 0.75rem', background: 'rgba(57,255,122,0.08)', borderRadius: 6, fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--accent-sign)' }}>
                  System is now in PURE PQC mode — all traffic uses ML-KEM · ML-DSA · SLH-DSA · SHA-3
                </div>
              </div>
            ) : (
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.75rem', color: 'var(--accent-warn)' }}>
                ✗ Gate BLOCKED — legacy paths retained until all dual-path attestations pass
              </div>
            )}
          </div>

          {/* What just happened */}
          <div className="card">
            <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)', textTransform: 'uppercase', letterSpacing: '0.1em', marginBottom: '0.75rem' }}>What Just Happened</div>
            <div style={{ display: 'flex', flexDirection: 'column', gap: '0.4rem', fontSize: '0.8rem', color: 'var(--text-secondary)', lineHeight: 1.7 }}>
              {[
                ['1', 'New legacy crypto layer initialised', '(RSA-2048, ECDH P-256, ECDSA P-384, AES-256-GCM)'],
                ['2', 'Three cryptographic joints wired', '(KEM, Sign, Hash — each bridges legacy ↔ PQC)'],
                ['3', 'Trust agent mesh created', '(TA-KEM, TA-Sign, TA-Hash — mode: DualPathAttestation)'],
                ['4', 'PQC providers registered', '(ML-KEM-768, ML-DSA-65, SLH-DSA-SHA2-128s)'],
                ['5', 'Attestation run', '(both paths exercised, compatibility checked per provider)'],
                ['6', result.pruned ? 'Prune gate FIRED' : 'Prune gate BLOCKED', result.pruned ? '(DeactivateLegacy() called on all joints)' : '(legacy retained — re-run to retry)'],
              ].map(([n, action, detail]) => (
                <div key={n} style={{ display: 'flex', gap: '0.75rem', padding: '0.5rem 0', borderBottom: '1px solid var(--border)' }}>
                  <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.7rem', color: 'var(--accent-hash)', minWidth: 16 }}>{n}.</span>
                  <div>
                    <span style={{ color: 'var(--text-primary)', fontWeight: 600 }}>{action}</span>
                    <span style={{ color: 'var(--text-muted)', fontFamily: 'var(--font-mono)', fontSize: '0.72rem', marginLeft: '0.5rem' }}>{detail}</span>
                  </div>
                </div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  )
}
