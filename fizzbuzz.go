package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(i int) {
	if i%3 == 0 && i%5 == 0 {
		println("FIZZBUZZ")
	} else if i%3 == 0 {
		println("FIZZ")
	} else if i%5 == 0 {
		println("BUZZ")
	} else {
		println(i)
	}
}

func main() {
	fmt.Println("FizzBuzz")
	fmt.Println("First test from 1 to ")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a number")
	text, _ := reader.ReadString('\n')
	i, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		println("Not a number")
	}

	count := 1
	for count <= i {
		check(count)
		count += 1
	}

	for {
		fmt.Println("Enter a number")
		text, _ := reader.ReadString('\n')
		i, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			println("Not a number")
		} else {
			check(i)
		}
	}
}
