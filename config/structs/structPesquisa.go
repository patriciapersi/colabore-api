package structs

type TipoQuestao string

const (
	SUBJETIVA TipoQuestao = "SUBJETIVA"
)

type Colaboradores struct {
	Matricula        string `json:"Matricula"`
	CPF              string `json:"CPF"`
	NrInscEmpregador string `json:"NrInscEmpregador"`
}

type Pergunta struct {
	ID          string                   `json:"id"`
	Tipo        TipoQuestao              `json:"tipo"`
	Ordem       string                   `json:"ordem"`
	Texto       string                   `json:"texto"`
	Obrigatoria bool                     `json:"obrigatoria"`
	NotaMinima  int                      `json:"notaMinima"`
	NotaMaxima  int                      `json:"notaMaxima"`
	Respostas   []map[string]interface{} `json:"respostas"`
}

type Pesquisa struct {
	ID                 string          `json:"id"`
	Inicio             string          `json:"inicio"`
	Fim                string          `json:"fim"`
	NrInscEmpregador   string          `json:"NrInscEmpregador"`
	Titulo             string          `json:"titulo"`
	MonitoramentoSaude bool            `json:"monitoramentoSaude"`
	PesquisaAnonima    bool            `json:"pesquisaAnonima"`
	IndependeMatricula bool            `json:"independeMatricula"`
	Versao             string          `json:"Versao"`
	Perguntas          []Pergunta      `json:"perguntas"`
	Colaboradores      []Colaboradores `json:"colaboradores"`
}

//STRUCT DE DELEÇÃO DE PESQUISA

type PesquisaDeletada struct {
	Matricula        string `json:"Matricula"`
	CPF              string `json:"CPF"`
	NrInscEmpregador string `json:"NrInscEmpregador"`
	PesquisaId       string `json:"PesquisaId"`
}

type DeleteRequest struct {
	Pesquisas []PesquisaDeletada `json:"pesquisas"`
}
