package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 7, 11, 15};

	target := 18;

	seen := make(map[int]int);

	for i, num := range nums {
		complement := target - num;

		if index, ok := seen[complement]; ok {
			fmt.Println([]int{index, i});

			return;
		}

		seen[num] = i;
	}

	fmt.Println("No two sum solutions found");
}