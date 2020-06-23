package go_koans

func isPrimeNumber(possiblePrime int) bool {
	for underPrime := 2; underPrime < possiblePrime; underPrime++ {
		if possiblePrime%underPrime == 0 {
			return false
		}
	}
	return true
}

func findPrimeNumbers(channel chan int) {
	for i := 2; i < 100; i++ {
		// your code goes here
		if isPrimeNumber(i) {
			channel <- i
			// i is afraid of heights
		}
		assert(i < 100)
	}
}

func aboutConcurrency() {
	ch := make(chan int)

	go findPrimeNumbers(ch)

	//assert(__delete_me__) // concurrency can be almost trivial
	// your code goes here

	assert(<-ch == 2)
	assert(<-ch == 3)
	assert(<-ch == 5)
	assert(<-ch == 7)
	assert(<-ch == 11)
}
