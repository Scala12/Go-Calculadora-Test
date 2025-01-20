package main

import "fmt"

func main() {
	x := soma(1, 2, 3)
	y := multiplica(10, 10)
	w := subtrai(10, 5)
	z := divide(20, 2)

	fmt.Println("Soma:", x, "Multiplicação:", y, "Subtração:", w, "Divisão:", z)
}

func soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func subtrai(i ...int) int {
	if len(i) < 2 {
		return 0 // Evita subtração com números insuficientes
	}
	total := i[0]
	for _, v := range i[1:] {
		total -= v
	}
	return total
}

func multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total *= v
	}
	return total
}

func divide(i ...int) int {
	if len(i) < 2 {
		return 0 // Evita divisão com números insuficientes
	}
	total := i[0]
	for _, v := range i[1:] {
		if v == 0 {
			return 0 // Evita divisão por zero
		}
		total /= v
	}
	return total
}
