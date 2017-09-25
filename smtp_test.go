package efilter

import "testing"

func TestValidateHost(t *testing.T) {
	a := TestAddress{
		Address: "mcarrillo@gecco.com.mx",
		Name:    "Mario",
	}
	err := ValidateHost(a)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}
