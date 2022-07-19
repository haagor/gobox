package main

import (
	"fmt"
	"sort"
)

type userConnection struct {
	id        int
	startDate int
	endDate   int
}

func getConnectionPeak(ucs []userConnection) int {
	sort.Slice(ucs, func(i, j int) bool {
		return ucs[i].startDate < ucs[j].startDate
	})

	peak := 1
	overlapLoss := []userConnection{ucs[0]}
	for i := 1; i < len(ucs); i++ {
		for _, uc := range overlapLoss {
			if ucs[i].startDate > uc.endDate {
				overlapLoss = overlapLoss[1:]
			}
		}
		overlapLoss = append(overlapLoss, ucs[i])
		sort.Slice(overlapLoss, func(i, j int) bool {
			return overlapLoss[i].endDate < overlapLoss[j].endDate
		})

		if peak < len(overlapLoss) {
			peak = len(overlapLoss)
		}

	}
	return peak
}

func main() {
	uc1 := userConnection{id: 1, startDate: 100, endDate: 200}
	uc2 := userConnection{id: 2, startDate: 50, endDate: 101}
	uc3 := userConnection{id: 3, startDate: 150, endDate: 250}
	uc4 := userConnection{id: 4, startDate: 120, endDate: 151}
	uc5 := userConnection{id: 5, startDate: 170, endDate: 210}
	uc6 := userConnection{id: 6, startDate: 199, endDate: 201}
	uc7 := userConnection{id: 7, startDate: 199, endDate: 201}
	uc8 := userConnection{id: 8, startDate: 199, endDate: 201}

	ucs := []userConnection{uc1, uc2, uc3, uc4, uc5, uc6, uc7, uc8}
	fmt.Println(getConnectionPeak(ucs))
}
