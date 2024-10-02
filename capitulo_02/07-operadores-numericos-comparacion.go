package capitulo_02

import "fmt"

func OperadoresNumericosComparacion() {

	a := 3
	b := 5

	// ==
	resultado := a+b == 8

	fmt.Println(resultado)

	// !=
	resultado = a+b != 8

	fmt.Println(resultado)

	// <
	resultado = 3 < 2

	fmt.Println(resultado)

	// <=
	resultado = 3 <= 1+2

	fmt.Println(resultado)

	// >
	resultado = 3 > 3-1

	fmt.Println(resultado)

	// >=
	resultado = 3 >= 2

	fmt.Println(resultado)

}
