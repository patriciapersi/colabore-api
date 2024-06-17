package config

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/go-resty/resty/v2"
	testutil "github.com/patriciapersi/colabore-api/util"
)

type API struct {
	Client          *resty.Client
	EndpointsAgente map[string]string
	EndpointsApp    map[string]string
}

func init() {
	testutil.LoadEnv()
}

func SetupHeadersAgente() map[string]string {
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("X_API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}

func SetupHeadersApp() map[string]string {
	token, _ := ReturnTokenId()

	return map[string]string{
		"Content-Type": "application/json",
		"awsauthtoken": token,
	}
}

func SetupApi() *API {

	api := resty.New().
		SetBaseURL("https://fortescolabore2.fortestecnologia.com.br/homolog-next")

	endpointsAgente := map[string]string{
		"Abono":                              "/api/agente/Abono",
		"ReverterAbono":                      "/api/agente/Abono/Redefinicao",
		"ImagemReferencia":                   "/api/agente/Arquivo/ImagemReferencia",             //NAO SERÁ AUTOMATIZADO
		"ArquivoRelatorioPontoURLTemporaria": "/api/agente/Arquivo/RelatorioPonto/URLTemporaria", //NAO SERÁ AUTOMATIZADO
		"AssinaturaLiberar":                  "/api/agente/Assinatura/Liberar",
		"Assinatura":                         "/api/agente/Assinatura",
		"AssinaturaContestacao":              "/api/agente/Assinatura/Contestacao",
		"AssinaturaContestacaoStatus":        "/api/agente/Assinatura/Contestacao/Status",
		"AssinaturaRelatorioPonto":           "/api/agente/Assinatura/RelatorioPonto",
		"AssinaturaDigital":                  "/api/agente/AssinaturaDigital",          //NAO SERÁ AUTOMATIZADO
		"AssinaturaDigitalFormData":          "/api/agente/AssinaturaDigital/FormData", //NAO SERÁ AUTOMATIZADO
		"Candidato":                          "/api/agente/Candidato",
		"CandidatoRetificar":                 "/api/agente/Candidato/Retificar",
		"Acesso":                             "/api/agente-acesso", //NAO SERÁ AUTOMATIZADO
		"Agente":                             "/api/agente",        //NAO SERÁ AUTOMATIZADO
		"Colaborador":                        "/api/agente/Colaborador",
		"ColaboradorAtivacao":                "/api/agente/Colaborador/Ativacao",
		"ColaboradorFreemium":                "/api/agente/Colaborador/Freemium",
		"ColaboradorPreemium":                "/api/agente/Colaborador/Premium",
		"ColaboradorEnderecoConfirmacao":     "/api/agente/Colaborador/Endereco/Confirmacao",
		"ColaboradorEnderecos":               "/api/agente/Colaborador/Enderecos",
		"ColaboradorDefinicoes":              "/api/agente/Colaborador/Definicoes",
		"Declaracao":                         "/api/agente/Declaracao",
		"Dispositivos":                       "/api/agente/Dispositivos",
		"DispositivosStatus":                 "/api/agente/Dispositivos/Status",
		"FeriasInformacoes":                  "/api/agente/Ferias/Informacoes",
		"Ferias":                             "/api/agente/Ferias",
		"FolhaDePagamento":                   "/api/agente/FolhaDePagamento",
		"Gestor":                             "/api/agente/Gestor",
		"GestorRH":                           "/api/agente/Gestor/RH",
		"GETimagem":                          "/api/agente/Imagem", //NAO SERÁ AUTOMATIZADO
		"LicencaReconhecimentoFacial":        "/api/agente/Licenca/ReconhecimentoFacial",
		"LicenciadoEmailParaNotificacoesDaEmpresa": "/api/agente/Licenciado/EmailParaNotificacoesDaEmpresa", //NAO SERÁ AUTOMATIZADO
		"LicenciadoLogo":                       "/api/agente/Licenciado/Logo", //NAO SERÁ AUTOMATIZADO
		"LicenciadoDefinicoes":                 "/api/agente/Licenciado/Definicoes",
		"Mensagem":                             "/api/agente/Mensagem",
		"Numerador":                            "/api/agente/Numerador", //NAO SERÁ AUTOMATIZADO
		"Pesquisa":                             "/api/agente/Pesquisa",
		"PesquisaRespostas":                    "/api/agente/Pesquisa/Respostas",
		"PesquisaRespostasPorPagina":           "/api/agente/Pesquisa/RespostasPorPagina",
		"PesquisaResposta":                     "/api/agente/Pesquisa/Resposta",
		"PontoBatidas":                         "/api/agente/Ponto/Batidas",
		"PontoProcessamentoRedefinicao":        "/api/agente/Ponto/Processamento/Redefinicao",
		"PontoAssinaturaRetroativaSolicitacao": "/api/agente/Ponto/AssinaturaRetroativa/Solicitacao", //NAO SERÁ AUTOMATIZADO
		"PontoObterTimeZone":                   "/api/agente/Ponto/ObterTimeZone",
		"Termo":                                "/api/agente/Termo", //NAO SERÁ AUTOMATIZADO
		"VersaoApp":                            "/api/agente/VersaoApp",
	}

	endpointsApp := map[string]string{
		"Appferias":      "/api/app/Ferias",
		"AbonoHistorico": "/api/app/Abono/Historico",
		"Folha":          "/api/app/FolhaDePagamento",
		"Mensagem":       "/api/app/Mensagem",
		"MensagemLida":   "/api/app/Mensagem/Ler",
	}

	return &API{
		Client:          api,
		EndpointsAgente: endpointsAgente,
		EndpointsApp:    endpointsApp,
	}

}

func ReturnTokenId() (string, error) {
	clientID := os.Getenv("CLIENT_ID")
	cpf := os.Getenv("CPF")
	password := os.Getenv("PASSWORD")

	// Configuração da sessão AWS
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})
	if err != nil {
		return "", fmt.Errorf("erro ao criar sessão AWS: %v", err)
	}

	// Criando cliente Cognito Identity Provider
	cognitoClient := cognitoidentityprovider.New(sess)

	// Autenticação do usuário
	authParams := map[string]*string{
		"USERNAME": aws.String(cpf),
		"PASSWORD": aws.String(password),
	}

	authInput := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: authParams,
		ClientId:       aws.String(clientID),
	}

	authOutput, err := cognitoClient.InitiateAuth(authInput)
	if err != nil {
		return "", fmt.Errorf("erro ao iniciar autenticação: %v", err)
	}

	// Obtendo tokens do resultado
	idToken := aws.StringValue(authOutput.AuthenticationResult.IdToken)

	return idToken, nil
}
