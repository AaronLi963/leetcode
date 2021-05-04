package main

import "fmt"

type Node struct {
	Name  string
	Edges map[string]float64
}

var (
	nodeMap = map[string]*Node{}
)

func converse(rates [][]interface{}, from, to string) float64 {
	createGraph(rates)

	traversed := map[string]bool{
		from: true,
	}
	stack := []*Node{nodeMap[from]}
	rateMap := map[string]float64{
		from: 1.0,
	}
	ans := 0.0
	for len(stack) != 0 {
		node := stack[0]
		stack = stack[1:]
		currentRate := rateMap[node.Name]
		for target, rate := range node.Edges {
			if target == to {
				ans = currentRate * rate
				break
			}
			if !traversed[target] {
				stack = append(stack, nodeMap[target])
				rateMap[target] = currentRate * rate
			}
		}
	}

	return ans
}

func createGraph(rates [][]interface{}) {
	for _, rate := range rates {
		from := rate[0].(string)
		to := rate[1].(string)
		rate := rate[2].(float64)
		fromNode, ok := nodeMap[from]
		if !ok {
			fromNode = &Node{Name: from, Edges: map[string]float64{}}
			nodeMap[from] = fromNode
		}
		toNode, ok := nodeMap[to]
		if !ok {
			toNode = &Node{Name: to, Edges: map[string]float64{}}
			nodeMap[to] = toNode
		}
		fromNode.Edges[to] = rate
		toNode.Edges[from] = 1 / rate
	}
}

func main() {
	rate1 := []interface{}{
		"USD",
		"JPY",
		110.0,
	}
	rate2 := []interface{}{
		"USD",
		"AUD",
		1.45,
	}
	rate3 := []interface{}{
		"JPY",
		"GBP",
		0.0070,
	}
	result := converse([][]interface{}{rate1, rate2, rate3}, "GBP", "AUD")
	fmt.Println(result)
}
