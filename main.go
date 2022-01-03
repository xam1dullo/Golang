package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	fmt.Println(pow)
	var _ = map[string]int{"a": 1, "b": 2}
	 age := 21
	 name := "Xamidullo"
	 fmt.Printf("my age is ", age , "My name is ",name)
}
