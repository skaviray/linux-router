package utils

import (
	"database/sql"
	"fmt"
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

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
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

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomName() sql.NullString {
	name := sql.NullString{
		String: RandomString(6),
		Valid:  true, // Set to true to indicate the value is not null
	}
	return name
}

// RandomMoney generates a random amount of money
func RandomVlan() sql.NullInt64 {
	return sql.NullInt64{
		Int64: RandomInt(0, 1024),
		Valid: true,
	}
	// return RandomInt(1500, 9000)
}

func RandomVxlan() sql.NullInt64 {
	return sql.NullInt64{
		Int64: RandomInt(0, 9000),
		Valid: true,
	}
	// return RandomInt(1500, 9000)
}

func RandomMtu() sql.NullInt64 {
	return sql.NullInt64{
		Int64: RandomInt(1500, 9000),
		Valid: true,
	}
	// return RandomInt(1500, 9000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
