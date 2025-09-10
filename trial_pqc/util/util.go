package util

import (
	"crypto/subtle"
	"fmt"
)

// SecurityLevel represents different post-quantum security levels
type SecurityLevel int

const (
	Level128 SecurityLevel = iota // ~AES-128 security
	Level192                      // ~AES-192 security
	Level256                      // ~AES-256 security
)

// String returns a human-readable description of the security level
func (s SecurityLevel) String() string {
	switch s {
	case Level128:
		return "Level 1 (~AES-128)"
	case Level192:
		return "Level 3 (~AES-192)"
	case Level256:
		return "Level 5 (~AES-256)"
	default:
		return "Unknown"
	}
}

// SecureCompare performs a constant-time comparison of two byte slices
func SecureCompare(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

// HexDump returns a formatted hex dump of data (useful for debugging)
func HexDump(data []byte, maxBytes int) string {
	if len(data) > maxBytes {
		return fmt.Sprintf("%x...", data[:maxBytes])
	}
	return fmt.Sprintf("%x", data)
}

// BytesToHex converts bytes to hex string with optional truncation
func BytesToHex(data []byte, truncate int) string {
	if truncate > 0 && len(data) > truncate {
		return fmt.Sprintf("%x... (%d bytes total)", data[:truncate], len(data))
	}
	return fmt.Sprintf("%x", data)
}
