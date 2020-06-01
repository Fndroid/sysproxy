package constant

const (
	WiFi NetworkType = iota
	Ethernet
	ThunderboltEthernet
	Unknown
)

type NetworkType int

func (nt NetworkType) String() string {
	switch nt {
	case WiFi:
		return "Wi-Fi"
	case Ethernet:
		return "Ethernet"
	case ThunderboltEthernet:
		return "ThThunderbolt Ethernet"
	default:
		return "Unknown"
	}
}
