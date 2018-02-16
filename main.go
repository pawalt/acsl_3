package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	board [8][8]int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	hexLine, _ := strings.ToLower(reader.ReadString('\n'))
	hexas := strings.Split(hexLine, ",")
	hexa = strings.TrimSpace(hexa)
	i, err := strconv.ParseInt(hexa, 16, 64)
	handleErr(err)
	octal := strconv.FormatInt(i, 8)
	handleErr(err)
	fmt.Println(octal)
}

func handleErr(err error) {
	if err != nil {
		fmt.Println("Error:", err)
	}
}
