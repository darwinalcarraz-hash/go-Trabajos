package conversor

import "fmt"

func ConvertirMoneda(dolares float64, opcionMoneda int) {
	switch opcionMoneda {
	case 1:
		fmt.Println("La conversión es: ", dolares*0.88, "Euros")
	case 2:
		fmt.Println("La conversión es: ", dolares*0.75, "Libras Esterlinas")
	case 3:
		fmt.Println("La conversión es: ", dolares*1365.0, "Wones")
	case 4:
		fmt.Println("La conversión es: ", dolares*0.000012, "BTC")
	default:
		fmt.Println("Opción de moneda no válida")
	}
}
