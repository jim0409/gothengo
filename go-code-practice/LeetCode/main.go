package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
	"unicode/utf8"
)

func main() {
	test24()
}

func test1() {
	s := "string"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i:i+1])
	}

	// a1 := 1000000000000000000000000000001 // overflows int
	// fmt.Println(a1)

}

func test2() {
	s := "string"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", s[i])
	}

}

func test3() {
	m := map[byte]int{
		'I': 1,
		'V': 2,
	}

	s := "IV" // expect to be 5

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v ", m[s[i]])
	}
}

func test4() {
	a := 1
	b := 2
	fmt.Println(a ^ b)
}

func test5() {
	a := uint32(1) // a = `1`
	c := a & 1
	a = a << 1 // a= `10`
	a = a << 1 // a = `100`
	fmt.Printf("%b\n", a)
	fmt.Printf("%b\n", c)

}

// https://www.tutorialspoint.com/go/go_bitwise_operators.htm
func test6() {
	/*

		Operator	Description	Example
		&	Binary AND Operator copies a bit to the result if it exists in both operands.	(A & B) will give 12, which is 0000 1100
		|	Binary OR Operator copies a bit if it exists in either operand.	(A | B) will give 61, which is 0011 1101
		^	Binary XOR Operator copies the bit if it is set in one operand but not both.	(A ^ B) will give 49, which is 0011 0001
		<<	Binary Left Shift Operator. The left operands value is moved left by the number of bits specified by the right operand.	A << 2 will give 240 which is 1111 0000
		>>	Binary Right Shift Operator. The left operands value is moved right by the number of bits specified by the right operand.	A >> 2 will give 15 which is 0000 1111

	*/

	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)

	// Line 1 - Value of c is 12
	// Line 2 - Value of c is 61
	// Line 3 - Value of c is 49
	// Line 4 - Value of c is 240
	// Line 5 - Value of c is 15
}

func test7() {
	a := 112122
	// for i, j := range a {
	// fmt.Println(i, j)
	// }
	for a > 0 {
		res := a % 10
		fmt.Println(res)
		// fmt.Println(res * res)
		a = a / 10
	}
}

func test8() {
	a := 8
	fmt.Println(a)
	fmt.Println(a & 1)
	fmt.Println(a&1 == 1)
	fmt.Println("=----------=")
	a >>= 1
	fmt.Println(a)
	fmt.Println(a & 1)
	fmt.Println(a&1 == 1)
}

func test9() {
	a := uint8(0x1) // 1 {0000 0001}
	fmt.Println(a)

	a >>= 1        // 向右移一位 --> 0
	fmt.Println(a) // 0

	b := uint8(0x4) // 4 {0000 0100}
	b >>= 1
	fmt.Println(b) // 2

	c := uint8(0xff) // 最大值 255 {1111 1111}
	fmt.Println(c)

	c >>= 1
	fmt.Println(c) // 向右移一位 {0111 1111} = 127

	d := uint8(0xf0) // {1111 0000}
	fmt.Println(d)
}

func test10() {
	a := 31
	for _, i := range string(a) {
		fmt.Println(i)
	}
}

func test11() {
	a := "test"
	fmt.Println(a[0])
}

func test12() {
	a := make(map[int]bool, 2)
	a[0] = true
	a[1] = false
	a[2] = true
	a[3] = true
	fmt.Println(a)
}

func test13() {
	a := "string"
	at := "string1"

	ra := []rune(a)
	rat := []rune(at)

	fmt.Println(ra)
	fmt.Println(rat)

	fmt.Println("-------")
	rla := rat[len(a)]
	fmt.Println(rla)
}

func test14() {
	// https://nanxiao.me/golang-byte-rune/
	byteSlice := func(b []byte) []byte {
		return b
	}

	runeSlice := func(r []rune) []rune {
		return r
	}

	b := []byte{0, 1}
	u8 := []uint8{2, 3}
	fmt.Printf("%T %T \n", b, u8)
	fmt.Println(byteSlice(b))
	fmt.Println(byteSlice(u8))

	r := []rune{4, 5}
	i32 := []int32{6, 7}
	fmt.Printf("%T %T \n", r, i32)
	fmt.Println(runeSlice(r))
	fmt.Println(runeSlice(i32))
}

