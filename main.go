package main // entry package

import "fmt" // import package

func main() {
	fmt.Println("Hello World")

	// VARIABLES
	// var/const  name  type
	var name string = "kriti"
	fmt.Println(name)

	// Go automatically understands it is a string.
	var car = "DEFENDER"
	fmt.Println(car)

	// Short Declaration (Most Used)
	// declare + assign
	name1 := "kriti"
	fmt.Println(name1)

	var age int
	fmt.Println(age)
	// var age ❌  no type declared

	// 5. Data Types
	// Integer int int8 int16 int32 int64
	// unsigned ( only positive ) -- uint uint8 uint16 uint32 uint64

	if age >= 18 {
		fmt.Println("Able to vote ")
	} else {
		fmt.Println("not able to ")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// no while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	// INFINITE LOOP
	// for {
	// 	fmt.Println("kriti")
	// }

	// RANGE LOOP
	// for index, value := range collection {
	// // work
	// }

	// ARRAYS
	arr := []int{10, 20, 30} // -- slice
	for i, value := range arr {
		fmt.Println(i, value)
	}

	for i, ch := range "kriti" {
		fmt.Println(i, ch)
	}

	// ARRAYS -- FIXED SIZE --- type alsoo to tell
	// var arrr[5] int
	// size exactly is = 5

	// SLICE -- dyanamic array
	// these builtin functions are for slice not for array
	arrr := []int{1, 2, 3, 4}
	arrr = append(arrr, 5)

	// maps
	ages := map[string]int{
		"kriti": 20,
	}

	fmt.Println(ages["kriti"])

	// insert
	ages["rahul"] = 11

	// custom map
	Custommap := map[string]int{}
	Custommap["kriti"] = 60
	Custommap["Ankita"] = 55

	for key, value := range Custommap {
		fmt.Println(key, "-->", value)
	}

	// functions
	//arrays_demo()
	//loop_demo()
}

// input type -- int 		return type -- string and int
func Checkevenodd(num int) (string, int) {
	if num%2 == 0 {
		return "Even", 0
	}
	return "Odd", 1
}

func Pointerss() {
	// demonstrating pointers
	i := 200

	var ptr *int = &i
	ptr1 := &i

	fmt.Println("Value of i:", i, "Pointer to i:", ptr, "Pointer1 to i:", ptr1)
	fmt.Println("value present at pointer", *ptr)

}

func loop_demo() {

	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}

	for i := range 2 {
		fmt.Println(i)
	}

	for i, char := range "Sanket" {
		fmt.Println("String range loop iteration:", i, char)
	}

	// Below is a while loop using a for loop

	j := 10

	for j > 0 {
		if j == 3 {
			j--
			continue
		} else if j == 5 {
			fmt.Println("Skipping 5")
			j--
			continue
		} else {
			fmt.Println("While loop iteration:", j)
			j--
			continue
		}

	}

	for {
		fmt.Println("Infinite loop iteration")
		break
	}
}

func arrays_demo() {

	var arr [2]int

	arr[0] = 33
	arr[1] = 34

	// slice
	var ar []int
	ar = append(ar, 1)

}
