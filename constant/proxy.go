package constant

const (
	HTTP ProxyType = iota
	HTTPS
	SOCKS
)

type ProxyType int

func (pt ProxyType) String() string {
	switch pt {
	case HTTPS:
		return "https"
	case SOCKS:
		return "socks"
	default:
		return "http"
	}
}

func (pt ProxyType) SetCommand() string {
	switch pt {
	case HTTPS:
		return "-setsecurewebproxy"
	case SOCKS:
		return "-setsocksfirewallproxy"
	default:
		return "-setwebproxy"
	}
}

func (pt ProxyType) StopCommand() string {
	switch pt {
	case HTTPS:
		return "-setsecurewebproxystate"
	case SOCKS:
		return "-setsocksfirewallproxystate"
	default:
		return "-setwebproxystate"
	}
}

func (pt ProxyType) ShowCommand() string {
	switch pt {
	case HTTPS:
		return "-getsecurewebproxy"
	case SOCKS:
		return "-getsocksfirewallproxy"
	default:
		return "-getwebproxy"
	}
}