func test15() {
	/*
		- https://juejin.im/post/5b44caebf265da0f491b8b83

		rune is an alias for int32 and is equivalent to int32 in all ways.
		It's used, by convention, to distinguish character values from integer values.
		(type rune = int32)
	*/

	var str = "hello 你好"               // 理論上應該是 8 (5個字符1個空格和兩個中文字)
	fmt.Println("len(str):", len(str)) // 12
	/*
		因為 golang 中 string 底層是通過byte數組實現的，中文字符在 unicode 下佔 2 個 bit
		在 utf-8 編碼下佔 3 個 bit，而 golang 的編碼正好是 utf-8
	*/

	// TODO: 預期要得到一個恰好為該字串符長度，而非編譯過後的長度
	fmt.Println("RuneCountInString:", utf8.RuneCountInString(str))
	fmt.Println("rune:", len([]rune(str)))

	/*
		golang 中還有另一個 `byte` 數據類型與 `rune` 相似，都是用來表達字符類型的變量類型
		差異在於；
		- byte 等同於 int8, 常用來處理 ASCII 字符
		- rune 等同於 int32, 常用來處理 unicode 或 utf-8 字符

		https://www.itread01.com/hkhkqple.html
	*/
}

func test16() {
	// `字串`轉`數字`
	i, err := strconv.Atoi("8888")
	if err != nil {
		panic(err)
	}
	i += 3
	println(i)

	// `數字`轉`字串`
	s := strconv.Itoa(333)
	s += "3"
	println(s)
}

func test17() {
	x := 12 // 1100
	y := 11 // 1011
	z := 10 // 1010

	fmt.Println(x ^ y) // 111
	fmt.Println(x ^ z) // 110

	a := x ^ y
	fmt.Println(a & 1) // 111 & 001

	b := x ^ z
	fmt.Println(b & 1) // 110 & 001
}

func test18() {
	s1 := "abab" // abab contained in `bababa`
	size1 := len(s1)
	fmt.Println(size1)

	ss1 := (s1 + s1)[1 : size1*2-1]
	fmt.Println(ss1)

	s2 := "aba" // aba doesn't contained in `baab`
	size2 := len(s2)
	fmt.Println(size2)

	ss2 := (s2 + s2)[1 : size2*2-1]
	fmt.Println(ss2)

	s3 := "abcabcabcabc" // abcabcabcabcabc contained in `bcabcabcabcabcabcabcab`
	size3 := len(s3)
	fmt.Println(size3)

	ss3 := (s3 + s3)[1 : size3*2-1]
	fmt.Println(ss3)
}

func test19() {
	// https://stackoverflow.com/questions/28465098/golang-bitwise-calculation
	// https://appdividend.com/2020/01/25/golang-operator-example-operators-in-go-tutorial/
	// https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
	// https://www.tutorialspoint.com/go/go_bitwise_operators.htm
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 ... need both true */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 ... at least one true */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 ... just one true */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}

func test20() {
	input1 := []int{1, 2, 4, 7, 6, 5, 3}
	sort.Ints(input1)
	fmt.Println(input1)
}

func test21() {
	a := math.MinInt32
	fmt.Println(a) // -2147483648
	b := math.MinInt16
	fmt.Println(b)
}

func test22() {
	// https://stackoverflow.com/questions/24122821/go-golang-time-now-unixnano-convert-to-milliseconds
	a := time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("%d \n", a)
}

func test23() {
	loop := 0
	s := 1
	for s <= 7 {
		fmt.Printf("loop round : %v %v\n", loop, s)
		s <<= 1
		loop++
	}

	fmt.Println("end loop : ", s)

	s -= 1
	fmt.Println(s)
}

// test over flow
func test24() {
	var a int
	var b uint = 8
	var c uint = 9
	fmt.Println(b - c)

	a = int(b - c)
	fmt.Println(a)
}

// test sort
func test25() {
	// https://openhome.cc/Gossip/Go/Sort.html
	s := []int{8, 2, 6, 3, 1, 4}
	sort.Ints(s)
	fmt.Println(sort.IntsAreSorted(s)) // true
	fmt.Println(s)                     // [1 2 3 4 6 8]
	fmt.Println(sort.SearchInts(s, 7)) // 5
}
