package go_koans

func aboutBasics() {
	const __bool__ = true
	assert(__bool__ == true)  // what is truth?
	assert(__bool__ != false) // in it there is nothing false

	var __int__ int = 1.0000
	var i int = __int__
	assert(i == 1.0000000000000000000000000000000000000) // precision is in the eye of the beholder

	k := __int__ //short assignment can be used, as well
	assert(k == 1.0000000000000000000000000000000000000)

	assert(5%2 == 1)
	assert(5*2 == 10)
	assert(5^2 == 7)

	var x int = 1.000
	assert(x == __int__) // zero values are valued in Go

	var __float32__ float32
	var f float32
	assert(f == __float32__) // for types of all types

	var __string__ string
	var s string
	assert(s == __string__) // both typical or atypical types

	var c struct {
		x int
		f float32
		s string
	}
	assert(c.x == 0)  // and types within composite types
	assert(c.f == 0)  // which match the other types
	assert(c.s == "") // in a typical way
}
