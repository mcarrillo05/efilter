package efilter

import "testing"

type TestAddress struct {
	Address string
	Name    string
}

func (a TestAddress) GetAddress() string {
	return a.Address
}

func TestValidateAddress(t *testing.T) {
	AddToBlackList("@hotmail.com", "prueba@", "mcarrillom05@gmail.com")
	list := []Address{
		TestAddress{Address: "mcarrillom05@github.com", Name: "Mario"},
		TestAddress{Address: "mcarrillom05.com", Name: "Mario"},
		TestAddress{Address: "mcarrillom05@", Name: "Mario"},
		TestAddress{Address: "mcarrillom05@github.", Name: "Mario"},
		TestAddress{Address: "mcarrillom05@.com", Name: "Mario"},
		TestAddress{Address: "mcarrillom05@gmail.com", Name: "Mario"},
		TestAddress{Address: "mcarrillom05@hotmail.com", Name: "Mario"},
		TestAddress{Address: "prueba@github.com", Name: "Mario"},
	}
	list = Filter(list)
	validEmail := false
	for _, interfaceL := range list {
		l := interfaceL.(TestAddress)
		t.Log(l.Name, l.GetAddress())
		switch l.GetAddress() {
		case "mcarrillom05@github.com":
			validEmail = true
		case "mcarrillom05.com":
			t.Log("mail must be incorrect")
			t.Fail()
		case "mcarrillom05@":
			t.Log("mail must be incorrect")
			t.Fail()
		case "mcarrillom05@github.":
			t.Log("mail must be incorrect")
			t.Fail()
		case "mcarrillom05@.com":
			t.Log("mail must be incorrect")
			t.Fail()
		case "mcarrillom05@gmail.com":
			t.Log("mail must be in black list")
			t.Fail()
		case "mcarrillom05@hotmail.com":
			t.Log("mail must be in black list")
			t.Fail()
		case "prueba@github.com":
			t.Log("mail must be in black list")
			t.Fail()
		}
	}
	if !validEmail {
		t.Log("mail must be correct")
		t.Fail()
	}
}
