package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	fmt.Println("Age pointer: ", *agePointer)
	adultYears := getAdultYears(agePointer)

	fmt.Println("Adult Years: ", adultYears)

	fmt.Println("Before overriding: ", age)
	overrideAge(&age)
	fmt.Println("After overriding: ", age)

}

func getAdultYears(age *int) int {
	fmt.Println("Age: ", age, &age, *&age)
	return *age - 18
}

func overrideAge(age *int) {
	*age += 10
}
