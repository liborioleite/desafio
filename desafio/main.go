package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Faturamento struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {

	fmt.Println("OBS: Eu deixei as funções que correspondem aos desafios comentados, é so descomentar aqui na função main e rodar no pront o comando go run main.go")

	// Desafio1()
	// Desafio2()
	// Desafio3()
	// Desafio4()
	// Desafio5()
}

func Desafio1() {
	indice := 13
	soma := 0
	k := 0

	for k < indice {
		k++
		soma += k
	}

	fmt.Println("O valor da soma é:", soma)
}

func Desafio2() {
	// 	2) Dado a sequência de Fibonacci, onde se inicia por 0 e 1 e o próximo valor sempre será a soma dos 2 valores anteriores
	// (exemplo: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34...), escreva um programa na linguagem que desejar onde, informado um número, ele calcule a sequência de Fibonacci e
	// retorne uma mensagem avisando se o número informado pertence ou não a sequência.
	// IMPORTANTE: Esse número pode ser informado através de qualquer entrada de sua preferência ou pode ser previamente definido no código;

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanln(&numero)

	a, b := 0, 1
	sequencia := []int{0}
	pertence := false

	if numero == 0 {
		pertence = true
	} else {
		for b <= numero {
			sequencia = append(sequencia, b)
			if b == numero {
				pertence = true
			}
			a, b = b, a+b
		}
	}

	fmt.Println("Calculo da sequencia:", sequencia)

	if pertence {
		fmt.Printf("O número %d pertence à sequência.\n", numero)
	} else {
		fmt.Printf("O número %d NÃO pertence à sequência.\n", numero)
	}

}

func Desafio3() {
	// 3) Dado um vetor que guarda o valor de faturamento diário de uma distribuidora, faça um programa, na linguagem que desejar, que calcule e retorne:
	// • O menor valor de faturamento ocorrido em um dia do mês;
	// • O maior valor de faturamento ocorrido em um dia do mês;
	// • Número de dias no mês em que o valor de faturamento diário foi superior à média mensal.

	// IMPORTANTE:
	// a) Usar o json ou xml disponível como fonte dos dados do faturamento mensal;
	// b) Podem existir dias sem faturamento, como nos finais de semana e feriados. Estes dias devem ser ignorados no cálculo da média;

	file, err := os.Open("arquivo1.json")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	var dados []Faturamento
	if err := json.NewDecoder(file).Decode(&dados); err != nil {
		fmt.Println("Erro ao decodificar JSON:", err)
		return
	}

	var soma float64
	var diasComFaturamento int
	var menor, maior float64
	menor = -1

	for _, dia := range dados {
		if dia.Valor > 0 {
			soma += dia.Valor
			diasComFaturamento++

			if menor == -1 || dia.Valor < menor {
				menor = dia.Valor
			}
			if dia.Valor > maior {
				maior = dia.Valor
			}
		}
	}

	media := soma / float64(diasComFaturamento)

	diasAcimaDaMedia := 0
	for _, dia := range dados {
		if dia.Valor > media {
			diasAcimaDaMedia++
		}
	}

	fmt.Printf("Menor faturamento: R$ %.2f\n", menor)
	fmt.Printf("Maior faturamento: R$ %.2f\n", maior)
	fmt.Printf("Dias com faturamento acima da média mensal (%.2f): %d\n", media, diasAcimaDaMedia)

}

func Desafio4() {
	// 4) Dado o valor de faturamento mensal de uma distribuidora, detalhado por estado:
	// • SP – R$67.836,43
	// • RJ – R$36.678,66
	// • MG – R$29.229,88
	// • ES – R$27.165,48
	// • Outros – R$19.849,53

	// Escreva um programa na linguagem que desejar onde calcule o percentual de representação que cada estado teve
	// dentro do valor total mensal da distribuidora.

	faturamento := map[string]float64{
		"SP":     67836.43,
		"RJ":     36678.66,
		"MG":     29229.88,
		"ES":     27165.48,
		"Outros": 19849.53,
	}

	var total float64
	for _, valor := range faturamento {
		total += valor
	}

	fmt.Printf("Faturamento total: R$ %.2f\n\n", total)

	for estado, valor := range faturamento {
		percentual := (valor / total) * 100
		fmt.Printf("%s: R$ %.2f (%.2f%%)\n", estado, valor, percentual)
	}
}

func Desafio5() {
	// 	5) Escreva um programa que inverta os caracteres de um string.

	// IMPORTANTE:
	// a) Essa string pode ser informada através de qualquer entrada de sua preferência ou pode ser previamente definida no código;
	// b) Evite usar funções prontas, como, por exemplo, reverse;

	var texto string
	fmt.Print("String: ")
	fmt.Scanln(&texto)

	caractere := []rune(texto)

	for i, j := 0, len(caractere)-1; i < j; i, j = i+1, j-1 {
		caractere[i], caractere[j] = caractere[j], caractere[i]
	}

	invertida := string(caractere)
	fmt.Println("String invertida:", invertida)

}
