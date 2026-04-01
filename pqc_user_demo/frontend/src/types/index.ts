export interface AlgorithmInfo {
  id: string
  name: string
  standard: string
  category: number
  parameterSet: string
  purpose: 'confidentiality' | 'integrity' | 'non-repudiation'
  description: string
}

export interface KEMResult {
  algorithm: string
  standard: string
  securityCategory: number
  publicKeyHex: string
  ciphertextHex: string
  sharedSecretHex: string
  verifiedMatch: boolean
  ciphertextBytes: number
  sharedSecretLen: number
  durationNs: number
  steps: string[]
}

export interface SignResult {
  algorithm: string
  standard: string
  securityCategory: number
  messageHex: string
  signatureHex: string
  signatureBytes: number
  verified: boolean
  tamperedVerify: boolean
  durationNs: number
  steps: string[]
}

export interface SLHResult {
  algorithm: string
  standard: string
  securityCategory: number
  hashFamily: string
  signatureBytes: number
  verified: boolean
  tamperedVerify: boolean
  durationNs: number
  steps: string[]
}

export interface HashResult {
  algorithm: string
  standard: string
  securityCategory: number
  digestHex: string
  digestLen: number
  avalancheDemo: string[]
  durationNs: number
  steps: string[]
}

export interface MeshPathResult {
  provider: string
  standard: string
  category: number
  legacyOk: boolean
  pqcPathOk: boolean
  compatible: boolean
  notes: string
}

export interface MeshAttestResult {
  timestamp: string
  allPassed: boolean
  results: MeshPathResult[]
  pruned: boolean
  prunedAt?: string
}

export type TabId = 'overview' | 'confidentiality' | 'integrity' | 'non-repudiation' | 'mesh'
