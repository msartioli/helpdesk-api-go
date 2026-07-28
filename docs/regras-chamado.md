# Regras de Negócio do Chamado

## Objetivo

Este documento descreve o modelo de dados e as regras de negócio de um chamado no sistema de Helpdesk, cobrindo desde a abertura até os estados finais, incluindo as transições de status permitidas.

## Informações do chamado

Todo chamado é composto pelos seguintes campos:

| Campo | Descrição |
|---|---|
| Número de identificação | Identificador único do chamado, gerado pelo sistema |
| Título | Resumo curto do problema relatado |
| Descrição | Detalhamento do problema |
| Solicitante | Pessoa que abriu o chamado |
| Responsável | Pessoa encarregada do atendimento |
| Categoria | Classificação do tipo de problema |
| Prioridade | Nível de urgência do chamado |
| Status | Situação atual do chamado |
| Solução | Descrição de como o problema foi resolvido |
| Data de criação | Data em que o chamado foi aberto |
| Data da última alteração | Data da modificação mais recente |
| Data de resolução | Data em que o chamado foi resolvido |

## Abertura do chamado

### Campos obrigatórios

Para abrir um chamado, os campos abaixo precisam ser preenchidos:

- título;
- descrição;
- solicitante;
- categoria;
- prioridade.

Se qualquer um desses campos estiver vazio, o chamado não pode ser criado.

### Campos definidos automaticamente

No momento da criação, o sistema é responsável por:

- gerar o número de identificação;
- definir o status inicial como `Aberto`;
- registrar a data de criação;
- registrar a data da última alteração.

Os campos **responsável**, **solução** e **data de resolução** não são obrigatórios na abertura e permanecem em branco até que o chamado avance no fluxo de atendimento.

## Prioridades

Um chamado pode ter uma das seguintes prioridades:

- `Baixa`
- `Média`
- `Alta`
- `Crítica`

## Status

Um chamado pode estar em um dos seguintes status:

- `Aberto`
- `Em atendimento`
- `Resolvido`
- `Cancelado`

## Fluxo de status

As transições de status permitidas são:

| Status atual | Pode mudar para |
|---|---|
| Aberto | Em atendimento |
| Aberto | Cancelado |
| Em atendimento | Resolvido |
| Em atendimento | Cancelado |

Qualquer transição que não esteja listada na tabela acima é considerada inválida e não deve ser permitida pelo sistema.

## Início do atendimento

Ao mover um chamado de `Aberto` para `Em atendimento`, é necessário:

- informar o responsável pelo atendimento;
- atualizar a data da última alteração.

Um chamado não pode passar para `Em atendimento` sem um responsável definido.

## Resolução do chamado

Ao mover um chamado de `Em atendimento` para `Resolvido`, é necessário:

- informar a solução aplicada;
- registrar a data de resolução;
- atualizar a data da última alteração.

Um chamado não pode ser marcado como `Resolvido` sem que uma solução tenha sido informada.

## Cancelamento do chamado

Um chamado só pode ser cancelado se estiver nos status:

- `Aberto`;
- `Em atendimento`.

Ao cancelar um chamado, a data da última alteração deve ser atualizada.

## Estados finais

`Resolvido` e `Cancelado` são considerados estados finais. Uma vez que um chamado atinge um desses status, ele não pode mais ser alterado.

## Exclusão de chamados

Chamados não são excluídos do sistema. O histórico deve ser mantido integralmente, mesmo para chamados já resolvidos ou cancelados.