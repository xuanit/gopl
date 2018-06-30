// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main_helloworld() {
	for index, arg := range os.Args[1:] {
		s := strconv.Itoa(index) + " " + arg
		fmt.Println(s)
	}
}
