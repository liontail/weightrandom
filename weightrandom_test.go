package weightrandom

import (
	"testing"

	. "github.com/franela/goblin"
	// . "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := Goblin(t)

	// Assign Item to All Items
	allItems := []Items{
		Items{
			Weight: 10,
			Data:   uint(1),
		},
		Items{
			Weight: 10,
			Data:   uint(2),
		},
		Items{
			Weight: 80,
			Data:   uint(3),
		},
	}

	// Send All Items and No. of item that you want
	result := DoRandom(allItems, 2)

	// Expect Items with most weight
	expect := false
	for _, key := range *result {
		if key.Data == uint(3) {
			expect = true
		}
	}
	g.Describe("Test_WeightRandom", func() {
		g.It("Should Get Item3 That has most Weight", func() {
			g.Assert(expect).Equal(true)
		})
	})
}
