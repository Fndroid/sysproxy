package network

import (
	"regexp"
)

type NetworkType struct {
	Name         string
	HardwarePort string
	Device       string
}

func (nt NetworkType) String() string {
	return nt.Name
}

func ParseFromText(t string) []NetworkType {
	re := regexp.MustCompile(`\(\d+\)\s(.+?)\n\(Hardware\sPort:\s(.+?), Device:\s(.+?)\)`)
	sts := re.FindAllStringSubmatch(t, -1)
	res := []NetworkType{}
	for _, st := range sts {
		if len(st) == 4 {
			res = append(res, NetworkType{Name: st[1], HardwarePort: st[2], Device: st[3]})
		}
	}
	return res
}
