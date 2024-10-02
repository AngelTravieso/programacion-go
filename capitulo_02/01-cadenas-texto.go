package capitulo_02

import "fmt"

func CadenasTexto() {

	texto := "-Hola, como estas? \n- Estoy bien, gracias"
	texto2 := "Podria decirse que estoy \"bien\"..."

	templateTexto := `
		- Hola, como estas?
		- Estoy bien, gracias.
	`

	fmt.Println(texto)
	fmt.Println(texto2)
	fmt.Println(templateTexto)

}
