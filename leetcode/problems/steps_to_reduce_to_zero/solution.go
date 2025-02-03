package stepstoreducetozero

func isEven(num int) bool {
	return num%2 == 0
}

func numberOfSteps(num int) int {
	var steps int = 0

	for num > 0 {

		if isEven(num) {
			num /= 2
		} else {
			num -= 1
		}

		steps += 1

	}

	return steps

}
