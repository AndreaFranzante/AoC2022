package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	getpuzzle "github.com/Tiv4n/AoC2022"
)

type Dir struct {
	parent  *Dir
	subdirs []*Dir
	name    string
	files   []int
}

func parseInput(lines []string) map[string]*Dir {
	fs := make(map[string]*Dir)
	currentDir := "/"
	rootDir := new(Dir)
	rootDir.name = "/"
	rootDir.parent = rootDir
	fs["/"] = rootDir

	for _, line := range lines {
		if line != "" {

			words := strings.Fields(line)
			switch words[0] {
			case "$":
				if words[1] == "cd" {
					switch words[2] {
					case "..":
						currentDir = fs[currentDir].parent.name
					case "/":
						currentDir = "/"
					default:
						currentDir += "/" + words[2]
					}
				}
			case "dir":
				dirName := fs[currentDir].name + "/" + words[1]
				_, ok := fs[dirName]
				if !ok {
					fs[dirName] = new(Dir)
					fs[dirName].name = dirName
					fs[dirName].parent = fs[currentDir]
				}
				fs[currentDir].subdirs = append(fs[currentDir].subdirs, fs[dirName])
			default:
				fileSize, _ := strconv.Atoi(words[0])
				fs[currentDir].files = append(fs[currentDir].files, fileSize)

			}

		}
	}
	return fs
}

func recurseDirSize(root *Dir) int {
	size := 0
	for _, dir := range root.subdirs {
		size += recurseDirSize(dir)
	}
	for _, file := range root.files {
		size += file
	}
	return size
}

func main() {
	_, inputLines := getpuzzle.GetInput(7, os.Getenv("AOC"))
	fs := parseInput(inputLines)

	firststep := 0
	total := 70000000
	needed := 30000000
	var sizes []int

	for key, _ := range fs {
		size := recurseDirSize(fs[key])
		sizes = append(sizes, size)
		if size <= 100000 {
			firststep += size
		}
	}

	toRemove := recurseDirSize(fs["/"]) - (total - needed)
	minValid := 99999999
	for _, size := range sizes {
		if size > toRemove && size < minValid {
			minValid = size
		}
	}

	fmt.Println("p1:", firststep)
	fmt.Println("p2:", minValid)
}
