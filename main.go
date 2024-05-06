package main

import (
	"flag"
	"fmt"
	"goPerm/permutation"
)

func init() {

}

/*
go言語による順列生成

	使い方：
	goPerm -n 6		<== 6桁の順列を生成する。
*/
func main() {
	// 桁数を引数から取得する（指定なければ5桁）
	var ketasu = flag.Int("n", 5, "順列の桁数")
	flag.Parse()
	// 順列を求める
	perm := new(permutation.Permutation)
	perm.Calc(*ketasu)
	// 結果表示
	for ix := 0; ix < len(perm.Result); ix += 1 {
		fmt.Println(perm.Result[ix])
	}
	fmt.Printf("%d桁の順列は%d個あります。", perm.Ketasu, len(perm.Result))
}
