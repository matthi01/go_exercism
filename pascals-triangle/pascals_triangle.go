package pascal

import "fmt"

func Triangle(n int) [][]int {
	fmt.Println("####################")
	fmt.Println(n)
	for tail := 0; tail < n; tail++ {
		for head := 0; head < n+1; head++ {
			fmt.Println("head:", head)
			fmt.Println("tail:", tail)
		}
	}
	return [][]int{{1}}
}
