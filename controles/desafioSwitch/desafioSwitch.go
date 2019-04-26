package main

import "fmt"

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch {
	case nota >= 9 && n <= 10:
		return "A"
	case nota >= 8 && n < 9:
		return "B"
	case nota >= 5 && n < 8:
		return "C"
	case nota >= 3 && n < 5:
		return "D"
	case nota >= 0 && n < 3:
		return "E"
	default:
		return "Nota inválida"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.9))
	fmt.Println(notaParaConceito(2.1))
	fmt.Println(notaParaConceito(11.1))

}
