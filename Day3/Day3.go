package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("Day3.txt")
	var count int
	count = 0
	values := [12]string{}
	var gammaBinary, epsilonBinary string

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var x []string
		x = strings.Split(scanner.Text(), "")
		for i, s := range x {
			values[i] += s
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		count += 1
	}
	for _, s := range values {
		if strings.Count(s, "1") > (count / 2) {
			gammaBinary += "1"
			epsilonBinary += "0"
		} else {
			gammaBinary += "0"
			epsilonBinary += "1"
		}
	}
	fmt.Println(gammaBinary)
	fmt.Println(epsilonBinary)
	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	binary, err := strconv.ParseInt(epsilonBinary, 2, 64)
	fmt.Println(gamma)
	fmt.Println(binary)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
