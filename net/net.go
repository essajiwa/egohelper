// Package egohelper contains helpful code (at least for the writer)
package egohelper

import (
	"bytes"
	"strings"
)

// GetDomain : Extract only domain name with the TLD from a host string.
// Subdomain and port (if any) will be cleaned up
func GetDomain(host string) string {
	if strings.Contains(host, ":") {
		host = strings.Split(host, ":")[0]
	}
	domain := strings.SplitAfter(host, ".")
	var buffer bytes.Buffer
	buffer.WriteString(domain[len(domain)-2])
	buffer.WriteString(domain[len(domain)-1])
	return buffer.String()
}
