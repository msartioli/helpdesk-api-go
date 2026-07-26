package main

import (
	"time"
)

type StatusChamado string

const (
	StatusAberto        StatusChamado = "Aberto"
	StatusEmAtendimento StatusChamado = "Em atendimento"
	StatusResolvido     StatusChamado = "Resolvido"
	StatusCancelado     StatusChamado = "Cancelado"
)

type Chamado struct {
	NumeroDoChamado       int
	Titulo                string
	Descricao             string
	Solicitante           string
	Responsavel           string
	Categoria             string
	Prioridade            string
	Status                StatusChamado
	Solucao               string
	DataDeCriacao         time.Time
	DataDaUltimaAlteracao time.Time
	DataDeResolucao       time.Time
}
