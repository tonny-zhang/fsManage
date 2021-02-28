package utils

import (
	"encoding/json"
	"errors"
	"net"
	"net/http"
	"strings"
)

func GetExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIPFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}
func getIPFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

// InitFormValue 初始化参数
func InitFormValue(r *http.Request) {
	if r.Method == "POST" {
		ct := r.Header["Content-Type"]
		if len(ct) > 0 {
			if strings.HasPrefix(ct[0], "application/json") {
				var params = make(map[string]string, 0)
				e := json.NewDecoder(r.Body).Decode(&params)
				if nil == e {
					// init form
					r.ParseForm()

					// add value
					for k, v := range params {
						r.Form.Add(k, v)
					}
					return
				}
			} else if ct[0] == "application/x-www-form-urlencoded" {
				r.ParseForm()
				return
			}
		}
		r.ParseMultipartForm(32 << 20)
	} else {
		r.ParseForm()
	}
}
