package utils

import (
	"fmt"
	"math/big"
	"math/rand"
	"net"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	seed := int64(42)
	rand.New(rand.NewSource(seed))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomIpInCidr() string {
	_, ipnet, _ := net.ParseCIDR("192.168.0.0/16")
	ip := ipnet.IP
	mask := ipnet.Mask
	ones, bits := mask.Size()

	if ones == bits {
		return ip.String()
	}

	ipStart := new(big.Int).SetBytes(ip)
	ipEnd := new(big.Int).Add(ipStart, big.NewInt(1<<(uint(bits-ones))-1))
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	randomIP := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), new(big.Int).Sub(ipEnd, ipStart))
	randomIP.Add(randomIP, ipStart)

	return net.IP(randomIP.Bytes()).String()
}

func RandomMAC() string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	// Generate a random MAC address in the format XX:XX:XX:XX:XX:XX
	mac := [6]byte{
		0x02, // Locally administered and unicast MAC address
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
		byte(rand.Intn(256)),
	}
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", mac[0], mac[1], mac[2], mac[3], mac[4], mac[5])
}

func RandomIP() string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	// Generate a random IP in the format X.X.X.X
	ip := net.IPv4(
		byte(rand.Intn(256)), // First octet
		byte(rand.Intn(256)), // Second octet
		byte(rand.Intn(256)), // Third octet
		byte(rand.Intn(256)), // Fourth octet
	)
	// Ip := sql.NullString{
	// 	String: ip.String(),
	// 	Valid:  true, // Set to true to indicate the value is not null
	// }
	return ip.String()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	// name := sql.NullString{
	// 	String: RandomString(6),
	// 	Valid:  true, // Set to true to indicate the value is not null
	// }
	return RandomString(6)
}

func RandomVlan() int64 {
	return RandomInt(0, 1024)
	// return sql.NullInt64{
	// 	Int64: RandomInt(0, 1024),
	// 	Valid: true,
	// }
	// return RandomInt(1500, 9000)
}

func RandomVxlan() int64 {
	return RandomInt(0, 16777215)
}
func RandomMtu() int64 {
	// return sql.NullInt64{
	// 	Int64: RandomInt(1500, 9000),
	// 	Valid: true,
	// }
	return RandomInt(1500, 9000)
}

func RandomAsNo() int64 {
	return RandomInt(1, 65535)
}

func RandomCIDR() string {
	seed := time.Now().UnixNano()
	rand.New(rand.NewSource(seed))
	// Generate random values for the first three octets
	octet1 := rand.Intn(256) // First octet (0-255)
	octet2 := rand.Intn(256) // Second octet (0-255)
	octet3 := rand.Intn(256) // Third octet (0-255)

	// Generate a random subnet mask prefix between 8 and 30
	prefix := rand.Intn(23) + 8 // Random prefix in range [8, 30]

	// Combine to form the CIDR
	return fmt.Sprintf("%d.%d.%d.0/%d", octet1, octet2, octet3, prefix)
}
