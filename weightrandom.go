package weightrandom

import (
	"math"
	"math/rand"
	"time"
)

func DoRandom(allItems []Items, number int) *[]Items {

	items := []Items{}
	if len(allItems) < number {
		number = len(allItems)
	}
	for i := 0; i < number; i++ {
		item, index := getItemFromRandom(allItems)
		items = append(items, *item)
		allItems = append(allItems[:index], allItems[(index+1):]...)
	}
	return &items
}

type choice struct {
	Start int
	End   int
	Item  Items
}

func getItemFromRandom(allItems []Items) (*Items, int) {
	rand.Seed(time.Now().UnixNano())
	allChoice, allAverageWeight := calculateWeight(allItems)
	pickNumber := rand.Intn(allAverageWeight)
	index := 0
	result := Items{}
	for indx, item := range *allChoice {
		if pickNumber < item.End && pickNumber >= item.Start {
			result = item.Item
			index = indx
		}
	}
	return &result, index

}

func calculateWeight(allItems []Items) (*[]choice, int) {
	result := []choice{}
	allWeight := 0
	for _, item := range allItems {
		allWeight += item.Weight
	}
	start := 0
	for _, item := range allItems {
		roundAverage := start + round(float64(item.Weight*100/allWeight), 0.5, 0)
		result = append(result, choice{
			Start: start,
			End:   roundAverage,
			Item:  item,
		})
		start = roundAverage
	}
	allAverageWeight := start
	return &result, allAverageWeight
}

func round(val float64, roundOn float64, places int) (newVal int) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = int(round / pow)
	return
}
