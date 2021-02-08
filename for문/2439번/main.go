package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &n)
	for i := 1; i <= n; i++ {
		for k:=1;k<=n-i; k++{
			fmt.Print(" ")
		}
		for j:=1;j<=i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}
}