package mobile

import (
	"net"

	"github.com/gopub/gox"
)

const (
	NoNetwork = 0
	Cellular  = 1
	WIFI      = 2
)

const (
	IDle       = 0
	Connecting = 1
	Connected  = 2
)

func GetIP() (net.IP, error) {
	return gox.GetOutboundIP()
}
