package solver

import (
	"fmt"
	"sort"
)

func SortStack(stackA []int) string {
	stackB := make([]int, 0, 100)
	switch len(stackA) {
	case 2:
		SortTwo(&stackA)
	case 3:
		SortThree(&stackA)
	default:
		
	}
	fmt.Println(stackB)
	return "hh"
}

func SortTwo(stackA *[]int) {
	a := *stackA
	if a[0] > a[1] {
		sa(a)
	}
}

func SortThree(stackA *[]int) {
	a := *stackA
	sorted := make([]int, 3)
	copy(sorted, a)
	sort.Ints(sorted)

	s, m, l := sorted[0], sorted[1], sorted[2]

	switch {
	case m == a[0] && s == a[1] && l == a[2]:
		fmt.Println("sa")
	case l == a[0] && m == a[1] && s == a[2]:
		fmt.Println("sa")
		fmt.Println("rra")
	case l == a[0] && s == a[1] && m == a[2]:
		fmt.Println("ra")
	case s == a[0] && l == a[1] && m == a[2]:
		fmt.Println("sa")
		fmt.Println("ra")
	case m == a[0] && l == a[1] && s == a[2]:
		fmt.Println("rra")
	}

	a[0], a[1], a[2] = s, m, l
}
