package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToInt(s string) int {
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}

func firstDigit(line string) int {
	l := len(line)
	for i := 0; i < l; i++ {
		val, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return val * 10
		}

		if l-i >= 3 {
			val = convertToInt(line[i : i+3])
			if val > 0 {
				return val * 10
			}
		}
		if l-i >= 4 {
			val = convertToInt(line[i : i+4])
			if val > 0 {
				return val * 10
			}
		}
		if l-i >= 5 {
			val = convertToInt(line[i : i+5])
			if val > 0 {
				return val * 10
			}
		}
	}
	return 0
}

func secondDigit(line string) int {
	l := len(line)
	for i := l - 1; i >= 0; i-- {
		val, err := strconv.Atoi(string(line[i]))
		if err == nil {
			return val
		}

		if l-i >= 3 {
			val = convertToInt(line[i : i+3])
			if val > 0 {
				return val
			}
		}

		if l-i >= 4 {
			val = convertToInt(line[i : i+4])
			if val > 0 {
				return val
			}
		}

		if l-i >= 5 {
			val = convertToInt(line[i : i+5])
			if val > 0 {
				return val
			}
		}
	}
	return 0
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sum := 0

	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		first := firstDigit(line)
		sum += first
		second := secondDigit(line)
		sum += second
	}

	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
