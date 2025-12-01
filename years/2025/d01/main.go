package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s1, err := part1(string(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s1)

	s2, err := part2(string(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s2)
}

func part1(input string) (int, error) {

	cursor := 50
	pass := 0

	for line := range strings.Lines(string(input)) {
		dir, num := line[0], line[1:len(line)-1]

		n, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		if dir == 'L' {
			cursor -= n
		} else {
			cursor += n
		}

		cursor = cursor % 100

		if cursor < 0 {
			cursor += 100
		}

		if cursor == 0 {
			pass++
		}


	}
	return pass, nil
}

func part2(input string) (int, error) {
	zeroCount := 0
	pos := 50

	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		dir := line[0]
		numStr := line[1:]

		offset, err := strconv.Atoi(numStr)
		if err != nil {
			return 0, err
		}

		if dir == 'L' {
			offset = -offset
		}

		if offset >= 0 {
			zeroCount += (pos + offset) / 100
		} else {
			zeroCount -= offset / 100
			offset %= 100
			if pos != 0 && (pos+offset) <= 0 {
				zeroCount++
			}
		}
		pos = (pos + offset%100 + 100) % 100
	}

	return zeroCount, nil
}
