package efilter

import (
	"errors"
	"regexp"
)

//Address represents a single mail address.
type Address interface {
	GetAddress() string //returns email address
}


//ErrBadFormat is used when address doesn't satisfy regex.
var (
	ErrBadFormat = errors.New("invalid address format")
	emailRegexp  = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

//ValidateFormat uses regex to validate address.
func ValidateFormat(a Address) error {
	if !emailRegexp.MatchString(a.GetAddress()) {
		return ErrBadFormat
	}
	return nil
}

//ValidateAddress address uses ValidateFormat and InBlackList.
func ValidateAddress(a Address) error {
	err := ValidateFormat(a)
	if err != nil {
		return err
	}
	return InBlackList(a)
}


//Filter validate all addresses and returns a new slice with filtered addresses.
func Filter(list []Address) (filtered []Address) {
	for _, a := range list {
		if err := ValidateAddress(a); err == nil {
			filtered = append(filtered, a)
		}
	}
	return filtered
}
