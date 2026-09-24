package operaciones

func Suma(a, b int) int {
	return a + b
}

func SumaYResta(a, b int) (int, int) {
	sumaTotal := a + b
	var restaResultado int

	if b > a {
		restaResultado = b - a
	} else {
		restaResultado = 0
	}

	return sumaTotal, restaResultado
}

//Funciones Varidicas

func Sumatoria(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
