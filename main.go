package main

import (
	"fmt"
	"strconv"
)

func limparTela() {
	// \033[H posiciona o cursor no canto superior esquerdo
	// \033[2J limpa toda a tela do terminal
	fmt.Print("\033[H\033[2J")
}

func mostrarCalculo(rangeValores []string, rangeOperadores []string) {
	var calculo string

	for i := range rangeValores {
		calculo += rangeValores[i]

		if i < len(rangeOperadores) {
            calculo += " " + rangeOperadores[i] + " "
        }
	}
	fmt.Println(calculo)
}

func main() {
	var valor string
	var operador string
	var valores []string
	var operadores []string

	fmt.Println("Digite algum valor para começar!")
	for true {
		_, err := fmt.Scan(&valor)
		if err != nil {
			return
		}

		_, err = strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println("O valor digitado não é numérico!")
			return
		}

		valores = append(valores, valor)

		limparTela()
		mostrarCalculo(valores, operadores)

		fmt.Println("Agora digite um operador(+ - / * =):")

		_, err = fmt.Scan(&operador)
		if err != nil {
			return
		}

		if operador == "=" {
			break
		} else if operador != "+" && operador != "-" && operador != "/" && operador != "*" {
			fmt.Println("Operador digitado errado!")
			return
		}

		operadores = append(operadores, operador)

		limparTela()
		mostrarCalculo(valores, operadores)
	}
}