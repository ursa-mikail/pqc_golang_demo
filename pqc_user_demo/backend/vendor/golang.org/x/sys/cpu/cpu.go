// Package cpu provides stub implementations for CPU feature detection.
// This is a minimal stub for vendor builds without the full x/sys package.
package cpu

// X86 contains the supported X86 CPU features.
var X86 struct {
	HasAVX2 bool
}

func init() {
	// Disable AVX2 in stub — falls back to generic (pure-Go) implementation
	X86.HasAVX2 = false
}
