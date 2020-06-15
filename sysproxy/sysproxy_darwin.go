package sysproxy

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	C "github.com/Fndroid/sysproxy/constant"
	N "github.com/Fndroid/sysproxy/network"
)

const COMMAND = "networksetup"

func networkTypes() ([]N.NetworkType, error) {
	cmd := exec.Command(COMMAND, "-listnetworkserviceorder")
	out, err := cmd.CombinedOutput()
	nts := []N.NetworkType{}
	if err != nil {
		return nts, err
	}
	for _, nt := range N.ParseFromText(string(out)) {
		err := testNetwork(nt)
		if err == nil {
			nts = append(nts, nt)
		}
	}
	return nts, nil
}

func testNetwork(nt N.NetworkType) error {
	grep := exec.Command("grep", nt.Device)
	netstat := exec.Command("netstat", "-nr")
	pipe, err := netstat.StdoutPipe()
	if err != nil {
		return err
	}
	defer pipe.Close()
	grep.Stdin = pipe
	netstat.Start()
	out, err := grep.Output()
	if err != nil {
		return err
	}
	if strings.Contains(string(out), "default") {
		return nil
	}
	return errors.New("testNetwork failed")
}

func SetBypass(domains []string) error {
	nts, err := networkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := append([]string{"-setproxybypassdomains", nt.String()}, domains...)
		cmd := exec.Command(COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func SetProxy(pt C.ProxyType, server string, port int) error {
	nts, err := networkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := []string{pt.SetCommand(), nt.String(), server, strconv.Itoa(port)}
		cmd := exec.Command(COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func StopProxy(pt C.ProxyType) error {
	nts, err := networkTypes()
	if err != nil {
		return err
	}
	for _, nt := range nts {
		args := []string{pt.StopCommand(), nt.String(), "off"}
		cmd := exec.Command(COMMAND, args...)
		cmd.Run()
	}
	return nil
}

func ShowProxy() (string, error) {
	nts, err := networkTypes()
	if err != nil {
		return "", err
	}
	result := []string{}
	for _, nt := range nts {
		str := ""
		for _, pt := range []C.ProxyType{C.HTTP, C.HTTPS, C.SOCKS} {
			args := []string{pt.ShowCommand(), nt.String()}
			cmd := exec.Command(COMMAND, args...)
			out, err := cmd.CombinedOutput()
			if err != nil {
				continue
			}
			o := format(string(out))
			if o.Enabled {
				if str == "" {
					str = fmt.Sprintf("%s: ", nt.String())
				}
				str += fmt.Sprintf("%s=%s:%d; ", pt.String(), o.Server, o.Port)
			}
		}
		result = append(result, str)
	}
	return strings.Join(result, "\n"), nil
}
