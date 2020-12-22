package main

import (
	"fmt"
	"unsafe"
)

const (
	bucketCntBits = 3
	bucketCnt     = 1 << bucketCntBits
)

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap

	nextOverflow *bmap
}

// A bucket for a Go map.
type bmap struct {
	tophash [bucketCnt]uint8
}

type T struct {
	Number int
}

func main() {
	s := []T{{1}, {2}, {3}, {4}, {5}}
	s2 := []*T{}
	for _, v := range s {
		fmt.Printf("pointer: %p, value: %v\n", &v, v)
		s2 = append(s2, &v)
	}
	for _, v := range s2 {
		fmt.Printf("%+v¥n\n", v)
		// &{Number:5}
	}

	tryStructLayout()

	str := "hello"
	// これは `b = []byte(str)` よりもパフォーマンスが良い
	// なぜならstrを別の場所にコピーすることなく直接pointerを参照してbyteに変換しているから。
	// これができる理由はメモリのレイアウトにあって、String とSliceのデータ構造を見ると Data uintptrとLen intが同じなので、レイアウトをとってきて、それをsliceにそのままキャストするという技術を使っている
	_ = *(*[]byte)(unsafe.Pointer(&str))

}

// String stringのレイアウト, 終端の文字列は保持しずに、開始のpointerと長さを保持している
type String struct {
	Data uintptr // 文字列データの先頭アドレス
	Len  int     // 文字列の長さ
}

// Slice sliceのレイアウト
type Slice struct {
	Data uintptr // スライス要素が連続して並んでいる領域の先頭アドレス
	Len  int     // スライスの長さ
	Cap  int     // スライスの容量
}

// hmap mapのレイアウト
type hmap struct {
	count      int
	flags      uint8
	B          uint8
	noverflow  uint16
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr
	extra      *mapextra
}

// T2 structのレイアウトはそのまま。最初のFieldとstruct自体のメモリアドレスは等しくなる
// 64bitのPCだと、intは8byteとして計算される
type T2 struct {
	A int
	B string
	C []int
	D int
}

func tryStructLayout() {
	var v T2
	fmt.Printf("v address = %p\n", &v)
	// v address = 0xc00010c040
	fmt.Printf("v.A address = %p\n", &v.A)
	// v.A address = 0xc00010c040
	fmt.Printf("v.B address = %p\n", &v.B)
	// v.B address = 0xc00010c048
	fmt.Printf("v.C address = %p\n", &v.C)
	// v.C address = 0xc00010c058
	fmt.Printf("A size = %d\n", uintptr(unsafe.Pointer(&v.B))-uintptr(unsafe.Pointer(&v.A)))
	// A size = 8
	fmt.Printf("B size = %d\n", uintptr(unsafe.Pointer(&v.C))-uintptr(unsafe.Pointer(&v.B)))
	// B size = 16
	fmt.Printf("C size = %d\n", uintptr(unsafe.Pointer(&v.D))-uintptr(unsafe.Pointer(&v.C)))
	// C size = 24

	var a struct{}
	b := struct{}{}
	fmt.Printf("pointer: %p, bool: %t\n", a, a)
	fmt.Printf("pointer: %p, bool: %t\n", b, b)
}

// emptyInterface interfaceのレイアウト
// プログラミング言語処理系のボクシング後の構造と等しい
// ランタイム時に変数の型情報を動的に取得するために利用するreflectパッケージでは、このtypをの値を reflect.Typeとして解釈して使っている
type emptyInterface struct {
	typ unsafe.Pointer // 型情報へのポインタ
	ptr unsafe.Pointer // 値へのポインタ
}
