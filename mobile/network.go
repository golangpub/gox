package mobile

import (
	"net"

	"github.com/gopub/gox"
)

func GetIP() (net.IP, error) {
	return gox.GetIP()
}
