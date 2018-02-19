package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	board     [8][8]int
	direction int
)

func main() {
	b, _ := ioutil.ReadFile("c_3_walk_test.txt")
	strIn := strings.Split(string(b), "\n")
	hexLine := strIn[0]
	hexLine = strings.ToLower(hexLine)
	hexLine = strings.TrimSpace(hexLine)
	hexas := strings.Split(hexLine, " ")
	for i := 0; i < len(hexas); i++ {
		hexa := hexas[i]
		iConv, _ := strconv.ParseInt(hexa, 16, 64)
		octal := strconv.FormatInt(iConv, 8)
		for j := 0; j < 8; j++ {
			val, _ := strconv.Atoi(octal[j : j+1])
			board[i][j] = val
		}
	}

	for i := 0; i < 5; i++ {
		dirLine := strIn[i+1]
		dirLine = strings.TrimSpace(dirLine)
		items := strings.Split(dirLine, " ")
		row, _ := strconv.Atoi(items[0])
		row--
		col, _ := strconv.Atoi(items[1])
		col--
		dir := getDir(items[2])
		dir += board[row][col]
		dir %= 8
		steps, _ := strconv.Atoi(items[3])
		for j := 0; j < steps; j++ {
			rDel, cDel := posDelta(dir)
			row += rDel
			col += cDel
			dir += 4
			dir += board[row][col]
			dir %= 8
		}
		fmt.Println(strconv.Itoa(row+1) + " " + strconv.Itoa(col+1))
	}
}

func posDelta(dir int) (int, int) {
	if dir == 0 {
		return 0, 1
	}
	if dir == 1 {
		return -1, 1
	}
	if dir == 2 {
		return -1, 0
	}
	if dir == 3 {
		return -1, -1
	}
	if dir == 4 {
		return 0, -1
	}
	if dir == 5 {
		return 1, -1
	}
	if dir == 6 {
		return 1, 0
	}
	if dir == 7 {
		return 1, 1
	}
	return 0, 0
}

func getDir(dir string) int {
	if dir == "R" {
		return 0
	}
	if dir == "B" {
		return 2
	}
	if dir == "L" {
		return 4
	}
	if dir == "A" {
		return 6
	}

	return -1
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
