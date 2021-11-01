package main

/*
Every valid email consists of a local name and a domain name, separated by the '@' sign.
Besides lowercase letters, the email may contain one or more '.' or '+'.

For example, in "alice@leetcode.com", "alice" is the local name, and "leetcode.com" is the domain name.
If you add periods '.' between some characters in the local name part of an email address,
mail sent there will be forwarded to the same address without dots in the local name. Note that this rule does not apply to domain names.

For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
If you add a plus '+' in the local name, everything after the first plus sign will be ignored.
This allows certain emails to be filtered. Note that this rule does not apply to domain names.

For example, "m.y+name@email.com" will be forwarded to "my@email.com".
It is possible to use both of these rules at the same time.

Given an array of strings emails where we send one email to each email[i], return the number of different addresses that actually receive mails.


Example 1:
Input: emails = ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails.

Example 2:
Input: emails = ["a@leetcode.com","b@leetcode.com","c@leetcode.com"]
Output: 3


Constraints:
1 <= emails.length <= 100
1 <= emails[i].length <= 100
email[i] consist of lowercase English letters, '+', '.' and '@'.
Each emails[i] contains exactly one '@' character.
All local and domain names are non-empty.
Local names do not start with a '+' character.
*/

// https://leetcode.com/problems/unique-email-addresses/discuss/220982/97-use-go
func numUniqueEmails(emails []string) int {
	mp := make(map[string]int)

	// 迭代 `emails array` 長度，尋遍 `email` 字串
	for i := 0; i < len(emails); i++ {
		em := []byte(emails[i])
		k := 0
		findDot := 0
		findPlus := false
		// 判斷 email[i] 內的字，判斷是否為特殊符號
		for _, val := range emails[i] {
			// 紀錄`@`的位置`
			if val == '@' {
				findPlus = false
				findDot++
			}
			// 紀錄`.`的位置
			if val == '.' || findPlus {
				if findDot == 0 {
					continue
				}
			}
			// 紀錄`+`的位置
			if val == '+' {
				findPlus = true
				continue
			}

			em[k] = byte(val)
			k++
		}
		// 如果恰好只有一個 `.`且 k > 0 則 對 mail map 做 hash -- string(em)
		if k > 0 && findDot == 1 {
			mp[string(em)[0:k]]++
		}
	}

	// 回傳相異郵件的數量
	return len(mp)
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/unique-email-addresses/discuss/992608/Simple-Go-solution-using-Go-builtins

func numUniqueEmails(emails []string) int {
	s := make(map[string]bool)
	for _, e := range emails {
		i := strings.Index(e, "@")
		l := string(e[:i])
		if j := strings.Index(l, "+"); j > -1 {
			l = string(l[:j])
		}
		l = strings.Replace(l, ".", "", -1)
		s[l+e[i:]] = true
	}
	return len(s)
}
*/
