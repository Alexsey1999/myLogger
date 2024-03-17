package myLogger

import "fmt"

func PrintColors() {
	colors := []string{"red", "green", "blue"}

	for _, val := range colors {
		fmt.Println(val)
	}
}
