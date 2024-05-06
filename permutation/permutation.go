package permutation

/*
Go言語による順列生成プログラム
2024-5-5 K.Komatsu
*/
type Permutation struct {
	Ketasu int
	Result []int
}

/*
引数で指定した桁数の数字の順列を生成してResultにセットする。ゼロは含まない。
例えば3桁の場合は、123,132,213,231,312,321の6個の数字を生成する。
*/
func (p *Permutation) Calc(ketasu int) {
	// 順列の最小値を開始値とする。
	value := p.init(ketasu)
	// 次に大きな順列の値を最大値になるまで繰り返し取得する。
	for ix := 0; value > 0; ix += 1 {
		p.Result[ix] = value
		value = p.nextVal(value)
	}
}

/*
引数の次に大きい順列値を返す。
これ以上無いときは0を返す。
*/
func (p *Permutation) nextVal(inValue int) int {
	// 入力値各桁の値を配列に逆順でセットする（例：12345なら5,4,3,2,1）
	array := make([]int, p.Ketasu)
	v := inValue
	for ix := 0; ix < p.Ketasu; ix += 1 {
		array[ix] = v % 10
		v /= 10
	}
	// array[ix-1]>array[ix]となるixを求める（但しix=1〜最終桁）
	v1 := 0
	for ix := 1; ix < p.Ketasu; ix += 1 {
		if array[ix-1] > array[ix] {
			v1 = ix
			break
		}
	}
	// 該当無ければ終了する
	if v1 == 0 {
		return 0
	}
	// array[ix]>array[v1]となるixを求める（但しix=0〜v1-1）
	v2 := 0
	for ix := 0; ix < v1; ix += 1 {
		if array[ix] > array[v1] {
			v2 = ix
			break
		}
	}
	// v1とv2の値を交換する（例：5,4,3,2,1 → 4,5,3,2,1）
	array[v1], array[v2] = array[v2], array[v1]
	// array[0]からarray[v1-1]を逆順にする
	if v1 > 1 {
		t := make([]int, v1)
		for ix := 0; ix < v1; ix += 1 {
			t[ix] = array[v1-ix-1]
		}
		for ix := 0; ix < v1; ix += 1 {
			array[ix] = t[ix]
		}
	}
	// 配列を整数に戻す（例：4,5,3,2,1 → 12354）
	outValue := 0
	for ix := p.Ketasu; ix > 0; ix -= 1 {
		outValue *= 10
		outValue += array[ix-1]
	}
	return outValue
}

/*
構造体を初期化し、指定した桁数の順列の最小値を返す（例：5桁なら12345）
*/
func (p *Permutation) init(ketasu int) int {
	// 桁数は1〜9の範囲に限る。
	if ketasu < 1 {
		ketasu = 1
	} else if ketasu > 9 {
		ketasu = 9
	}
	// 桁数を保存
	p.Ketasu = ketasu
	// 結果を格納するSliceの領域を確保する。
	size := 1
	for ix := 1; ix <= ketasu; ix += 1 {
		size *= ix
	}
	p.Result = make([]int, size)
	// 指定された桁数の最小値を求める。
	minVal := 0
	for ix := 1; ix <= ketasu; ix += 1 {
		minVal *= 10
		minVal += ix
	}
	return minVal
}
