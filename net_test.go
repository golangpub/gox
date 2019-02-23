package gox

import "testing"

func TestGetIP(t *testing.T) {
	ip := GetIP()
	t.Log(ip)
}

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP()
	t.Log(ip)
}
