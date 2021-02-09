package main

import (
	"fmt"
)

func main () {
	var a int
	fmt.Scanf("%d",&a)
	for i := 0; i < a; i++ {
		var a,b int
		fmt.Scanf("%d %d",&a,&b)
		fmt.Println(a+b)
	}

}
	