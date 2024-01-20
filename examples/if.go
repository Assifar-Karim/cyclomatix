package examples

import "fmt"

func if_fc1() {
	v := 0
	if v >= 1 {
		y := 1
		fmt.Println(y)
	}
}

func if_fc2() {
	v := 0
	if v >= 1 {
		y := 1
		fmt.Println(y)
	} else {
		a := 77
		fmt.Println(a)
	}
}

func if_fc3() {
	v := 0
	if v >= 1 {
		y := 1
		fmt.Println(y)
	} else if v <= -2 {
		fmt.Println("Hello 1")
		z := 7
		fmt.Println(z)
	} else {
		a := 77
		fmt.Println(a)
	}
	fmt.Println("Hello 2")
}

func if_fc4(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func if_fc5(v int) {
	if i := if_fc4(v); i == 5 {
		fmt.Println("Bingo")
	}
}
