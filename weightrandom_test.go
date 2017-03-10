package weightrandom

import (
	"testing"

	. "github.com/franela/goblin"
	// . "github.com/onsi/gomega"
)

func Test(t *testing.T) {
	g := Goblin(t)
	item1 := Items{
		Weight: 10,
		Data:   uint(1),
	}
	item2 := Items{
		Weight: 10,
		Data:   uint(2),
	}
	item3 := Items{
		Weight: 80,
		Data:   uint(3),
	}
	allItems := []Items{}
	allItems = append(allItems, item1)
	allItems = append(allItems, item2)
	allItems = append(allItems, item3)

	result := DoRandom(allItems, 2)

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
