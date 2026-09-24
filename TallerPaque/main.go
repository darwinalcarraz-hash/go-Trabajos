package main

import (
	"Taller/Herramientas/conversor"
	"Taller/Herramientas/vocales"
	"bufio"
	"fmt"
	"os"
)

func main() {
	var opcion string

	for {
		fmt.Println("=======================Menú de opciones=======================")
		fmt.Println("1. Conversor de monedas")
		fmt.Println("2. Contador de vocales")
		fmt.Println("Presione 0 o escriba 'salir' para terminar el programa")
		fmt.Print("Elija una opción: ")
		fmt.Scanln(&opcion)

		if opcion == "0" || opcion == "salir" {
			break
		}

		switch opcion {
		case "1":
			var dolares float64
			var tipoMoneda int

			fmt.Print("Ingrese dólares: ")
			fmt.Scanln(&dolares)

			fmt.Println("\nSeleccione la moneda destino:")
			fmt.Println("1. Euros")
			fmt.Println("2. LB (Libras Esterlinas)")
			fmt.Println("3. Won (Sur Coreano)")
			fmt.Println("4. BTC")
			fmt.Print("Elija una opción (1-4): ")
			fmt.Scanln(&tipoMoneda)

			conversor.ConvertirMoneda(dolares, tipoMoneda)

		case "2":
			fmt.Print("Ingrese la frase: ")
			fmt.Scanln()

			lector := bufio.NewReader(os.Stdin)
			frase, _ := lector.ReadString('\n')

			vocales.ContarVocales(frase)
		}
	}
}
