/*
At a lemonade stand, each lemonade costs $5. Customers are standing in a queue to buy from you and order one at a time (in the order specified by bills). Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill. You must provide the correct change to each customer so that the net transaction is that the customer pays $5.

Note that you do not have any change in hand at first.

Given an integer array bills where bills[i] is the bill the ith customer pays, return true if you can provide every customer with the correct change, or false otherwise.
*/
package main

func lemonadeChange(bills []int) bool {
	change := map[int]int{}
	for _, v := range bills {
		change[v]++
		if v != 5 {
			currChange := v - 5
			for i := currChange / 5; i > 0 && currChange > 0; {
				if change[5*i] > 0 && currChange-5*i >= 0 {
					change[5*i]--
					currChange -= 5 * i
				} else {
					i--
				}
			}
			if currChange != 0 {
				return false
			}
		}
	}
	return true
}

//5,5,5,5,10,5,10,10,10,20
