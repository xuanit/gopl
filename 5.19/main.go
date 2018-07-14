package main

import "fmt"

type bailout struct {
	X int
	I internal
}
type internal struct{ V int }

func noReturn() (result int) {
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			result = 2
		}
	}()

	panic(bailout{})
}

func main() {
	i := internal{V: 1}
	i2 := internal{V: 1}
	fmt.Println(bailout{X: 1, I: i} == bailout{X: 1, I: i2})
}
