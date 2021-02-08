package main

import (
	"fmt"
)

func main () {
	var a,b int
	fmt.Scanf("%d\n%d",&a,&b)
	if a>0{
		if b>0{
			fmt.Print(1)
		} else{
			fmt.Print(4)
		}
	} else{
		if b>0{
			fmt.Print(2)
		} else{
			fmt.Print(3)
		}
	}

}
	