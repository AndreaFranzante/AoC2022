package main

import (
	"os"

	getpuzzle "github.com/Tiv4n/AoC2022"

	"fmt"
	"strconv"
)

/*
The jungle must be too overgrown and difficult to navigate in vehicles or access from the air; the Elves' expedition traditionally goes on foot. As your boats approach land, the Elves begin taking inventory of their supplies. One important consideration is food - in particular, the number of Calories each Elf is carrying (your puzzle input).

The Elves take turns writing down the number of Calories contained by the various meals, snacks, rations, etc. that they've brought with them, one item per line. Each Elf separates their own inventory from the previous Elf's inventory (if any) by a blank line.

For example, suppose the Elves finish writing their items' Calories and end up with the following list:

1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
This list represents the Calories of the food carried by five Elves:

The first Elf is carrying food with 1000, 2000, and 3000 Calories, a total of 6000 Calories.
The second Elf is carrying one food item with 4000 Calories.
The third Elf is carrying food with 5000 and 6000 Calories, a total of 11000 Calories.
The fourth Elf is carrying food with 7000, 8000, and 9000 Calories, a total of 24000 Calories.
The fifth Elf is carrying one food item with 10000 Calories.
In case the Elves get hungry and need extra snacks, they need to know which Elf to ask: they'd like to know how many Calories are being carried by the Elf carrying the most Calories. In the example above, this is 24000 (carried by the fourth Elf).

Find the Elf carrying the most Calories. How many total Calories is that Elf carrying?
*/
// funziona per la prima soluzione, trovare il valore di calorie massimo dell'elfo
func compute_max(input []int) (int, int) {
	var max int
	var ind int
	for index, i := range input {
		if i > max {
			max = i
			ind = index
		} else {
			continue
		}
	}
	return ind, max
}

func compute_sum_of_biggest_N(input []int, N int) int {
	var result int
	for i := 0; i < N; i++ {
		idx, max := compute_max(input)
		result += max
		input = append(input[:idx], input[idx+1:]...)
	}
	return result

}

func main() {
	N := 3
	var elf_calories []int
	_, input := getpuzzle.GetInput(1, os.Getenv("AOC"))
	var singleElfCalories int
	for _, val := range input {
		if val == "" {
			elf_calories = append(elf_calories, singleElfCalories)
			singleElfCalories = 0
		} else {
			supp, _ := strconv.Atoi(val)
			singleElfCalories += supp

		}
	}
	index, value := compute_max(elf_calories)
	fmt.Println("The elf with the most calories ", index, " is: with: ", value)
	top_three_calories_sum := compute_sum_of_biggest_N(elf_calories, N)
	fmt.Println("The sum of top ", N, " most calories elf is: ", top_three_calories_sum)

}

/*
The elf with the most calories  6  is: with:  67633
The sum of top  3  most calories elf is:  199628
*/
