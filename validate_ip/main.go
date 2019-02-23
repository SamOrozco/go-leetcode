package main

import (
	"golang.org/x/text/runes"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const ipv6 = "IPv6"
const ipv6SegmentCount = 8
const ipv6SegmentLength = 4
const ipv4 = "IPv4"
const ipv4SegmentCount = 4
const neither = "Neither"
const period = "."
const colon = ":"

func main() {
	if validIPAddress("172.145.255.121") != ipv4 {
		print("fail")
		os.Exit(1)
	}

	if validIPAddress("2001:0db8:85a3:0000:0000:8a2e:0370:7334") != ipv6 {
		print("fail")
		os.Exit(1)
	}
	println("pass")
}

func validIPAddress(IP string) string {
	if !valid(IP) { // empty
		return neither
	}

	if strings.Contains(IP, period) {
		if isValidIPv4(IP) {
			return ipv4
		}
	} else if strings.Contains(IP, colon) {
		if isValidIPv6(IP) {
			return ipv6
		}
	}
	return neither

}

func valid(val string) bool {
	return len(val) > 0
}

func isValidIPv4(ip string) bool {
	segments := strings.Split(ip, period)
	if len(segments) != ipv4SegmentCount {
		return false
	}
	for _, seg := range segments {
		if !validIpv4Segment(seg) {
			return false
		}
	}
	return true
}

func validIpv4Segment(seg string) bool {
	if val, err := strconv.Atoi(seg); err == nil {
		return val > 0 && val < 256
	}
	return false
}

func isValidIPv6(ip string) bool {
	runes := []rune(ip)
	currentRunes := make([]rune, ipv6SegmentLength)
	currSegCount := 0
	segmentCount := 0
	for i := 0; i < len(ip); i++ {
		currentRunes[currSegCount] = runes[i]
		currSegCount++
		if currSegCount == ipv6SegmentLength {
			return false
		}
		if runes[i] ==  ':' {
			if currSegCount > ipv6SegmentLength {
				return false
			}
			if !validIpv6Segment(currentRunes) {
				return false
			}
			segmentCount++
		}
	}
	return segmentCount == 8
}

func validIpv6Segment(runes []rune) bool {
	return isValidHex(runes)
}

func isValidHex(rs []rune) bool {
	set := runes.In(unicode.ASCII_Hex_Digit)
	for _, r := range rs {
		if !set.Contains(r) {
			return false
		}
	}
	return true
}
