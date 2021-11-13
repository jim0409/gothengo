package main

import (
	"fmt"
	"reflect"
)

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

e.g.
	2 -> II
	12 -> XII
	27 -> XXVII

1. Roman numerals are usually written largest to smallest from left to right.
2. However, the numeral for four is not IIII. Instead, the number four is written as IV
    ,Because the one is before the five we subtract it making four

I can be placed before V(5) and X(10) to make 4 and 9
X can be placed before L(50) and C(100) to make 40 and 90
C can be placed before D(500) and M(1000) to make 400 and 900

e.g.
	III -> 3
	IV -> 4
	IX -> 9
	LVIII -> 58 (L=50, V=5, III=3)
	MCMXCIV -> 1994 (M=1000, CM=900, XC=90, IV=4)

*/
// 這題要考驗programer對於 golang 中運行 byte的概念

func romanToInt(s string) int {
	res := 0
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		temp := m[s[i]]

		sign := 1
		if temp < last {
			sign = -1
		}
		res += sign * temp
		last = temp
	}

	return res
}

func main() {
	a := "I"
	b := 'I'

	fmt.Printf("The value %v has type %v\n", a, reflect.TypeOf(a)) // I -> string
	fmt.Printf("The value %v has type %v\n", b, reflect.TypeOf(b)) // 73 -> int32

	s := "I"
	fmt.Println(romanToInt(s))

}

/*
	這題有點魔性 ... 表面上看起來要考慮到 III 以及 II 的狀況 ... 但是反向來做，完全不用。所以取數值要反過來取(從尾巴開始看)
	temp 緩存 取 byte 後的數字，
	判斷數字是否小於接下來的數字(備註: 羅馬數字表示有一定的規則... 能放置在前面的只有 I ...)
	備註: 不用額外去考慮 `III` -> loop ... 1 + 1 + 1 (非嚴格遞增)
*/
