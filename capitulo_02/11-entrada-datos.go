package capitulo_02

import "fmt"

func EntradaDatos() {

	var edad int
	fmt.Print("Edad? ")

	// fmt.Scan lee los datos del teclado
	fmt.Scan(&edad)

	fmt.Println("Tienes", edad, "años")

	// fmt.Scanf permite especificar con mas detalle el formato de la entrada
	var hora, minuto, segundo int

	fmt.Print("HH:MM:SS? ")

	fmt.Scanf("%d:%d:%d", &hora, &minuto, &segundo)

	fmt.Printf("%d horas, %d minutos, %d segundos", hora, minuto, segundo)

}
