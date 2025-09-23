package util

import (
	"testing"
)

func TestSecurityLevel_String(t *testing.T) {
	tests := []struct {
		level    SecurityLevel
		expected string
	}{
		{Level128, "Level 1 (~AES-128)"},
		{Level192, "Level 3 (~AES-192)"},
		{Level256, "Level 5 (~AES-256)"},
		{SecurityLevel(99), "Unknown"},
	}

	for _, test := range tests {
		result := test.level.String()
		if result != test.expected {
			t.Errorf("SecurityLevel(%d).String() = %q, want %q",
				int(test.level), result, test.expected)
		}
	}
}

func TestSecureCompare(t *testing.T) {
	tests := []struct {
		name     string
		a, b     []byte
		expected bool
	}{
		{"identical bytes", []byte("hello"), []byte("hello"), true},
		{"different bytes", []byte("hello"), []byte("world"), false},
		{"different lengths", []byte("hello"), []byte("hi"), false},
		{"empty slices", []byte{}, []byte{}, true},
		{"one empty", []byte("test"), []byte{}, false},
		{"nil slices", nil, nil, true},
		{"one nil", []byte("test"), nil, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := SecureCompare(test.a, test.b)
			if result != test.expected {
				t.Errorf("SecureCompare(%v, %v) = %v, want %v",
					test.a, test.b, result, test.expected)
			}
		})
	}
}

func TestHexDump(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		maxBytes int
		expected string
	}{
		{"short data", []byte{0x01, 0x02, 0x03}, 10, "010203"},
		{"truncated data", []byte{0x01, 0x02, 0x03, 0x04}, 2, "0102..."},
		{"empty data", []byte{}, 10, ""},
		{"exact length", []byte{0xaa, 0xbb}, 2, "aabb"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HexDump(test.data, test.maxBytes)
			if result != test.expected {
				t.Errorf("HexDump(%v, %d) = %q, want %q",
					test.data, test.maxBytes, result, test.expected)
			}
		})
	}
}

func TestBytesToHex(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		truncate int
		expected string
	}{
		{"no truncation", []byte{0x01, 0x02, 0x03}, 0, "010203"},
		{"with truncation", []byte{0x01, 0x02, 0x03, 0x04}, 2, "0102... (4 bytes total)"},
		{"truncate larger than data", []byte{0xaa, 0xbb}, 10, "aabb"},
		{"empty data", []byte{}, 5, ""},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := BytesToHex(test.data, test.truncate)
			if result != test.expected {
				t.Errorf("BytesToHex(%v, %d) = %q, want %q",
					test.data, test.truncate, result, test.expected)
			}
		})
	}
}

// Benchmark SecureCompare vs naive comparison
func BenchmarkSecureCompare(b *testing.B) {
	data1 := make([]byte, 1000)
	data2 := make([]byte, 1000)
	// Fill with same data
	for i := range data1 {
		data1[i] = byte(i % 256)
		data2[i] = byte(i % 256)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SecureCompare(data1, data2)
	}
}
