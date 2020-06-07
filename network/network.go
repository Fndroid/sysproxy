package network

import (
	"regexp"
)

type NetworkType struct {
	Name   string
	Device string
}

func (nt NetworkType) String() string {
	return nt.Name
}

func ParseFromText(t string) []NetworkType {
	re := regexp.MustCompile(`\(Hardware\sPort:\s(.+?), Device:\s(.+?)\)`)
	sts := re.FindAllStringSubmatch(t, -1)
	res := []NetworkType{}
	for _, st := range sts {
		if len(st) == 3 {
			res = append(res, NetworkType{Name: st[1], Device: st[2]})
		}
	}
	return res
}
