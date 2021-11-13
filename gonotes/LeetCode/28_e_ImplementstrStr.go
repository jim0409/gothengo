package main

import (
	"fmt"
	"log"
	"strings"
)

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

*/

func strStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	if ok := strings.Contains(haystack, needle); ok {
		haystackl := len(haystack)
		window := len(needle)
		// log.Println(haystackl)
		// log.Println(window)

		for i, _ := range haystack {
			if i+window <= haystackl {
				log.Println("value of i:", i)
				tmp := fmt.Sprintf("%v", haystack[i:i+window])
				// log.Println(tmp)
				// log.Println(needle)
				if needle == tmp {
					log.Println("enterHere")
					return i
				}
			}
		}
	}

	return -1
}

func main() {
	haystack := "mississippi"

	needle := "pi"
	log.Println(strStr(haystack, needle))
}

/*
func strStr(haystack string, needle string) int {
    for i := 0; i < len(haystack)-len(needle)+1; i++ {
        j := 0
        for ; j < len(needle); j++ {
            if needle[j] != haystack[i+j] {
                break
            }
        }
        if j == len(needle) {
            return i
        }
    }
    return -1
}

*/
