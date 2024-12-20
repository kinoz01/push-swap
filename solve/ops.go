package solver

import "fmt"


func sa(a []int) {
	a[0], a[1] = a[1], a[0]
	fmt.Println("sa")
}

func sb(b []int) {
	b[0], b[1] = b[1], b[0]
	fmt.Println("sb")
}

func ss(a, b []int) {
	sa(a)
	sb(b)
	fmt.Println("ss")
}

func rb(b []int) {
	for i := 0; i < len(b)-1; i++ {
		b[i], b[i+1] = b[i+1], b[i]
	}
	fmt.Println("rb")
}

func rrb(b []int) {
	for i := len(b) - 1; i > 0; i-- {
		b[i], b[i-1] = b[i-1], b[i]
	}
	fmt.Println("rrb")
}



func reverse(a []int) {
	for i := len(a) - 1; i > 0; i-- {
		a[i], a[i-1] = a[i-1], a[i]
	}
}
