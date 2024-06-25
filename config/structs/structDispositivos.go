package structs

type Dispositivos struct {
	Cnpj          string `json:"Cnpj"`
	DispositivoId string `json:"DispositivoId"`
	Status        int    `json:"Status"`
	ListaEmpresas string `json:"ListaEmpresas"`
}
