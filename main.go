package main

import "fmt"

func main() {
	map1 := map[string]int{
		"a": 1,
		"b": 2,
	}
	delete(map1, "b")

	map1["c"] = 5

	fmt.Println(map1)

	if value, ok := map1["d"]; ok {
		fmt.Println(ok, value)
	} else {
		fmt.Println("no found", ok)
		fmt.Println("hello")
	}

}
