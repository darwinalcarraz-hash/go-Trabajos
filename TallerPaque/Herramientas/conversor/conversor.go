package conversor

import "fmt"

func ConvertirMoneda(dolares float64, moneda string) {
	const tasaEuro = 0.88
	const tasaLibra = 0.75
	const tasaWon = 1365.0
	const tasaBTC = 0.000012

	switch moneda {
	case "Euros", "euros", "EURO":
		fmt.Println("La conversión es: ", dolares*tasaEuro)
	case "LB", "lb", "Libras":
		fmt.Println("La conversión es: ", dolares*tasaLibra)
	case "Won", "won", "WON":
		fmt.Println("La conversión es: ", dolares*tasaWon)
	case "BTC", "btc":
		fmt.Println("La conversión es: ", dolares*tasaBTC)
	default:
		fmt.Println("Moneda no válida")
	}
}
