package structural

import "fmt"

func DoBrigde(param []int) {
	// so the condition is stated before it used
	for _, v := range param {
		if v%4 == 0 {
			fmt.Println("a")
		}
	}
}

func Bridge() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	DoBrigde(a)

	b := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	DoBrigde(b)
}
