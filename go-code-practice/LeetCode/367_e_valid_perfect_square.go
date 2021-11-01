package main

/*
Given a positive integer num, write a function which returns True
if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:
Input: 16
Output: true

Example 2:
Input: 14
Output: false
*/

/*
// runtime 708ms; memory 1.9MB
func isPerfectSquare(num int) bool {
	if num==1 {
		return true
	}

	i:=1
	for i<num{
		if i*i == num{
			return true
		}
		i++
	}
	return false
}
*/

// 用 2 分法 逐次收斂來加速
// m1 := (1+num)/2
// m2 := {func(m1)+m1}/2
func isPerfectSquare(num int) bool {
	lo := 1
	hi := num
	for lo <= hi {
		med := (lo + hi) >> 1 // 中間值 為 lo(頭) + hi(尾) / 2
		val := med * med
		if num == val {
			return true
		}

		// 判斷 平方值 是否小於 num，將中間值移動
		if val < num {
			lo = med + 1
		} else {
			hi = med - 1
		}
	}
	return false
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/valid-perfect-square/discuss/450533/Go%3A-100


func isPerfectSquare(num int) bool {
    lo := 1
    hi := num
    for lo <= hi {
        med := (lo+hi) >> 1
        val := med * med
        if num == val {
            return true
        }
        if val < num {
            lo = med + 1
        }else{
            hi = med - 1
        }
    }
    return false
}


// triangular number
- https://simple.wikipedia.org/wiki/Triangular_number

o

o
oo

o
oo
ooo

o
oo
ooo
oooo

func isPerfectSquare(num int) bool {
    pre := 0
    cur := 0
    for idx := 0; pre + cur < num; idx++{
        pre = cur
        cur = cur + idx
        if pre + cur == num{
            return true
        }
    }
    return false
}
*/
