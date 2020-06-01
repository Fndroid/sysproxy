package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"

	C "github.com/Fndroid/sysproxy/constant"
	"github.com/Fndroid/sysproxy/sysproxy"
)

var bypass = flag.String("bypass", "", "bypass join by ,")
var httpProxy = flag.String("http", "", "http proxy server and port")
var httpsProxy = flag.String("https", "", "https proxy server and port")
var socksProxy = flag.String("socks", "", "socks proxy server and port")
var stop = flag.Bool("s", false, "disable all proxies")

func main() {
	flag.Parse()

	if bypass != nil {
		dms := strings.Split(*bypass, ",")
		sysproxy.SetBypass(dms)
	}

	if *httpProxy != "" {
		host, port, err := net.SplitHostPort(*httpProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				sysproxy.SetProxy(C.HTTP, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("http proxy format error")
		}
	}

	if *httpsProxy != "" {
		host, port, err := net.SplitHostPort(*httpsProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				sysproxy.SetProxy(C.HTTPS, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("https proxy format error")
		}
	}

	if *socksProxy != "" {
		host, port, err := net.SplitHostPort(*socksProxy)
		if err == nil {
			portInt, err := strconv.Atoi(port)
			if err == nil {
				sysproxy.SetProxy(C.SOCKS, host, portInt)
			} else {
				fmt.Println("port is not a integer")
			}
		} else {
			fmt.Println("socks proxy format error")
		}
	}

	if stop != nil && *stop {
		sysproxy.StopProxy(C.HTTP)
		sysproxy.StopProxy(C.SOCKS)
		sysproxy.StopProxy(C.HTTPS)
	}
}
