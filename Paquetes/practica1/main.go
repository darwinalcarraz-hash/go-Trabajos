package main

import (
	"fmt"
	"practica/operaciones"
	"practica/saludo"
)

func main() {
	fmt.Println("------Bienvenid@s a la clase de paquetes------")
	mensaje := saludo.Saludar("josefina")
	fmt.Println(mensaje)

	fmt.Println(operaciones.Suma(5, 6))

	total := operaciones.Suma(10, 5)
	fmt.Println("La suma es: ", total)

	suma, resta := operaciones.SumaYResta(4, 10)
	fmt.Println("La suma es: ", suma)
	fmt.Println("La resta es: ", resta)

	variadica := operaciones.Sumatoria(5, 6, 7, 8, 9)
	fmt.Println("La sumatoria es: ", variadica)
}
