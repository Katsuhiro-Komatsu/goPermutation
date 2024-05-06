package main

import (
	"fmt"
	"goPerm/permutation"
)

//import permutation

/*
go言語による順列生成
*/
func main() {
	const ketasu = 9
	perm := new(permutation.Permutation)
	perm.Calc(ketasu)
	for ix := 0; ix < len(perm.Result); ix += 1 {
		fmt.Println(perm.Result[ix])
	}
	fmt.Printf("%d桁の順列は%d個あります。", perm.Ketasu, len(perm.Result))
}
