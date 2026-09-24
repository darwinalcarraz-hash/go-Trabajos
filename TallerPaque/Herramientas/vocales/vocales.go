package vocales

import "fmt"

func ContarVocales(frase string) {
	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, letra := range frase {
		switch letra {
		case 'a', 'A':
			a++
		case 'e', 'E':
			e++
		case 'i', 'I':
			i++
		case 'o', 'O':
			o++
		case 'u', 'U':
			u++
		}
	}

	fmt.Println("Vocal A: ", a)
	fmt.Println("Vocal E: ", e)
	fmt.Println("Vocal I: ", i)
	fmt.Println("Vocal O: ", o)
	fmt.Println("Vocal U: ", u)
}
