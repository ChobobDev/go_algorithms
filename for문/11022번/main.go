package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var t, a, b int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &t)
	defer writer.Flush()

	for i := 1; i <= t; i++ {
		fmt.Fscanln(reader, &a, &b)
		fmt.Fprintf(writer, "Case #%d: %d + %d = %d\n", i, a, b, a+b)
	}
}