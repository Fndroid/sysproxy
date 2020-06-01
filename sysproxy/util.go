package sysproxy

import (
	"strconv"
	"strings"
)

type ProxyInfo struct {
	Enabled bool
	Server string
	Port int
}

func format(s string) ProxyInfo {
	lines := strings.Split(s, "\n")
	pi := ProxyInfo{}
	for _, line := range lines {
		kv := splitTrim(line, ":")
		if len(kv) == 2 {
			switch kv[0] {
			case "Enabled":
				pi.Enabled = kv[1] == "Yes"
			case "Server":
				pi.Server = kv[1]
			case "Port":
				portInt, err := strconv.Atoi(kv[1])
				if err == nil {
					pi.Port = portInt
				}
			}
		}
	}
	return pi
}

func splitTrim(s string, sep string) []string {
	ps := strings.Split(s, sep)
	o := []string{}
	for _, p := range ps {
		o = append(o, strings.TrimSpace(p))
	}
	return o
}