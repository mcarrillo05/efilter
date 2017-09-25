package efilter

import (
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

var (
	ErrMXRecords = errors.New("unable to get MX records")
)

func ValidateHost(a Address) error {
	mx, err := net.LookupMX(strings.Split(a.GetAddress(), "@")[1])
	if err != nil {
		return ErrMXRecords
	}
	for _, m := range mx {
		client, err := smtp.Dial(m.Host + ":25")
		if err != nil {
			return err
		}
		defer client.Close()
		err = client.Hello("gecco.com.mx")
		if err != nil {
			return err
		}
		err = client.Mail("mario.carrillo@gecco.com.mx")
		if err != nil {
			return err
		}
		err = client.Rcpt(a.GetAddress())
		if err == nil {
			return nil
		}
		fmt.Println(err)
	}
	return nil
}
