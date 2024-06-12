package main

import (
	"fmt"

	// O import deve levar em consideração o que está definido no go.mod adicionando o nome do pacote a ser utilizado
	"github.com/PCPedroso/pos-fc-pacotes/math"
	"github.com/google/uuid"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Soma())
	fmt.Println(uuid.New())
}
