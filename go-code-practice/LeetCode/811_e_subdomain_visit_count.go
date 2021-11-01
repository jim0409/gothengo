package main

import "strconv"

/*
A website domain like "discuss.leetcode.com" consists of various subdomains.
At the top level, we have "com", at the next level, we have "leetcode.com", and at the lowest level, "discuss.leetcode.com".

When we visit a domain like "discuss.leetcode.com", we will also visit the parent domains "leetcode.com" and "com" implicitly.

Now, call a "count-paired domain" to be a count (representing the number of visits this domain received),
followed by a space, followed by the address. An example of a count-paired domain might be "9001 discuss.leetcode.com".

We are given a list cpdomains of count-paired domains.
We would like a list of count-paired domains, (in the same format as the input, and in any order),
that explicitly counts the number of visits to each subdomain.

Example 1:
Input:
["9001 discuss.leetcode.com"]
Output:
["9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"]
Explanation:
We only have one website domain: "discuss.leetcode.com".
As discussed above, the subdomain "leetcode.com" and "com" will also be visited. So they will all be visited 9001 times.

Example 2:
Input:
["900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"]
Output:
["901 mail.com","50 yahoo.com","900 google.mail.com","5 wiki.org","5 org","1 intel.mail.com","951 com"]
Explanation:
We will visit "google.mail.com" 900 times, "yahoo.com" 50 times, "intel.mail.com" once and "wiki.org" 5 times.
For the subdomains, we will visit "mail.com" 900 + 1 = 901 times, "com" 900 + 50 + 1 = 951 times, and "org" 5 times.

Notes:

The length of cpdomains will not exceed 100.
The length of each domain name will not exceed 100.
Each address will have either 1 or 2 "." characters.
The input count in any count-paired domain will not exceed 10000.
The answer output can be returned in any order.
*/

// 1. split domain array with `times` and `domain name`
// 2. count domain name from `root` to `sub`
// https://leetcode.com/problems/subdomain-visit-count/discuss/346844/Go-golang-8ms-solution
func subdomainVisits(cpdomains []string) []string {

	tempSlice := []string{}
	tempMap := make(map[string]int)
	var n int
	var s string

	// 先拆分次數
	for _, v := range cpdomains {
		// n 表示訪問次數
		// s 表示域名
		// 這段相當於做 string.Split(v,",")
		for I, V := range v {
			if V == ' ' {
				n, _ = strconv.Atoi(v[:I])
				s = v[I+1:]
			}
		}

		// 紀錄每一個域名的訪問次數
		tempMap[s] += n
		for i, v := range s {
			var tempS string
			if v == '.' {
				tempS = s[i+1:]
				tempMap[tempS] += n
			}
		}
	}

	for i, v := range tempMap {
		a := strconv.Itoa(v) + " " + i
		tempSlice = append(tempSlice, a)
	}

	return tempSlice

}

func main() {

}

/*
refer:
- https://leetcode.com/problems/subdomain-visit-count/discuss/224211/Golang-100

import (
    "strings"
    "strconv"
)

func subdomainVisits(cpdomains []string) []string {
    domainsMap := make(map[string]int, len(cpdomains) * 2)

    for _, cpdomain := range cpdomains {
        raw        := strings.Split(cpdomain, " ")
        visits,  _ := strconv.Atoi(raw[0])
        domains    := strings.Split(raw[1], ".")

        for i, _:= range domains {
            domainsMap[strings.Join(domains[i:], ".")] += visits
        }
    }

    answer := make([]string, 0, len(domainsMap))

    for domain, visits := range domainsMap {
        answer = append(answer, strconv.Itoa(visits) + " " + domain)
    }

    return answer
}
*/
