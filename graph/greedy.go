package graph

import (
	mapset "github.com/deckarep/golang-set"
)

// 贪婪算法
func greedy(states mapset.Set, stations map[string]mapset.Set) mapset.Set {
	res := mapset.NewSetFromSlice([]interface{}{})
	for len(states.ToSlice()) > 0 {
		bestStation := ""
		bestCovered := mapset.NewSetFromSlice([]interface{}{})
		for station, subStates := range stations {
			coverStates := states.Intersect(subStates)
			if len(coverStates.ToSlice()) > len(bestCovered.ToSlice()) {
				bestCovered = coverStates
				bestStation = station
			}
		}
		res.Add(bestStation)
		states = states.Difference(bestCovered)
	}
	return res
}
