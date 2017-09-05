package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int
	var output []string
	fmt.Scanf("%d", &input)

	first_inputs := make([]int, input)
	second_inputs := make([]int, input)

	for i := range first_inputs {
		fmt.Scan(&first_inputs[i])
	}
	for i := range second_inputs {
		fmt.Scan(&second_inputs[i])
		result := second_inputs[i] + first_inputs[i]
		output = append(output, strconv.Itoa(result))
	}

	fmt.Println(strings.Join(output, " "))

}
