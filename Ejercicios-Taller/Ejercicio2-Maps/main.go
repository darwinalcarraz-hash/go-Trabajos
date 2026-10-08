package main

import "fmt"

func determinarGanador(votacion map[string]int) {
	mayorVotos := -1
	actividadGanadora := ""

	for actividad, votos := range votacion {
		if votos > mayorVotos {
			mayorVotos = votos
			actividadGanadora = actividad
		}
	}

	fmt.Println("\n=========================================")
	fmt.Println("La actividad ganadora es: ", actividadGanadora, " con ", mayorVotos, " votos")
	fmt.Println("=========================================")
}

func main() {
	votosActividades := map[string]int{
		"deportes":    0,
		"videojuegos": 0,
		"cine":        0,
		"música":      0,
	}

	fmt.Println("=== REGISTRO DE VOTOS (5 ESTUDIANTES) ===")

	for i := 1; i <= 5; i++ {
		var seleccion int
		fmt.Println("\nVoto número: ", i)
		fmt.Println("1. Deportes")
		fmt.Println("2. Videojuegos")
		fmt.Println("3. Cine")
		fmt.Println("4. Música")
		fmt.Print("Seleccione un número (1-4): ")
		fmt.Scanln(&seleccion)

		switch seleccion {
		case 1:
			votosActividades["deportes"]++
		case 2:
			votosActividades["videojuegos"]++
		case 3:
			votosActividades["cine"]++
		case 4:
			votosActividades["música"]++
		default:
			fmt.Println("Opción no válida. Voto perdido.")
		}
	}

	fmt.Println("\n======= RESULTADOS DE LA VOTACIÓN =======")
	for actividad, totalVotos := range votosActividades {
		fmt.Println("- ", actividad, ": ", totalVotos, " votos")
	}

	determinarGanador(votosActividades)
}
