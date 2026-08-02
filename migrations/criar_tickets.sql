BEGIN;

CREATE TABLE tickets (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,

    titulo VARCHAR(150) NOT NULL,
    descricao TEXT NOT NULL,
    solicitante VARCHAR(120) NOT NULL,

    responsavel VARCHAR(120),

    categoria VARCHAR(80) NOT NULL,
    prioridade VARCHAR(30) NOT NULL,

    status VARCHAR(30) NOT NULL DEFAULT 'ABERTO',

    solucao TEXT,

    data_criacao TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data_ultima_alteracao TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data_resolucao TIMESTAMPTZ,

    CONSTRAINT tickets_titulo_nao_vazio
        CHECK (BTRIM(titulo) <> ''),

    CONSTRAINT tickets_descricao_nao_vazia
        CHECK (BTRIM(descricao) <> ''),

    CONSTRAINT tickets_solicitante_nao_vazio
        CHECK (BTRIM(solicitante) <> ''),

    CONSTRAINT tickets_categoria_nao_vazia
        CHECK (BTRIM(categoria) <> ''),

    CONSTRAINT tickets_prioridade_nao_vazia
        CHECK (BTRIM(prioridade) <> ''),

    CONSTRAINT tickets_status_nao_vazio
        CHECK (BTRIM(status) <> ''),

    CONSTRAINT tickets_responsavel_nao_vazio
        CHECK (
            responsavel IS NULL
            OR BTRIM(responsavel) <> ''
        ),

    CONSTRAINT tickets_solucao_nao_vazia
        CHECK (
            solucao IS NULL
            OR BTRIM(solucao) <> ''
        )
);

CREATE OR REPLACE FUNCTION atualizar_data_ticket()
RETURNS TRIGGER AS $$
BEGIN
    NEW.data_ultima_alteracao = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_atualizar_data_ticket
BEFORE UPDATE ON tickets
FOR EACH ROW
EXECUTE FUNCTION atualizar_data_ticket();

COMMIT;