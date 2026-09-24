package main

import (
	"Taller/Herramientas/conversor"
	"Taller/Herramientas/vocales"
	"fmt"
)

func main() {
	var opcion string

	for {
		fmt.Println("\n1. Conversor de monedas")
		fmt.Println("2. Contador de vocales")
		fmt.Println("0 o salir")
		fmt.Print("Elija una opción: ")
		fmt.Scanln(&opcion)

		if opcion == "0" || opcion == "salir" {
			break
		}

		switch opcion {
		case "1":
			var dolares float64
			var moneda string
			fmt.Print("Ingrese dólares: ")
			fmt.Scanln(&dolares)
			fmt.Print("Ingrese moneda (Euros, LB, Won, BTC): ")
			fmt.Scanln(&moneda)

			conversor.ConvertirMoneda(dolares, moneda)

		case "2":
			var frase string
			fmt.Print("Ingrese la frase (una sola palabra): ")
			fmt.Scanln(&frase)

			vocales.ContarVocales(frase)
		}
	}
}
