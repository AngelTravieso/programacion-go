package capitulo_02

import "fmt"

func OperadoresLogicos() {

	// &&
	a := 6 == 6 && 3 > 2

	fmt.Println(a)

	// ||
	b := false || 3 < 8

	fmt.Println(b)

	// !
	c := !(7 < 1)

	fmt.Println(c)

}
