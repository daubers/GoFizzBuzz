package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("FizzBuzz")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Enter a number")
		text, _ := reader.ReadString('\n')
		i, err := strconv.Atoi(strings.TrimSpace(text))
		if err != nil {
			println("Not a number")
		} else if i%3 == 0 && i%5 == 0 {
			println("FIZZBUZZ")
		} else if i%3 == 0 {
			println("FIZZ")
		} else if i%5 == 0 {
			println("BUZZ")
		} else {
			println(i)
		}
	}
}
