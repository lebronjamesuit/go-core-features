package main

import "fmt"

// {1, 2, 4, 8, 16, 32, 64, 128}
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
var m = make(map[string]string)

func main() {
	for i, v := range pow {
		fmt.Printf("this is value %d", v)
		fmt.Println(" ")
		fmt.Printf("this is index %d", i)
		fmt.Println(" ")
	}
	m["k1"] = "hehe"
	m["k2"] = "haha"
	m["k3"] = "lol"

	for index, value := range m {
		fmt.Printf("this is value %s", value)
		fmt.Println(" ")
		fmt.Printf("this is index %s", index)
		fmt.Println(" ")
	}

}
