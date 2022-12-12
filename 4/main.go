package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	getpuzzle "github.com/Tiv4n/AoC2022"
)

func main() {
	var sector1, sector2 string
	var totalOverlap int
	_, input := getpuzzle.GetInput(4, os.Getenv("AOC"))

	for _, row := range input {
		if row != "" {
			sectors := strings.Split(row, ",")
			sector1 = sectors[0]
			sector2 = sectors[1]
			totalOverlap += findFullOverlap(sector1, sector2)
		}
	}
	fmt.Println("The total number of complete overlap are: ", totalOverlap)
	//The total number of complete overlap are:  450
	//PART 2

}

func findFullOverlap(first_range, second_range string) int {
	low1, _ := strconv.Atoi(strings.Split(first_range, "-")[0])
	low2, _ := strconv.Atoi(strings.Split(second_range, "-")[0])
	high1, _ := strconv.Atoi(strings.Split(first_range, "-")[1])
	high2, _ := strconv.Atoi(strings.Split(second_range, "-")[1])
	if (low1 >= low2 && high1 <= high2) || (low1 <= low2 && high1 >= high2) {
		return 1
	}
	//part 2
	if (low1 <= low2 && high1 >= low2) || (low1 <= high2 && high1 >= high2) || (low2 <= low1 && high2 >= low1) || (low2 <= high1 && high2 >= high1) {
		return 1
	}
	return 0
}
