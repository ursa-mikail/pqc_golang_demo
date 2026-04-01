const API = '/api'

export const api = {
  getAlgorithms: () => fetch(`${API}/algorithms`).then(r => r.json()),

  kemDemo: (variant: string, message: string) =>
    fetch(`${API}/kem/demo`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ variant, message }),
    }).then(r => r.json()),

  signDemo: (variant: string, message: string) =>
    fetch(`${API}/sign/demo`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ variant, message }),
    }).then(r => r.json()),

  slhDemo: (variant: string, message: string) =>
    fetch(`${API}/slhdsa/demo`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ variant, message }),
    }).then(r => r.json()),

  hashDemo: (variant: string, data: string) =>
    fetch(`${API}/hash/demo`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ variant, data }),
    }).then(r => r.json()),

  meshAttest: () =>
    fetch(`${API}/mesh/attest`, { method: 'POST' }).then(r => r.json()),
}
