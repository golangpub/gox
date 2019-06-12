package gox

import (
	"net"

	"github.com/gopub/log"
)

// GetIP gets ip address such as "10.0.0.1"
//func GetIP() string {
//	ifaces, err := net.Interfaces()
//	if err != nil {
//		return err.Error()
//	}
//	// handle err
//	var pdpIP string
//	for _, i := range ifaces {
//		//en0-wifi, en1-wireless lan, eth0?
//		isEn := strings.Index(i.Name, "e") == 0 || strings.Index(i.Name, "w") == 0
//
//		//unknown interface?
//		if !isEn && strings.Index(i.Name, "pdp") != 0 {
//			continue
//		}
//
//		addrs, err := i.Addrs()
//		if err != nil {
//			return err.Error()
//		}
//		// handle err
//		for _, addr := range addrs {
//			switch v := addr.(type) {
//			case *net.IPNet:
//				if strings.Count(v.IP.String(), ".") == 3 {
//					if isEn {
//						//
//						return v.IP.String()
//					}
//					pdpIP = v.IP.String()
//				}
//			case *net.IPAddr:
//				//ignore
//			}
//			// process IP address
//		}
//	}
//
//	return pdpIP
//}

// Get preferred outbound ip of this machine
func GetIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetMacAddrs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	addrs := make([]string, len(ifaces))
	for i, ifa := range ifaces {
		addrs[i] = ifa.HardwareAddr.String()
	}

	return addrs, nil
}
