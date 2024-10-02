package capitulo_02

import "fmt"

func BasesNumeracion() {

	// Binario
	const binario = 0b10001101

	// Octal
	const octal = 012072

	// Hexadecimal
	const hexadecimal = 0xAF32

	// Cuando se quieren agrupar bloques de digitos, podemos insertar un "_", go no lo tendra en cuenta al momento de establecer el valor
	const produtoInternoBruto = 1_419_000_000_000
	const bitsAgrupados = 0b_1000_1001_0110

	fmt.Println(binario)
	fmt.Println(octal)
	fmt.Println(hexadecimal)
	fmt.Println(produtoInternoBruto)
	fmt.Println(bitsAgrupados)

}
