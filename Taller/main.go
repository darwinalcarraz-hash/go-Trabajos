package main

import "fmt"

func main() {
	var opcion string

	for {
		fmt.Println("\n=====================================")
		fmt.Println("          MENÚ DE OPCIONES")
		fmt.Println("=====================================")
		fmt.Println("1. Calcular promedio de un curso")
		fmt.Println("2. Calcular la suma del 1 al n")
		fmt.Println("3. Convertir de Celsius a Fahrenheit")
		fmt.Println("4. Convertir de Fahrenheit a Celsius")
		fmt.Println("0 o 'salir'. Terminar programa")
		fmt.Print("Elija una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("\nSaliendo del programa... ¡Hasta luego!")
			break
		}

		switch opcion {
		case "1":
			opcionEstudiantes()
		case "2":
			opcionSumaN()
		case "3":
			opcionCelsiusAFahrenheit()
		case "4":
			opcionFahrenheitACelsius()
		default:
			fmt.Println("Opción no válida. Por favor, intente de nuevo.")
		}
	}
}

// Función averageGrade que calcula el promedio de todas las notas ingresadas.
func averageGrade(notas []float64) float64 {
	var suma float64 = 0
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

// Opción 1 registro de alumnos, cálculo de promedio, aprobación y rendimiento
func opcionEstudiantes() {
	var cantidad int
	fmt.Print("\nIngrese la cantidad de estudiantes: ")
	fmt.Scan(&cantidad)

	if cantidad <= 0 {
		fmt.Println("La cantidad de estudiantes debe ser mayor a 0.")
		return
	}

	notas := make([]float64, cantidad)
	for i := 0; i < cantidad; i++ {
		fmt.Printf("Ingrese la nota del estudiante %d (0 a 100): ", i+1)
		fmt.Scan(&notas[i])
	}

	promedio := averageGrade(notas)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	if promedio >= 70 {
		fmt.Println("Estado del curso: Aprobado")
	} else {
		fmt.Println("Estado del curso: Reprobado")
	}

	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Mensaje: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Mensaje: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Mensaje: Satisfactory performance")
	default:
		fmt.Println("Mensaje: Needs improvement")
	}
}

// Opción 2 suma consecutiva de todos los números desde el 1 hasta n
func opcionSumaN() {
	var n int
	fmt.Print("\nIngrese un número entero (n): ")
	fmt.Scan(&n)

	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}
	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, suma)
}

// Opción 3 conversión de unidades de temperatura Celsius a Fahrenheit
func opcionCelsiusAFahrenheit() {
	var celsius float64
	fmt.Print("\nIngrese la temperatura en grados Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f °C equivalen a %.2f °F\n", celsius, fahrenheit)
}

// Opción 4 conversión de unidades de temperatura Fahrenheit a Celsius
func opcionFahrenheitACelsius() {
	var fahrenheit float64
	fmt.Print("\nIngrese la temperatura en grados Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f °F equivalen a %.2f °C\n", fahrenheit, celsius)
}