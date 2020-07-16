package main

import "fmt"

func cal(a, b int) (r1, r2 int) {
	r1 = a*b
	r2 = a+b 
	return 
}
func main() {
	mul, add := cal(2, 3)
	fmt.Println("Hello mul: ", mul)
	fmt.Println("Hello add: ", add)
}