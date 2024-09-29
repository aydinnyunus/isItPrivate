package main

import (
	"fmt"
	"net/netip"
)

func main() {
	fmt.Println("Language: Go")
	ip := netip.MustParseAddr("::1")
	fmt.Printf("=== Checking IP: %s ===\n", ip)
	fmt.Printf("IsPrivate: %v\n", ip.IsPrivate())
	fmt.Printf("IsLoopback: %v\n", ip.IsLoopback())
	fmt.Println()
}
