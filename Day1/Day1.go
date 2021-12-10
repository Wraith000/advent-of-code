package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var total, previous, x, y int
	queue := []int{}
	total, x = 0, 0
	previous = 0
	file, err := os.Open("Day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		push, err := strconv.Atoi(scanner.Text())
		queue = append(queue, push)
		x += push
		if len(queue) == 3 {
			if x > previous {
				total += 1
			}
			previous = x
			y, queue = queue[0], queue[1:]
			x -= y
		}
		if err != nil {
			// handle error
			fmt.Println(err)
			os.Exit(2)
		}
	}
	fmt.Println(total)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
