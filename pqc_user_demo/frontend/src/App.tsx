import { useState } from 'react'
import { TabId } from './types'
import OverviewPanel from './components/OverviewPanel'
import ConfidentialityPanel from './components/ConfidentialityPanel'
import IntegrityPanel from './components/IntegrityPanel'
import NonRepudiationPanel from './components/NonRepudiationPanel'
import MeshPanel from './components/MeshPanel'

const TABS: { id: TabId; label: string; glyph: string; color: string }[] = [
  { id: 'overview',          label: 'Overview',         glyph: '⬡', color: '#e8f0ff' },
  { id: 'confidentiality',   label: 'Confidentiality',  glyph: '⬢', color: '#00d4ff' },
  { id: 'integrity',         label: 'Integrity',        glyph: '◈', color: '#39ff7a' },
  { id: 'non-repudiation',   label: 'Non-Repudiation',  glyph: '◆', color: '#ff9500' },
  { id: 'mesh',              label: 'Trust Mesh',       glyph: '⬡', color: '#c855ff' },
]

export default function App() {
  const [tab, setTab] = useState<TabId>('overview')

  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
      {/* Header */}
      <header style={{
        borderBottom: '1px solid var(--border)',
        background: 'rgba(8,12,20,0.95)',
        backdropFilter: 'blur(20px)',
        position: 'sticky',
        top: 0,
        zIndex: 100,
      }}>
        {/* Top bar */}
        <div style={{
          display: 'flex',
          alignItems: 'center',
          justifyContent: 'space-between',
          padding: '0.9rem 2rem',
          borderBottom: '1px solid var(--border)',
        }}>
          <div style={{ display: 'flex', alignItems: 'center', gap: '0.75rem' }}>
            <div style={{
              width: 36, height: 36,
              background: 'linear-gradient(135deg, #00d4ff22, #c855ff22)',
              border: '1px solid var(--border-bright)',
              borderRadius: 8,
              display: 'flex', alignItems: 'center', justifyContent: 'center',
              fontSize: '1.1rem',
            }}>⬡</div>
            <div>
              <div style={{ fontFamily: 'var(--font-sans)', fontWeight: 800, fontSize: '1rem', letterSpacing: '-0.01em' }}>
                PQC<span style={{ color: 'var(--accent-kem)' }}>Lab</span>
              </div>
              <div style={{ fontFamily: 'var(--font-mono)', fontSize: '0.6rem', color: 'var(--text-muted)', letterSpacing: '0.15em' }}>
                POST-QUANTUM CRYPTOGRAPHY
              </div>
            </div>
          </div>

          <div style={{ display: 'flex', gap: '0.5rem', alignItems: 'center' }}>
            <span className="badge badge-cat">FIPS 202</span>
            <span className="badge badge-kem">FIPS 203</span>
            <span className="badge badge-sign">FIPS 204</span>
            <span className="badge badge-slh">FIPS 205</span>
            <div style={{
              marginLeft: '0.5rem',
              padding: '0.3rem 0.7rem',
              background: 'rgba(57,255,122,0.1)',
              border: '1px solid rgba(57,255,122,0.3)',
              borderRadius: 4,
              display: 'flex', alignItems: 'center', gap: '0.4rem',
              fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--accent-sign)',
            }}>
              <span className="status-dot ok" />
              QUANTUM-SAFE
            </div>
          </div>
        </div>

        {/* Nav tabs */}
        <nav style={{ display: 'flex', gap: 0, padding: '0 2rem' }}>
          {TABS.map(t => (
            <button
              key={t.id}
              onClick={() => setTab(t.id)}
              style={{
                background: 'none',
                border: 'none',
                borderBottom: tab === t.id ? `2px solid ${t.color}` : '2px solid transparent',
                color: tab === t.id ? t.color : 'var(--text-muted)',
                padding: '0.85rem 1.2rem',
                cursor: 'pointer',
                fontFamily: 'var(--font-sans)',
                fontWeight: 600,
                fontSize: '0.82rem',
                display: 'flex',
                alignItems: 'center',
                gap: '0.4rem',
                transition: 'all 0.15s',
                whiteSpace: 'nowrap',
                letterSpacing: '0.02em',
              }}
            >
              <span style={{ fontSize: '0.9rem' }}>{t.glyph}</span>
              {t.label}
            </button>
          ))}
        </nav>
      </header>

      {/* Main content */}
      <main style={{ flex: 1, padding: '2rem', maxWidth: 1280, margin: '0 auto', width: '100%' }}>
        {tab === 'overview'        && <OverviewPanel onNavigate={setTab} />}
        {tab === 'confidentiality' && <ConfidentialityPanel />}
        {tab === 'integrity'       && <IntegrityPanel />}
        {tab === 'non-repudiation' && <NonRepudiationPanel />}
        {tab === 'mesh'            && <MeshPanel />}
      </main>

      <footer style={{
        borderTop: '1px solid var(--border)',
        padding: '1rem 2rem',
        display: 'flex',
        justifyContent: 'space-between',
        alignItems: 'center',
      }}>
        <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)' }}>
          Backend: Go · cloudflare/circl (Kyber / Dilithium) · golang.org/x/crypto/sha3
        </span>
        <span style={{ fontFamily: 'var(--font-mono)', fontSize: '0.65rem', color: 'var(--text-muted)' }}>
          NIST PQC Standards · FIPS 202 / 203 / 204 / 205
        </span>
      </footer>
    </div>
  )
}
