package ip

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		address string
		want int
		wantByte string
	}{
		{"0.0.0.1", 1, "00000000.00000000.00000000.00000001"},
		{"0.0.1.0", 1<<8, "00000000.00000000.00000001.00000000"},
		{"0.1.0.0", 1<<16, "00000000.00000001.00000000.00000000"},
		{"1.0.0.0", 1<<24, "00000001.00000000.00000000.00000000"},
		{"0.0.0.156", 156, "00000000.00000000.00000000.10011100"},
		{"0.0.255.0", 1<<16-256, "00000000.00000000.11111111.00000000"},
	}
	for _, tt := range tests {
		t.Run(tt.address, func(t *testing.T) {
			a := ParseIP(tt.address)
			if int(a) != tt.want {
				t.Errorf("address: %s, got: %d, want: %d", tt.address, a, tt.want)
			}
			if a.String() != tt.address {
				t.Errorf("ip: %d, got: %s, want: %s", int(a), a.String(), tt.address)
			}
			if a.ByteString() != tt.wantByte {
				t.Errorf("address: %s, got: %s, want: %s", a.String(), a.ByteString(), tt.wantByte)
			}
		})
	}
}