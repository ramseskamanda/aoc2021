package part2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "math/bits"
)

const (
	FILE     = "input.txt"
	NUM_BITS = 12
)

func Run() {
	file, err := os.Open(FILE)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	oxygen := find_oxygen(lines, 0)

	co2 := find_co2(lines, 0)

	fmt.Printf("%s (should be 10111)\n", oxygen)
	fmt.Printf("%s (should be 01010)\n", co2)

	_max, err := strconv.ParseUint(oxygen, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	_min, err := strconv.ParseUint(co2, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("oxygen: %v (should be 23), co2: %v (should be 10)\n", _max, _min)
	fmt.Printf("%v\n", _max*_min)
}

func find_oxygen(lines []string, index int) string {
	if len(lines) == 0 {
		return "encountered empty lines"
	}

	if len(lines) == 1 || index >= NUM_BITS {
		return lines[0]
	}

	count := 0
	for _, line := range lines {
		if line[index] == '1' {
			count++
		}
	}

	char_to_keep := '1'
	if float64(count) < float64(len(lines))/float64(2) {
		char_to_keep = '0'
	}

	fmt.Println("size:", len(lines))
	fmt.Println("1s count:", count)
	fmt.Println("midpoint:", float64(len(lines))/float64(2))
	fmt.Println("char_to_keep:", string(char_to_keep))
	fmt.Println("___________")

	new_lines := []string{}
	for _, line := range lines {
		if line[index] == byte(char_to_keep) {
			new_lines = append(new_lines, line)
		}
	}

	return find_oxygen(new_lines, index+1)
}

func find_co2(lines []string, index int) string {
	if len(lines) == 0 {
		return "encountered empty lines"
	}

	if len(lines) == 1 || index >= NUM_BITS {
		return lines[0]
	}

	count := 0
	for _, line := range lines {
		if line[index] == '0' {
			count++
		}
	}

	char_to_keep := '0'
	if count > len(lines)/2 {
		char_to_keep = '1'
	}

	new_lines := []string{}
	for _, line := range lines {
		if line[index] == byte(char_to_keep) {
			new_lines = append(new_lines, line)
		}
	}

	return find_co2(new_lines, index+1)
}
