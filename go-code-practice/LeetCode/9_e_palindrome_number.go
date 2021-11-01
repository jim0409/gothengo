package main

import "fmt"

/*
	Determine whether an integer is a palindrome.
	An integer is a palindrome when it reads the same backward as forward.

	e.g.
	121 -> true

	-121 -> false
	10 -> false

*/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	tmpInt := x
	resultArr := []int{}

	// first for loop concrete a comparision array
	for tmpInt >= 10 {
		resultArr = append(resultArr, tmpInt%10)
		tmpInt = tmpInt / 10
	}
	resultArr = append(resultArr, tmpInt)

	// second loop compare the digit between this array
	lenResultArr := len(resultArr) - 1
	for i := 0; i < lenResultArr; i++ {

		// fmt.Println(i, resultArr[i])
		// fmt.Println(lenResultArr-i, resultArr[lenResultArr-i])

		if resultArr[i] == resultArr[lenResultArr-i] {
			continue
		} else {
			return false
		}
	}

	return true
}

func main() {
	/*
		Input:
		121 -> true
		-121 -> false
	*/

	a := 121
	fmt.Println(isPalindrome(a))

}

/*
func isPalindrome(x int) bool {
    s := strconv.Itoa(x)
    r := []rune(s)
    for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1 {
        if r[i] != r[j] {return false}
    }
    return true
}
*/
