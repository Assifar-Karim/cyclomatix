package examples

import "fmt"

func switch_fc1(idx int) string {
	switch idx {
	case 0:
		return "Monday"
	case 1:
		return "Tuesday"
	case 2:
		return "Wednesday"
	case 3:
		return "Thursday"
	case 4:
		return "Friday"
	case 5:
		return "Saturday"
	case 6:
		return "Sunday"
	default:
		return "Weirday"
	}
}

func switch_fc2(idx int) {
	switch day := switch_fc1(idx); day {
	case "Monday", "Tuesday", "Wednesday":
		fmt.Println("First half of the week")
	case "Thursday", "Friday":
		fmt.Println("Second half of the week")
	case "Saturday", "Sunday":
		fmt.Println("WEEK END")
	default:
		fmt.Println("ERROR! ERROR! MACHINE EXPLODING!!!!")
	}
}
