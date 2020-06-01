package sysproxy

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	C "github.com/Fndroid/sysproxy/constant"
)

const COMMAND = "networksetup"

func networkType() C.NetworkType {
	for _, t := range []C.NetworkType{C.Ethernet, C.WiFi, C.ThunderboltEthernet} {
		if testWebProxy(t) {
			return t
		}
	}
	return C.Unknown
}

func testWebProxy(nt C.NetworkType) bool {
	cmd := exec.Command(COMMAND, "-getwebproxy", nt.String())
	err := cmd.Run()
	if err != nil {
		return false
	}
	return true
}

func SetBypass(domains []string) error {
	nt := networkType()
	if nt == C.Unknown {
		return errors.New("unknown network type")
	}
	args := append([]string{"-setproxybypassdomains", nt.String()}, domains...)
	cmd := exec.Command(COMMAND, args...)
	return cmd.Run()
}

func SetProxy(pt C.ProxyType, server string, port int) error {
	nt := networkType()
	if nt == C.Unknown {
		return errors.New("unknown network type")
	}
	args := []string{pt.SetCommand(), nt.String(), server, strconv.Itoa(port)}
	cmd := exec.Command(COMMAND, args...)
	return cmd.Run()
}

func StopProxy(pt C.ProxyType) error {
	nt := networkType()
	if nt == C.Unknown {
		return errors.New("unknown network type")
	}
	args := []string{pt.StopCommand(), nt.String(), "off"}
	cmd := exec.Command(COMMAND, args...)
	return cmd.Run()
}

func ShowProxy() (string, error) {
	nt := networkType()
	if nt == C.Unknown {
		return "", errors.New("unknown network type")
	}
	result := []string{}
	for _, pt := range []C.ProxyType{C.HTTP, C.HTTPS, C.SOCKS} {
		args := []string{pt.ShowCommand(), nt.String()}
		cmd := exec.Command(COMMAND, args...)
		out, err := cmd.CombinedOutput()
		if err != nil {
			continue
		}
		o := format(string(out))
		if o.Enabled {
			result = append(result, fmt.Sprintf("%s=%s:%d", pt.String(), o.Server, o.Port))
		}
	}
	return strings.Join(result, "\n"), nil
}
