package main

import (
	"fmt"
)

func recurSum(N int) int{
if N==1{
	return 1
}

//func Pow(base,exponent)
exponent := N
res:= 1  
for exponent != 0 {  
  res *= N
  exponent -= 1  
}

return res + recurSum(N-1) 
}

func main() {
	input := 4
	fmt.Println(recurSum(input))
}
