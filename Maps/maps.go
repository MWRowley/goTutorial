package main

import (
	"fmt"
	"time"
)

func main() {
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 = map[string]uint{"Adam": 12, "Sarah": 24}
	fmt.Println(myMap2["Adam"]) // a map will always return a zero value if the key does not exist
	var age, ok = myMap2["Jason"]
	delete(myMap2, "Adam") // delete a key
	if ok {
		fmt.Printf("The age is %v", age)
	} else {
		fmt.Printf("Invalid Name \n")
	}

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	intArr := [...]int32{1, 2, 3}
	fmt.Println(intArr)

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice, n))
	fmt.Printf("Total time without preallocation: %v \n", timeLoop(testSlice2, n))

}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	return time.Since(t0)

}
