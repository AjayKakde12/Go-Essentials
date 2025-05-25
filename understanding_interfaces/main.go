package main

import "fmt"

func add[T int | float64 | string](a, b T) T {
	// aInt, aIsInt := a.(int)
	// bInt, bIsInt := b.(int)
	// if aIsInt && bIsInt {
	// 	return aInt + bInt
	// }

	// aFloat, aIsFloat := a.(float64)
	// bFloat, bIsFloat := b.(float64)
	// if aIsFloat && bIsFloat {
	// 	return aFloat + bFloat
	// }

	// aString, aIsString := a.(string)
	// bString, bIsString := b.(string)
	// if aIsString && bIsString {
	// 	return aString + bString
	// }

	return a + b
}

func main() {
	sum := add(2, 3)
	fmt.Println("Addition: ", sum)
}
