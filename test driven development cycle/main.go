package main

import (
	"fmt"
)

func Addition(val1, val2 any) any {

	switch v1 := val1.(type) {
	case int:
		v2, err := val2.(int)
		if err {
			return v1 + v2
		}
	case string:
		v2, err := val2.(string)
		if err {
			return v1 + v2
		}
	case float64:
		v2, err := val2.(float64)
		if err {
			return v1 + v2
		}
	default:
		fmt.Println("not a valid type")
	}
	return nil
}

func main() {
	var name string
	fmt.Scan(name)
	fmt.Println(Addition("tanish ", "shivhare"))
}
