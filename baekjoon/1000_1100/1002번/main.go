package main

import (
	"fmt"
	"math"
)
func calc() {
	var x1,y1,r1,x2,y2,r2 int64
	fmt.Scanf("%d %d %d %d %d %d",&x1,&y1,&r1,&x2,&y2,&r2)
	var distance,rdif,rsum float64
	distance = math.Sqrt(math.Pow(float64(x1-x2),2) + math.Pow(float64(y1-y2),2))
	rsum=float64(r1+r2)
	rdif=math.Abs(float64(r1-r2))
	if(distance==0){
		if(r1==r2){
			fmt.Println(-1)
		}else{
			fmt.Println(0)
		}
		
	}else{
		if(distance == rsum || distance==rdif){
			fmt.Println(1)
		}else if(distance<rsum && distance>rdif){
			fmt.Println(2)
		}else{
			fmt.Println(0)
		}
	}

}

func main () {
	var a int
	fmt.Scanf("%d",&a)
	for i := 0; i < a; i++ {
		calc()	
	}
}
	