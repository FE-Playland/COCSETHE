package model

import (
	"math/big"
	"testing"
)

var address = "0x411e2725F2Cd5358AEe921bEB4De8299BC0181EB"

func TestBalanceToReadableEtherWillNotFail(t *testing.T) {

	cases := []struct {
		weiAmount   *big.Int
		etherAmount *big.Float
	}{
		{
			weiAmount:   big.NewInt(10000000000000000),
			etherAmount: big.NewFloat(0.01),
		},
		{

			weiAmount:   big.NewInt(1000000000000000000),
			etherAmount: big.NewFloat(1),
		},
		{

			weiAmount:   big.NewInt(100000000000000),
			etherAmount: big.NewFloat(0.0001),
		},
	}

	for _, c := range cases {

		balance := Balance{
			Amount: c.weiAmount,
			Unit:   "wei",
		}

		etherAmount := balance.ToEther()

		if etherAmount.String() != c.etherAmount.String() {
			t.FailNow()
		}

	}

}
