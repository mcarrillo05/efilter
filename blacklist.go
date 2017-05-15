package efilter

import (
	"errors"
	"strings"
)

//ErrBlackList is used when blackList contains address.
var (
	blackList    []string
	ErrBlackList = errors.New("address is in black list")
)

//AddToBlackList can recieve full address, domain or address:
//full address: mcarrillom05@github.com, this address will be removed from list.
//domain: @github.com, all addresses with github.com will be removed from list.
//address: mcarrillom05@, all addresses with mcarrillom05 will be removed.
func AddToBlackList(address ...string) {
	for _, a := range address {
		blackList = append(blackList, strings.ToLower(a))
	}
}

//InBlackList checks if address satisfies blacklist rules.
func InBlackList(a Address) error {
	for _, b := range blackList {
		if strings.Contains(strings.ToLower(a.GetAddress()), b) {
			return ErrBlackList
		}
	}
	return nil
}
