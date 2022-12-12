package main

import (
	"fmt"
	"os"
	"strings"

	getpuzzle "github.com/Tiv4n/AoC2022"
)

func findMarker(lenght int, input string) {
	var marker string
	for id, char := range input {
		if idx := strings.Index(marker, string(char)); id == -1 {
			marker += string(char)
		} else {

			marker = marker[idx+1:]
			marker += string(char)
		}
		if len(marker) == lenght {
			fmt.Println("the packet marker is: ", marker, "\nThe last item is taken at: ", id+1)
			//1: The last item is taken at:  1658
			//2: The last item is taken at:  2260
			break
		}
	}

}

const msg_size = 14
const pkt_size = 4

func main() {
	datastream, _ := getpuzzle.GetInput(6, os.Getenv("AOC"))
	findMarker(pkt_size, datastream)
	findMarker(msg_size, datastream)
}
