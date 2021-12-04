package part1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "math/bits"
)

func Run() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	total_lines := 0
	ones := make([]int, 12)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total_lines++
		line := scanner.Text()
		for i, c := range line {
			if c == '1' {
				ones[i]++
			}
		}
	}

	fmt.Println(ones)

	max := ""
	min := ""
	for _, count := range ones {
		if total_lines/2 < count {
			max += "1"
			min += "0"
		} else {
			max += "0"
			min += "1"
		}
	}

	fmt.Printf("max: %s, min: %s\n", max, min)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	_max, err := strconv.ParseUint(max, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	_min, err := strconv.ParseUint(min, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("max: %v, min: %v\n", _max, _min)
	fmt.Printf("%v\n", _max*_min)
}

// 10011
