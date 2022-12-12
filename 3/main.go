package main

import (
	"os"

	getpuzzle "github.com/Tiv4n/AoC2022"

	"fmt"
)

func ComputePoints(row string) []int {
	var points []int
	for _, char := range row {
		if char <= 90 {
			points = append(points, int(char-'A')+27)
		} else if char >= 97 {
			points = append(points, int(char-'a')+1)
		}
	}
	return points

}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func CompareCompartments(first []int, second []int) int {
	var output int
	var done []int
	for _, first_item := range first {
		for _, second_item := range second {
			if second_item == first_item {
				done = append(done, second_item)

			}
		}

	}
	done = removeDuplicate[int](done)
	//fmt.Println(done)
	for _, d := range done {
		output += d
	}
	return output
}

func main() {
	var sum_common_items int
	//var firstCompartments []int
	//var secondCompartments []int
	var points []int
	var threesets [][]int
	_, input := getpuzzle.GetInput(3, os.Getenv("AOC"))
	for id, row := range input {
		if row != "" {
			points = ComputePoints(row)

			//firstCompartments = points[:len(points)/2]
			//secondCompartments = points[len(points)/2:]
			//sum_common_items += CompareCompartments(firstCompartments, secondCompartments)

			//part 2
			threesets = append(threesets, points)

			if (id+1)%3 == 0 {
				sum_common_items += FindCommonBadges(threesets)
				threesets = nil
			}

		}
	}

	fmt.Println("The sum of priorities of each element in both compartments is: ", sum_common_items)
	//getpuzzle.SendAnswer(3, 1, string(sum_common_items))
}

func FindCommonBadges(sets [][]int) int {
	var output int
	var done []int
	for _, first_item := range sets[0] {
		for _, second_item := range sets[1] {
			for _, third_item := range sets[2] {
				if second_item == first_item && second_item == third_item {
					done = append(done, second_item)

				}
			}
		}

	}
	done = removeDuplicate[int](done)
	for _, d := range done {
		output += d
	}
	return output
}
