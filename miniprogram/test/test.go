package main

import "fmt"

func main() {
	aa := make(map[string]interface{})
	aa["2"] = 3

	i64, ok := aa["2"].(int)
	fmt.Println(ok)
	fmt.Println(i64)
}
