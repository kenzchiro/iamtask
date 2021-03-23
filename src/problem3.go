package main

import (
	"fmt"
)

func swap(str string, i, j int) string {
      var result []byte
      for k := 0; k < len(str); k++ {
              if k == i {
                      result = append(result, str[j])
              } else if k == j {
                      result = append(result, str[i])
              } else {
                      result = append(result, str[k])
              }
      }
      return string(result)
}


func needSwap(str string, start int, curr int) bool {

	for i := start; i < curr; i++ {
		if str[i] == str[curr] {
			return false
		}
	}
	return true
}

func findPermutations(str string, index int, n int) {
	if index >= n {
		fmt.Println(" ", str)
		return
	}
	for i := index; i < n; i++ {
		check := needSwap(str, index, i) 			// check for distnict possible position
		if check {
		 	str = swap(str, index, i)				// swap character at index i with current character
			findPermutations(str, index+1, n)		// recursion for string [i+1:n]
			str = swap(str, index, i)				// backtrack (restore the string to its original state)
		}
	}
}
func main() {

	str := "112"
	n := len(str)
	findPermutations(str, 0, n)
}
