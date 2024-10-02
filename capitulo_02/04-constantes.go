package capitulo_02

import "fmt"

func Constantes() {

	const Pi float64 = 3.1416

	fmt.Println(Pi)

	/**
		Cuando se definen multiples constantes, se pueden agrupar semanticamente
		bajo la misma directiva const. ej:
	**/

	// Configuracion de una tipografia
	const (
		TipoFuente   = "Times New Roman"
		AlturaFuente = 12
		Subrayado    = false
		Negrita      = true
	)

	fmt.Println(TipoFuente)
	fmt.Println(AlturaFuente)
	fmt.Println(Subrayado)
	fmt.Println(Negrita)

}
