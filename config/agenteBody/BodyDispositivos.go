package agentebody

import (
	"github.com/patriciapersi/colabore-api/config/structs"
)

func Dispositivo(nrInsc string) structs.Dispositivos {
	return structs.Dispositivos{
		Cnpj:          "63542443000124",
		DispositivoId: "31e18fe4151b96cb",
		Status:        1,
		ListaEmpresas: []string{nrInsc},
	}
}
