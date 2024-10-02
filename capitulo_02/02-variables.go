package capitulo_02

import "fmt"

func Variables() {

	var dias int
	var meses int = 12

	// definicion de la variable con un valor inicial
	paso := 1

	// cambio del valor de la variable
	paso = 2

	/**
		Las guias de estilo recomiendan declarar variables con el operador := siempre que se pueda, aunque las variables globales deberan declararse mediante la forma precedida con var
	**/

	fmt.Println(dias)
	fmt.Println(meses)
	fmt.Println(paso)

}
