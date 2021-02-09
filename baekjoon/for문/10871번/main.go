package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &n, &x)
	defer writer.Flush()
    
	var seq = make([]int, n)
	for i := range seq {
		fmt.Fscanf(reader, "%d ", &seq[i])
		if seq[i] < x {
			fmt.Fprintf(writer, "%d ", seq[i])
		}
	}
	fmt.Fprint(writer, "\n")
}