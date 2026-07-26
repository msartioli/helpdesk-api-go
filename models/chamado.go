package main

import (
	"fmt"
	"time"
)

type Chamado struct {
	s
	NumeroDoChamado       int
	Titulo                string
	Descricao             string
	Solicitante           string
	Responsavel           string
	Categoria             string
	Prioridade            string
	Status                string
	Solucao               string
	DataDeCriacao         time.Time
	DataDaUltimaAlteracao time.Time
	DataDeResolucao       time.Time
}

func VerStruct() {
	fmt.Println("Hello World")
	return
}
