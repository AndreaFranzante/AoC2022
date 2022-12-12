package main

import (
	"os"

	getpuzzle "github.com/Tiv4n/AoC2022"

	"fmt"
	"strings"
)

const win = 6
const draw = 3
const lose = 0
const paper = 2
const rock = 1
const scissors = 3

// A rock B paper C scissors, X rock, Y paper, Z scissors
func rockpaperscissor(opponent, mychoice string) int {
	if (opponent == "A" && mychoice == "Y") || (opponent == "B" && mychoice == "Z") || (opponent == "C" && mychoice == "X") {
		return win
	} else if (opponent == "A" && mychoice == "X") || (opponent == "C" && mychoice == "Z") || (opponent == "B" && mychoice == "Y") {
		return draw
	} else {
		return lose
	}

}

func compute_strategy(input []string) int {
	var result int
	for _, strategy := range input {
		letters := strings.Split(strategy, " ")
		if len(letters) > 1 {
			switch letters[1] {
			case "X":
				result += 1 + rockpaperscissor(letters[0], letters[1])
			case "Y":
				result += 2 + rockpaperscissor(letters[0], letters[1])
			case "Z":
				result += 3 + rockpaperscissor(letters[0], letters[1])
			}
		}

	}
	return result
}

func main() {
	_, input := getpuzzle.GetInput(2, os.Getenv("AOC"))
	res := compute_strategy(input)
	fmt.Println("the total score for this strategy is: ", res)
	//getpuzzle.SendAnswer(2, 1, string(res)) not work
	res = compute_strategy2(input)
	fmt.Println("the total score for this 2nd strategy is: ", res)
}

// part 2: X lose, Y draw, Z win

func compute_strategy2(input []string) int {
	var result int
	for _, strategy := range input {
		letters := strings.Split(strategy, " ")
		if len(letters) > 1 {
			switch letters[1] {
			case "X":
				switch letters[0] {
				case "A":
					result += lose + scissors
				case "B":
					result += lose + rock
				case "C":
					result += lose + paper
				}
			case "Y":
				switch letters[0] {
				case "A":
					result += draw + rock
				case "B":
					result += draw + paper
				case "C":
					result += draw + scissors
				}
			case "Z":
				switch letters[0] {
				case "A":
					result += win + paper
				case "B":
					result += win + scissors
				case "C":
					result += win + rock
				}

			}
		}

	}
	return result
}
