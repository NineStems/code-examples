package main

import (
	"fmt"
	"math/big"
)

func main(){
	var x, y, z big.Int
	fmt.Scan(&x,&y) 
	fmt.Println(z.Add(&x,&y))
}