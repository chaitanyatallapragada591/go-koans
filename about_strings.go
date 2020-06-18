package go_koans

import "fmt"

func aboutStrings() {
	var __string__ string = "bc"
	var __int__ int = 3
	assert("a"+__string__ == "abc") // string concatenation need not be difficult
	assert(len("abc") == __int__)   // and bounds are thoroughly checked
	var __byte__ byte = 97
	assert("abc"[0] == __byte__)  // their contents are reminiscent of C
	assert("smith"[2:] == "ith")  // slicing may omit the end point
	assert("smith"[:4] == "smit") // or the beginning
	assert("smith"[2:4] == "it")  // or neither
	assert("smith"[:] == "smith") // or both

	assert("smith" == "smith") // they can be compared directly
	assert("smith" < "z")      // i suppose maybe this could be useful.. someday

	bytes := []byte{'a', 'b', 'c'}
	assert(string(bytes) == "abc") // strings can be created from byte-slices

	bytes[0] = 'z'
	assert(string(bytes) == "zbc") // byte-slices can be mutated, although strings cannot
	//var worldString string = "world"
	assert(fmt.Sprintf("hello %s", "world") == "hello world")                                  // our old friend sprintf returns
	assert(fmt.Sprintf("hello \"%s\"", "world") == "hello \"world\"")                          // quoting is familiar
	assert(fmt.Sprintf("hello %q", "world") == "hello \"world\"")                              // although it can be done more easily
	assert(fmt.Sprintf("your balance: %d and %0.2f", 3, 4.5589) == "your balance: 3 and 4.56") // "the root of all evil" is actually a misquotation, by the way
}
