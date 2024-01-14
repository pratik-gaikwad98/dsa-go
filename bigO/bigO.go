package bigo

import "fmt"

// Date: 2024-01-14 11:10:01.698728208 +0530 IST

func CalculateBigO() {
	LinearTime()

}

// O(n) -> Linear Time(Increases time according to iteration)
//  1. if We have to find element 30 loop will iterate till 3
//  2. if the loop size is 10 and 30 element is on 10 so the loop will iteterate 10 times
//     iteration depend on length of array thats y it is O(n) >> n times loop will iterate
func LinearTime() {
	myArray := []int{10, 20, 30}
	find30 := 30
	for location, value := range myArray {
		if find30 == value {
			fmt.Printf("Found on Location: %v & Value is : %d\n", location, value)
		}
	}
}
