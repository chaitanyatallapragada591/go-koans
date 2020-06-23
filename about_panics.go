package go_koans

func divideFourBy(i int) int {
	var result int

	defer func() {
		if r := recover(); r != nil {
			println("Cannot divide by Zero")

		}
	}()
	result = 4 / i
	return result
}

const __divisor__ = 0

var n int

func aboutPanics() {
	//assert(__delete_me__) // panics are exceptional errors at runtime

	n = divideFourBy(__divisor__)
	n = divideFourBy(2)

	assert(n == 2) // panics are exceptional errors at runtime
}
