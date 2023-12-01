package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToInt(s string) int {
	names := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	return names[s]
}

func retrieveDigit(i int, line string) int {
	l := len(line)
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
	return 0
}

func firstDigit(line string) int {
	l := len(line)
	for i := 0; i < l; i++ {
		val := retrieveDigit(i, line)
		if val > 0 {
			return val * 10
		}
	}
	return 0
}

func secondDigit(line string) int {
	l := len(line)
	for i := l - 1; i >= 0; i-- {
		val := retrieveDigit(i, line)
		if val > 0 {
			return val
		}
	}
	return 0
}

func main() {
	sum := 0

	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += firstDigit(line)
		sum += secondDigit(line)
	}

	fmt.Println(fmt.Sprintf("Sum =  %d", sum))
}
