package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	sum := 0

	fmt.Println("Hello")
	f, err := os.Open("input.txt")
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num := 0

		l := len(line)
		for i := 0; i < l; i++ {
			val, err := strconv.Atoi(string(line[i]))
			if err != nil {
				continue
			}
			num += val * 10
			break
		}
		for i := l - 1; i >= 0; i-- {
			fmt.Sprintln(line[i])
			val, err := strconv.Atoi(string(line[i]))
			if err != nil {
				continue
			}
			num += val
			break
		}
		sum += num
	}

	fmt.Println(fmt.Sprintf(" =  %d", sum))
}
