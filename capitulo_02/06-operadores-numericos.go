package capitulo_02

import "fmt"

func OperadoresNumericos() {

	a := 8 + 3*(1+2)%5

	fmt.Println(a)

	a = 10
	b := 20

	// Incremento
	a++

	// Decremento
	b++

	fmt.Println(a)
	fmt.Println(b)

}
