package main

/*
At a lemonade stand, each lemonade costs $5.

Customers are standing in a queue to buy from you, and order one at a time (in the order specified by bills).

Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill.

You must provide the correct change to each customer, so that the net transaction is that the customer pays $5.

Note that you don't have any change in hand at first.

Return true if and only if you can provide every customer with correct change.

Example 1:
Input: [5,5,5,10,20]
Output: true
Explanation:
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.

Example 2:
Input: [5,5,10]
Output: true

Example 3:
Input: [10,10]
Output: false

Example 4:
Input: [5,5,10,10,20]
Output: false
Explanation:
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can't give change of $15 back because we only have two $10 bills.
Since not every customer received correct change, the answer is false.

Note:

0 <= bills.length <= 10000
bills[i] will be either 5, 10, or 20.
*/

func lemonadeChange(bills []int) bool {
	if len(bills) == 0 {
		return true
	}
	if bills[0] > 5 {
		return false
	}
	bill_num := make(map[int]int)
	for _, v := range bills {
		switch v {
		case 5:
			bill_num[5]++
		case 10:
			if bill_num[5] > 0 {
				bill_num[10]++
				bill_num[5]--
			}
		case 20:
			if bill_num[10] > 0 && bill_num[5] > 0 {
				bill_num[20]++ // actually unnecessary
				bill_num[10]--
				bill_num[5]--
			} else if bill_num[5] > 2 {
				bill_num[20]++
				bill_num[5] -= 3
			} else {
				return false
			}
		default:
			return false
		}
	}
	return true
}

func main() {

}

/*
refer:
- https://leetcode.com/problems/lemonade-change/discuss/580749/go-clean-code-4ms-beats-100

func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		switch {
		case bill == 5:
			five++
		case bill == 10 && five > 0:
			five, ten = five-1, ten+1
		case bill == 20 && five > 0 && ten > 0:
			five, ten = five-1, ten-1
		case bill == 20 && five > 2:
			five -= 3
		default:
			return false
		}
	}
	return true
}
*/
