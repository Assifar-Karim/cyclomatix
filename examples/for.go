package examples

import "fmt"

func for_fc1() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

func for_fc2() {
	i := 0
	for i < 2 {
		fmt.Println(i)
		i++
	}
}

func for_fc3() {
	for i := 0; i < 2; i++ {
		if i == 1 {
			break
		}
		fmt.Println(i)
	}
}

func for_fc4() {
	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

func for_fc5() string {
	for i := 0; i < 2; i++ {
		if i%2 == 1 {
			return "Bingo"
		}
		fmt.Println(i)
	}
	return "7:3"
}

func for_fc6() {
	arr := []int{1, 3, 5, 7, 11, 13}

	for _, v := range arr {
		fmt.Println(v)
	}
}
