package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	getpuzzle "github.com/Tiv4n/AoC2022"
)

type stack struct {
	stack [][]string
	nmb   int
}

func (stk *stack) print() {
	var k int
	for j := len(stk.stack[k]) - 1; j >= 0; j-- {
		for i := 0; i < len(stk.stack); i++ {
			if j >= len(stk.stack[k]) {
				fmt.Printf("   ")
			} else {
				fmt.Printf("[%s]", stk.stack[i][j])
			}
			fmt.Printf(" ")
			k += 1
		}
		fmt.Println()
		k = 0

	}
	for i := 0; i < stk.nmb; i++ {
		fmt.Printf(" %d ", i+1)
		fmt.Printf(" ")
	}
	fmt.Println()

}

func (stk *stack) Create(input []string) {
	for id, line := range input {
		if id == 0 {
			line = strings.Trim(line, " ")
			sliced_line := strings.Split(line, "  ")
			stk.nmb = len(sliced_line)
			stk.stack = make([][]string, stk.nmb)
		} else {
			if string(line[0]) == " " {
				line = strings.Replace(line, "   ", "-", 1)
			}
			if strings.Contains(line, "     ") {
				line = strings.ReplaceAll(line, "     ", " - ")
			}
			if string(line[len(line)-1]) == " " {
				line = strings.Replace(line, "    ", " -", 1)
				line = strings.Replace(line, "-  ", "- ", 1)
				line = strings.Replace(line, "-   ", "- -", 1)
			}
			sliced_line := strings.Split(line, " ")
			fmt.Println(sliced_line)
			fmt.Println(len(sliced_line))
			for nmb, value := range sliced_line {
				if len(value) >= 3 {
					//fmt.Println(value)
					value = value[:len(value)-1]
					value = strings.Replace(value, "[", "", 1)
					value = strings.Replace(value, "]", "", 1)
					//fmt.Println(value)
					stk.stack[nmb] = append(stk.stack[nmb], value)
				}
			}

		}
	}

}
func (stk *stack) Move(quantity, from, to int) {
	//allign indexes
	from -= 1
	to -= 1
	if quantity > len(stk.stack[from]) {
		quantity = len(stk.stack[from])
	}
	//not invert the order
	if quantity != 0 {
		fmt.Println("Move: ", quantity, "BEFORE TO: ", stk.stack[to], "FROM: ", stk.stack[from])
		/*part 1
		var new_order []string
		old_order := stk.stack[from][len(stk.stack[from])-quantity:]
		for id, _ := range old_order {
			new_order = append(new_order, old_order[len(old_order)-1-id])
		}
		stk.stack[to] = append(stk.stack[to], new_order...)*/
		stk.stack[to] = append(stk.stack[to], stk.stack[from][len(stk.stack[from])-quantity:]...)
		stk.stack[from] = stk.stack[from][:len(stk.stack[from])-quantity]
		fmt.Println("AFTER TO:", stk.stack[to], "FROM: ", stk.stack[from])
	}
}

func main() {
	var stack stack
	var starting_part []string
	_, input := getpuzzle.GetInput(5, os.Getenv("AOC"))

	for id, row := range input {
		if row == "" && id < 20 {
			//fmt.Println(row)
			for i := 1; i <= id; i++ {
				starting_part = append(starting_part, input[id-i])
			}
			stack.Create(starting_part)
			fmt.Println(stack.stack)
		}
		if strings.Contains(row, "move") {
			words := strings.Split(row, " ")
			qty, _ := strconv.Atoi(words[1])
			from_id, _ := strconv.Atoi(words[3])
			to_id, _ := strconv.Atoi(words[5])
			if id == 20 {
				fmt.Println(qty, from_id, to_id)
			}
			stack.Move(qty, from_id, to_id)

		}
	}
	stack.print()
	//TLNGFGMFN
}
