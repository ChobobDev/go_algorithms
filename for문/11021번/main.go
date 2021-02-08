package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	Writer := bufio.NewWriter(os.Stdout)

	var N int
	if Scanner.Scan() {
		N, _ = strconv.Atoi(Scanner.Text())
	}

	for i := 0; i < N; i++ {
		if Scanner.Scan() {
			a := strings.Fields(Scanner.Text())
			x, _ := strconv.Atoi(a[0])
			y, _ := strconv.Atoi(a[1])
			Writer.WriteString("Case #")
			Writer.WriteString(strconv.Itoa(i + 1))
			Writer.WriteString(": ")
			Writer.WriteString(strconv.Itoa(x + y))
			Writer.WriteString("\n")
		}
	}
	Writer.Flush()
}