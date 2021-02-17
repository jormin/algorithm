package graph

// 狄克斯特拉算法
func dijkstra(graph map[string]map[string]int, costs map[string]int, parents map[string]string) []string {
	processed := map[string]int{}
	for {
		n := findLowestCostNode(costs, processed)
		if n == "" {
			break
		}
		for s, c := range graph[n] {
			newCost := costs[n] + c
			if costs[s] < 0 || costs[s] > newCost {
				costs[s] = newCost
				parents[s] = n
			}
		}
		processed[n] = 1
	}
	n := "final"
	res := []string{n}
	for {
		if n == "start" {
			break
		}
		res = append(res, parents[n])
		n = parents[n]
	}
	return reverse(res)
}

// 找出花费最小的节点
func findLowestCostNode(costs map[string]int, processed map[string]int) string {
	lowestCost := 9999999999
	lowestCostNode := ""
	for k, v := range costs {
		if v > 0 && v < lowestCost {
			if _, ok := processed[k]; !ok {
				lowestCost = v
				lowestCostNode = k
			}
		}
	}
	return lowestCostNode
}

// 反转
func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
