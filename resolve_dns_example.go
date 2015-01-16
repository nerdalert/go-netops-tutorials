package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	ip := "docker.io"
	ip2, _ := dnsLookup(ip)
	fmt.Println("docker.io resolves to --> ", ip2)

	// print a blank line for a space in output
	fmt.Println()

	// now lookup a fake hostname
	l33t := "badH0stN@me.c0m"
	dnsResp, _ := dnsLookup(l33t)
	fmt.Println("badH0stN@me.c0m doesnt resolve and prints a nil pointer, thats uglies --> ", dnsResp)

	// print a blank line for a space in output using print line
	fmt.Printf("\n")

	// now properly handle the nil response from the failed DNS lookup
	dnsResp2, err := dnsLookup(l33t)
	if err != nil {
		log.Printf("lookup failed for [ %s ]", l33t)
	} else {
		fmt.Println("g00gl3.c0m resolves to --> ", dnsResp2)
	}

	// print a blank line for a space in output using print format
	fmt.Printf("\n")
	// now properly handle a good response from a successful lookup
	ovs := "openvswitch.org"
	dnsResp3, err1 := dnsLookup(ovs)
	if err1 != nil {
		log.Println("lookup failed")
	} else {
		fmt.Println("openvswitch.org resolves to --> ", dnsResp3)
	}
}

func dnsLookup(ipStr string) (string, error) {
	ipAddr, err := net.ResolveIPAddr("ip", ipStr)
	return ipAddr.String(), err
}
