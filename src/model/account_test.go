package model

import (
	"testing"
)

var (
	accountPk = "08346faf1f20d2dcbf63887ca87ffcc211809cad438fbe6e62c587403d2bf4fa"
)

func TestCheckAccountIsValid(t *testing.T) {
	_, err := NewEtherAccount(accountPk)

	if err != nil {
		t.Fatal(err)
	}
}
