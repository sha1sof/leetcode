package task1

func distMoney(money int, children int) int {
	if money < children {
		return -1
	}

	if money > 8*children {
		return children - 1
	}
	if money == 8*children-4 {
		return children - 2
	}

	return (money - children) / 7
}

/*package task1

func distMoney(money int, children int) int {
	if money < 8 || children < 1 {
		return -1
	}
	result := 0

	for money >= 0{
		if money > 7 && children > 0 {
			result++
			money -=8
			children--
		} else if money != 4 && children > 0{
			result++
			children--
		} else if money == 4 && children>{}
	}

	if result == 0 {
		return -1
	} else{
		return result
	}
}
*/
