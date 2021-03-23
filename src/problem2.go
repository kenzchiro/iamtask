package main

import (
	"fmt"
)

func groupWords(strs []string) [][]string {
	result := [][]string{}
	m := make(map[rune][]string)

	for c, v := range strs {
		var a, d rune
		for _, i := range v {
			a ^= (i * i)
			d += (i * i)
		}
		hash := a ^ d*d
		m[hash] = append(m[hash], strs[c])
	}

	for _, v := range m {
		result = append(result, v)
	}
	return result
}

func main() {
	words := []string{"VMRCO", "VORCM", "MCRTOX", "ZXTAC", "XZATC", "XMTCOR", "OXVS", "AZTXC", "VXOS", "RZAT", "MRCOTX", "SVXO", "TRAZ", "ZTAR", "MVOCR"}
	results := groupWords(words)

	//for show result
	for i, res := range results {
		fmt.Printf("%d.", i+1)
		for j, ch := range res {
			fmt.Printf("%s", ch)
			if j < (len(res) - 1) {
				fmt.Printf("-")
			}
		}
		fmt.Println("")
	}
}
