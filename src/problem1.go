package main

import (
	"fmt"
	"strings"
)

func isPalindrome(input string) bool {
word:= strings.ToLower(input)
l:=0
r:=len(input)-1

for l < r {
   if word[l] != word[r]{
     return false
   }
l++
r--
}
return true
}

func main() {
	input := "Deleveled"
	fmt.Println(input,isPalindrome(input))
}
