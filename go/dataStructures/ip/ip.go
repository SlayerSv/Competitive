package ip

import (
	"fmt"
)

type IP int

func ParseIP(address string) IP {
	var a, b, c, d int
	_, err := fmt.Sscanf(address, "%d.%d.%d.%d", &a, &b, &c, &d)
	if err != nil {
		panic(err)
	}
	var ip int
	ip += a << 24
	ip += b << 16
	ip += c << 8
	ip += d
	return IP(ip)
}

func (ip IP) String() string {
	i := int(ip)
	return fmt.Sprintf("%d.%d.%d.%d", (i & 0xFF000000)>>24, (i & 0xFF0000)>>16, (i & 0xFF00)>>8, i & 0xFF)
}

func (ip IP) ByteString() string {
	i := int(ip)
	return fmt.Sprintf("%08b.%08b.%08b.%08b", (i & 0xFF000000)>>24, (i & 0xFF0000)>>16, (i & 0xFF00)>>8, i & 0xFF)
}