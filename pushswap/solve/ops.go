package solver

import "fmt"

func pa(a, b *[]int) {
	*a = append(*a, (*b)[0])
	reverse(*a)
	*b = (*b)[1:]
	fmt.Println("pa")
}

func pb(a, b *[]int) {
	*b = append(*b, (*a)[0])
	reverse(*b)
	*a = (*a)[1:]
	fmt.Println("pb")
}

func sa(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("sa")
}

func ra(a []int) {
	for i := 0; i < len(a)-1; i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}
	fmt.Println("ra")
}

func rra(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
	fmt.Println("rra")
}

func reverse(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}
