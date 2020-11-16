package main

import (
	"fmt"
)

func main(a int) (square int, cube int) {
	square = a*a
	cube = square*a
	return square,cube,nil
}
}
