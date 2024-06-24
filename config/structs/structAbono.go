package structs

type StatusSolicitacao string

const (
	PENDENTE  StatusSolicitacao = "1"
	ACEITO    StatusSolicitacao = "2"
	REJEITADO StatusSolicitacao = "3"
)

// Define uma struct para representar o Abono
type Abono struct {
	NrInscEmpregador string            `json:"NrInscEmpregador"`
	Evento           string            `json:"Evento"`
	CPF              string            `json:"CPF"`
	Matricula        string            `json:"Matricula"`
	Nome             string            `json:"Nome"`
	DataAbono        string            `json:"DataAbono"`
	MotivoId         string            `json:"MotivoId"`
	StatusSol        StatusSolicitacao `json:"StatusSol"`
	Turnos           []string          `json:"Turnos"`
}

// Define uma struct para representar o corpo da solicitação
type PostAbonoBody struct {
	Abonos []Abono `json:"Abonos"`
}
