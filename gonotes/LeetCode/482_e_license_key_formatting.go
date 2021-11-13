package main

/*
You are given a license key represented as a string S which consists only alphanumeric character and dashes.
The string is separated into N+1 groups by N dashes.

Given a number K, we would want to reformat the strings such that each group contains exactly K characters,
except for the first group which could be shorter than K, but still must contain at least one character.

Furthermore, there must be a dash inserted between two groups and all lowercase letters should be converted to uppercase.
Given a non-empty string S and a number K, format the string according to the rules described above.

重組字串，將字串依照`K`的大小重新組合成若干個群，每個群需要用`-`來做區隔


Example 1:
Input: S = "5F3Z-2e-9-w", K = 4
Output: "5F3Z-2E9W"
Explanation: The string S has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.

Example 2:
Input: S = "2-5g-3-J", K = 2
Output: "2-5G-3J"
Explanation: The string S has been split into three parts,
each part has 2 characters except the first part as it could be shorter as mentioned above.
Note:
The length of string S will not exceed 12,000, and K is a positive integer.
String S consists only of alphanumerical characters (a-z and/or A-Z and/or 0-9) and dashes(-).
String S is non-empty.
*/

func licenseKeyFormatting(S string, K int) string {
	ret := make([]byte, 0, len(S))
	// 從最後一位元開始組合回來，組合好以後在反轉回原本的字串
	for i := len(S) - 1; i >= 0; i-- {
		// 遇到'-'忽略
		if S[i] == '-' {
			continue
		}

		// 如果當前組合長度除以組數的餘數恰好為K則加上`-`
		if len(ret)%(K+1) == K {
			ret = append(ret, '-')
		}

		// 將字元append進ret前，將小寫轉成大寫
		ret = append(ret, toUpper(S[i]))
	}
	reverse(ret)
	return string(ret)
}

func toUpper(c byte) byte {
	if 'a' <= c && c <= 'z' {
		return c - 'a' + 'A'
	}
	return c
}

func reverse(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}

func main() {

}

/*
solution:
- https://leetcode.com/problems/license-key-formatting/discuss/318360/Go-slice-magic


func licenseKeyFormatting(S string, K int) string {
    S = strings.ToUpper(S)
    S = strings.Replace(S, "-", "", -1)

    for i := len(S) - 1 - K; i >= 0; i = i - K {
        S = S[0:i + 1] + "-" + S[i + 1:]
    }

    return S
}
*/

/*
solution2:
- https://leetcode.com/problems/license-key-formatting/discuss/223701/2-foolish-solutions-by-golang-GO-4ms-292ms


func licenseKeyFormatting(S string, K int) string {
    var result []byte
    var temp byte
    var flag int
    for i:=0;i<len(S);i++{
        if S[i]!='-'{
            temp=S[i]
            if int(S[i])-int('Z')>0{
                temp=byte(int(S[i])-32)
            }
            result=append(result,temp)
        }
    }
    if len(result)<=K{
        return string(result)
    }
    var len1 int//len1 is the number of '-' in the last answer
    if len(result)%K==0{
        len1=len(result)/K-1
    }else{
        len1=len(result)/K
    }
    result1:=make([]byte,len1+len(result))
    j:=len1+len(result)-1
    flag=0
    for i:=len(result)-1;i>-1;i--{
        result1[j]=result[i]
        j--
        flag++
        if flag==K&&j>0{
            result1[j]='-'
            flag=0
            j--
        }
    }
    return string(result1)
}
*/
