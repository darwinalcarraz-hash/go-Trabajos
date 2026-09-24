package vocales

import "fmt"

func ContarVocales(frase string) {
	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, letra := range frase {
		switch letra {
		case 'a', 'A', 'á', 'Á', 'ä', 'Ä':
			a++
		case 'e', 'E', 'é', 'É', 'ë', 'Ë':
			e++
		case 'i', 'I', 'í', 'Í', 'ï', 'Ï':
			i++
		case 'o', 'O', 'ó', 'Ó', 'ö', 'Ö':
			o++
		case 'u', 'U', 'ú', 'Ú', 'ü', 'Ü':
			u++
		}
	}

	fmt.Println("Vocal A: ", a)
	fmt.Println("Vocal E: ", e)
	fmt.Println("Vocal I: ", i)
	fmt.Println("Vocal O: ", o)
	fmt.Println("Vocal U: ", u)
}
