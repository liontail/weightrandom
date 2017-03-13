package weightrandom

import (
	"testing"

	"github.com/Sirupsen/logrus"
	. "github.com/franela/goblin"
	// . "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := Goblin(t)

	// Assign Item to All Items
	allItems := []Items{
		Items{
			Weight: 10,
			Data:   "Item1",
		},
		Items{
			Weight: 10,
			Data:   "Item2",
		},
		Items{
			Weight: 80,
			Data:   "Item3",
		},
	}

	// Send All Items and No. of item that you want
	result := DoRandom(allItems, 2)

	// Expect Items with most weight
	expect := false
	for _, key := range *result {
		if key.Data == "Item3" {
			expect = true
		}
	}
	g.Describe("Test_WeightRandom", func() {
		g.It("Should Get Item3 That has most Weight", func() {
			logrus.Infoln("Result : ", result)
			g.Assert(expect).Equal(true)
		})
	})
}
