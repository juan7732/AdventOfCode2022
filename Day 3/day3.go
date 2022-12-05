package main

import(
	"fmt"
	"bufio"
	"os"
)

func main () {
	part1()
	part2()
}

func part1 () {
	filePath := "./day3input.txt"
	file, _ := os.Open(filePath)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0

	for fileScanner.Scan(){
		totalPriority += convertPriorityRuneToValue(findPriorityRune(fileScanner.Text()))
	}

	fmt.Println(totalPriority)
}

func part2 (){
	filePath := "./day3input.txt"
	file, _ := os.Open(filePath)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	totalPriority := 0
	loopCount := 0
	var prevSet set 

	for fileScanner.Scan(){
		if loopCount == 2 {
			currentSet := StringToSet(fileScanner.Text())
			intersect := currentSet.Intersect(&prevSet)
			for k, _ := range intersect.m {
				totalPriority += convertPriorityRuneToValue([]rune(k)[0])
			}
		} else if loopCount == 1 {
			currentSet := StringToSet(fileScanner.Text())
			prevSet = currentSet.Intersect(&prevSet)
		} else if loopCount == 0 {
			prevSet = StringToSet(fileScanner.Text())
		}

		loopCount = (loopCount + 1) % 3
	}

	fmt.Println(totalPriority)
}

func findPriorityRune (rucksack string) rune{
	rucksackCompartmentOne := rucksack[0:len(rucksack)/2]
	rucksackCompartmentTwo := rucksack[len(rucksack)/2:len(rucksack)]
	ruckSetOne := StringToSet(rucksackCompartmentOne)
	ruckSetTwo := StringToSet(rucksackCompartmentTwo)

	priorityRune := ruckSetOne.Intersect(&ruckSetTwo)
	for k, _ := range priorityRune.m {
		return []rune(k)[0]
	}
	return '*'
}

func convertPriorityRuneToValue (priorityRune rune) int {
	if priorityRune > 96 {
		return int(priorityRune) - 96
	} else {
		return int(priorityRune) - 38
	}
}

// since there are no sets in go, let's make one
// thanks to https://www.davidkaya.com/sets-in-golang/#:~:text=Set%20is%20an%20abstract%20data%20type%20that%20can,are%20ways%20of%20imitating%20a%20Set%20in%20Go.

var exists = struct{}{}

type set struct {
	m map[string] struct{}
}

func NewSet() *set {
	s := &set{}
	s.m = make(map[string]struct{})
	return s
}

func (s *set) Add(value string) {
	s.m[value] = exists
}

func (s *set) Remove(value string){
	delete(s.m, value)
}

func (s *set) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *set) Intersect(value *set) set {
	n := NewSet()

	for k, _ := range s.m {
		if value.Contains(k){
			n.Add(k)
		}
	}
	return *n
}

func StringToSet(value string) set {
	n := NewSet()
	chars := []rune(value)
	for i:= range(chars) {
		n.Add(string(chars[i]))
	}
	return *n
}