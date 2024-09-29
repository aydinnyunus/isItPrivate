package main

import (
	"fmt"
	"net/netip"
)

func main() {
	// Checking 127.0.0.1
	fmt.Println("=== Checking IP: 127.0.0.1 ===")
	ip := netip.MustParseAddr("127.0.0.1")
	fmt.Printf("IsPrivate: %v\n", ip.IsPrivate())          // Output: false
	fmt.Printf("IsLinkLocalUnicast: %v\n", ip.IsLinkLocalUnicast()) // Output: false
	fmt.Printf("IsLoopback: %v\n", ip.IsLoopback())        // Output: true
	fmt.Println()

	// Checking 169.254.169.254
	fmt.Println("=== Checking IP: 169.254.169.254 ===")
	ip = netip.MustParseAddr("169.254.169.254")
	fmt.Printf("IsPrivate: %v\n", ip.IsPrivate())          // Output: false
	fmt.Printf("IsLinkLocalUnicast: %v\n", ip.IsLinkLocalUnicast()) // Output: true
	fmt.Printf("IsLoopback: %v\n", ip.IsLoopback())        // Output: false
	fmt.Println()

	// Checking 8.8.8.8
	fmt.Println("=== Checking IP: 8.8.8.8 ===")
	ip2 := netip.MustParseAddr("8.8.8.8")
	fmt.Printf("IsPrivate: %v\n", ip2.IsPrivate())         // Output: false
	fmt.Printf("IsLinkLocalUnicast: %v\n", ip2.IsLinkLocalUnicast()) // Output: false
	fmt.Printf("IsLoopback: %v\n", ip2.IsLoopback())       // Output: false
	fmt.Println()
}
