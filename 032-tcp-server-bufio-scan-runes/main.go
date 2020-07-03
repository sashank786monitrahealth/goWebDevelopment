package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	var s string = "I felt so good like anything was possible\n I hit cruise control "
	var scanner *bufio.Scanner = bufio.NewScanner(strings.NewReader(s))

	//scanner.Split(bufio.ScanWords)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		var line string = scanner.Text()
		fmt.Println(line)
	}
}
