package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	res := ""
	
	for ind, el := range ip {
		
		if ind == (len(ip) - 1) {
			res += fmt.Sprint(el)
		} else {
			res += fmt.Sprint(el, ".")
		}
	}
	
	return fmt.Sprintf("%q", res)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
