package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	reportData := []string{}
	var bitCount int
	bitCount = 0
	var oxygenRating string
	var co2Rating string
	//region File
	file, err := os.Open("Day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reportData = append(reportData, scanner.Text())
	}
	//endregion File
	tmp := append([]string(nil), reportData...)
	for ok := true; ok; {
		OneCount := []string{}
		ZeroCount := []string{}
		for _, v := range tmp {
			x := v[bitCount : bitCount+1]
			if x == "1" {
				OneCount = append(OneCount, v)
			} else {
				ZeroCount = append(ZeroCount, v)
			}
		}

		if len(OneCount) > len(ZeroCount) {
			tmp = OneCount
		} else if len(ZeroCount) > len(OneCount) {
			tmp = ZeroCount
		} else {
			tmp = OneCount
		}
		bitCount++
		if len(tmp) == 1 {
			ok = false
			oxygenRating = tmp[0]
		}

	}
	bitCount = 0
	tmp = append([]string(nil), reportData...)
	for ok := true; ok; {
		OneCount := []string{}
		ZeroCount := []string{}
		for _, v := range tmp {
			x := v[bitCount : bitCount+1]
			if x == "1" {
				OneCount = append(OneCount, v)
			} else {
				ZeroCount = append(ZeroCount, v)
			}
		}

		if len(OneCount) < len(ZeroCount) {
			tmp = OneCount
		} else if len(ZeroCount) < len(OneCount) {
			tmp = ZeroCount
		} else {
			tmp = ZeroCount
		}
		bitCount++
		if len(tmp) == 1 {
			ok = false
			co2Rating = tmp[0]
		}

	}

	fmt.Println(strconv.ParseInt(oxygenRating, 2, 64))
	fmt.Println(strconv.ParseInt(co2Rating, 2, 64))
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
