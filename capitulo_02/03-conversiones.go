package capitulo_02

import "fmt"

func Conversiones() {

	// int8 = int de 8 bytes
	// go requiere de manera explicita la conversion

	var segundos int8 = 30
	horas := int(segundos)

	fmt.Println(segundos)
	fmt.Println(horas)

	// Si la conversion se hace desde un tipo de coma flotante, se trunca la parte con decimales
	distancia := 12.78
	km := int(distancia)

	fmt.Println(km)

}
