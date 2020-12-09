package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	// part1()
	found, a, b, c := part2()
	if found {
		fmt.Println("found it!")
		fmt.Println(a * b * c)
	}
}
func part1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(number))
	}
	// fmt.Println(numbers)
	// basic algorithm: subtract first number from 2020, loop over rest until find a match, then break
	var i, j, n1, n2, test int
	var found bool
	for i = 0; i < len(numbers)-1; i++ {
		if found {
			break
		}
		n1 = numbers[i]
		test = 2020 - numbers[i]
		for j = i + 1; j < len(numbers); j++ {
			n2 = numbers[j]
			fmt.Println(n1, test, n2, test-n2)

			if test-n2 == 0 {
				found = true
				break
			}
		}
	}
	fmt.Println(n1, n2, n1*n2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func part2() (bool, int, int, int) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, int(number))
	}
	// fmt.Println(numbers)
	// basic algorithm: sort array, then start from both ends and work towards middle
	var i, a, b, c, start, end int
	var found bool

	sort.Ints(numbers)
	for i = 0; i < len(numbers)-1; i++ {
		a = numbers[i]
		start = i + 1
		end = len(numbers) - 1
		for start < end {
			b = numbers[start]
			c = numbers[end]
			fmt.Println(i, start, end, a, b, c)
			if a+b+c == 2020 {
				found = true
				return found, a, b, c
			} else if a+b+c > 2020 {
				end = end - 1
			} else {
				start = start + 1
			}
		}
	}
	// for i = 0; i < len(numbers)-1; i++ {
	// 	if found {
	// 		break
	// 	}
	// 	n1 = numbers[i]
	// 	test = 2020 - numbers[i]
	// 	for j = i + 1; j < len(numbers); j++ {
	// 		n2 = numbers[j]
	// 		fmt.Println(n1, test, n2, test-n2)

	// 		if test-n2 == 0 {
	// 			found = true
	// 			break
	// 		}
	// 	}
	// }
	// fmt.Println(n1, n2, n1*n2)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return false, 0, 0, 0
}
