package main

import (
	"fmt"
	"bufio"
    "os"
	"strconv"
)

func main () {
	part1()
	part2()
}

func part1() {
	filePath := "./day1input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	maxSize, tmp := 0, 0

	for fileScanner.Scan(){
		if fileScanner.Text() == "" {
			if tmp > maxSize {
				maxSize = tmp
			}
			tmp = 0
		} else {
			val, _ := strconv.Atoi(fileScanner.Text())
			tmp += val
		}
	}
	fmt.Println(maxSize)
	if err = file.Close(); err != nil {
		fmt.Printf("Could not close the file due to this %s error \n", err)
	}
}

func part2() {
	filePath := "./day1input.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("Could not open the file due to this %s error \n", err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	smallest, medium, largest, tmp := 0,0,0,0

	for fileScanner.Scan(){
		if fileScanner.Text() == "" {
			if tmp > smallest {
				smallest = tmp
			}
			if smallest > medium {
				medium, smallest = smallest, medium
			}
			if medium > largest {
				medium, largest = largest, medium
			}
			tmp = 0
		} else {
			val, _ := strconv.Atoi(fileScanner.Text())
			tmp += val
		}
	}

	fmt.Println(smallest + medium + largest)
}