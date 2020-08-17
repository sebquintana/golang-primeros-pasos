package tp2

import "fmt"

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error.
func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if len(numeros) == 0 {
		return 0, fmt.Errorf("SumarLista Error: lista vacía")
	}

	r := make(chan int, len(numeros)/2)

	total := 0

	for i := 0; i < len(numeros); i += 2 {
		go func(index int) {
			if index+1 >= len(numeros) {
				r <- numeros[index]
			} else {
				r <- sumFunc(numeros[index], numeros[index+1])
			}
		}(i)
	}

	for i := 0; i < len(numeros); i += 2 {
		total += <-r
	}

	return total, nil
}

// SumarLista recibe una function sumadora y una lista
// de numeros. Usando esos parametro retorna la suma de todos
// los numeros. Si la suma no se puede realizar por algun motivo
// se retorna un error. Sin concurrencia
/*func SumarLista(sumFunc sumador, numeros ...int) (int, error) {

	if len(numeros) == 0 {
		return 0, fmt.Errorf("SumarLista Error: lista vacía")
	}

	total := 0

	for i := 0; i < len(numeros); i += 2 {
		if i+1 >= len(numeros) {
			total += numeros[i]
		} else {
			total += sumFunc(numeros[i], numeros[i+1])
		}
	}

	return total, nil
}*/
