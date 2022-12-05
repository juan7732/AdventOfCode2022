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

func score (opponent, player rune) int {
	// A & X == rock
	// B & Y == paper
	// C & Z == scissors
	// opponent e {A,B,C}
	// player e {X,Y,Z}
	// loss = 0 points, tie = 3 points, win = 6 points
	A := map[rune] int {
		'X': 3,
		'Y': 6,
		'Z': 0,
	}
	B := map[rune] int {
		'X': 0,
		'Y': 3,
		'Z': 6,
	}
	C := map[rune] int {
		'X': 6,
		'Y': 0,
		'Z': 3,
	}
	playScore := map[rune] int {
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	rules := map[rune] map[rune]int {
		'A': A,
		'B': B,
		'C': C,
	}

	totalScore := 0
	totalScore += rules[opponent][player]
	totalScore += playScore[player]
	return totalScore
}

func part1 () {
	filepath := "./day2input.txt"
	file, _ := os.Open(filepath)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	finalScore := 0

	for fileScanner.Scan() {
		runes := []rune(fileScanner.Text())
		finalScore += score(runes[0], runes[len(runes) - 1])
	}
	fmt.Println(finalScore)
}

func scoreVariant (opponent, player rune) int {
	// A = rock
	// B = paper
	// C = scissors
	// X = Lose
	// Y = Draw
	// Z = Win
	// opponent e {A,B,C}
	// player e {X,Y,Z}
	// loss = 0 points, tie = 3 points, win = 6 points
	A := map[rune] int {
		'X': 3,
		'Y': 1,
		'Z': 2,
	}
	B := map[rune] int {
		'X': 1,
		'Y': 2,
		'Z': 3,
	}
	C := map[rune] int {
		'X': 2,
		'Y': 3,
		'Z': 1,
	}
	playScore := map[rune] int {
		'X': 0,
		'Y': 3,
		'Z': 6,
	}
	rules := map[rune] map[rune]int {
		'A': A,
		'B': B,
		'C': C,
	}

	totalScore := 0
	totalScore += rules[opponent][player]
	totalScore += playScore[player]
	return totalScore
}


func part2 () {
	filepath := "./day2input.txt"
	file, _ := os.Open(filepath)

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	finalScore := 0

	for fileScanner.Scan() {
		runes := []rune(fileScanner.Text())
		finalScore += scoreVariant(runes[0], runes[len(runes) - 1])
	}
	fmt.Println(finalScore)
}