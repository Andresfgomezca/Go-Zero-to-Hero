package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	s := make([]string, len(ip))
	for i, val := range ip {
		s[i] = fmt.Sprint(int(val))

	}
	return strings.Join(s, ".")
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		//this method separates the key of the map next to the string concatened with the address
		fmt.Printf("%v: %v\n", name, ip)
	}
}
