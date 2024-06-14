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

func SetupHeadersAgente() map[string]string {
	testutil.LoadEnv()
	return map[string]string{
		"Content-Type":                  "application/json",
		"x-api-key":                     os.Getenv("X_API_KEY"),
		"cnpj-licenciado":               os.Getenv("CNPJ_LICENCIADO"),
		"colabore-signature-expiration": os.Getenv("COLABORE_SIGNATURE_EXPIRATION"),
		"colabore-signature":            os.Getenv("COLABORE_SIGNATURE"),
	}
}

func SetupHeadersApp() map[string]string {
	testutil.LoadEnv()

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
		"ImagemReferencia":                         "/api/agente/Arquivo/ImagemReferencia",
		"ArquivoRelatorioPontoURLTemporaria":       "/api/agente/Arquivo/RelatorioPonto/URLTemporaria",
		"AssinaturaLiberar":                        "/api/agente/Assinatura/Liberar",
		"Assinatura":                               "/api/agente/Assinatura",
		"AssinaturaContestacao":                    "/api/agente/Assinatura/Contestacao",
		"AssinaturaContestacaoStatus":              "/api/agente/Assinatura/Contestacao/Status",
		"AssinaturaRelatorioPonto":                 "/api/agente/Assinatura/RelatorioPonto",
		"AssinaturaDigital":                        "/api/agente/AssinaturaDigital",
		"AssinaturaDigitalFormData":                "/api/agente/AssinaturaDigital/FormData",
		"Candidato":                                "/api/agente/Candidato",
		"CandidatoRetificar":                       "/api/agente/Candidato/Retificar",
		"Acesso":                                   "/api/agente-acesso", //NAO SERÁ AUTOMATIZADO
		"Agente":                                   "/api/agente",        //NAO SERÁ AUTOMATIZADO
		"Colaborador":                              "/api/agente/Colaborador",
		"ColaboradorAtivacao":                      "/api/agente/Colaborador/Ativacao",
		"ColaboradorFreemium":                      "/api/agente/Colaborador/Freemium",
		"ColaboradorPreemium":                      "/api/agente/Colaborador/Premium",
		"ColaboradorEnderecoConfirmacao":           "/api/agente/Colaborador/Endereco/Confirmacao",
		"ColaboradorEnderecos":                     "/api/agente/Colaborador/Enderecos",
		"ColaboradorDefinicoes":                    "/api/agente/Colaborador/Definicoes",
		"Declaracao":                               "/api/agente/Declaracao",
		"Dispositivos":                             "/api/agente/Dispositivos",
		"DispositivosStatus":                       "/api/agente/Dispositivos/Status",
		"FeriasInformacoes":                        "/api/agente/Ferias/Informacoes",
		"Ferias":                                   "/api/agente/Ferias",
		"FolhaDePagamento":                         "/api/agente/FolhaDePagamento",
		"Gestor":                                   "/api/agente/Gestor",
		"GestorRH":                                 "/api/agente/Gestor/RH",
		"GETimagem":                                "/api/agente/Imagem",
		"LicencaReconhecimentoFacial":              "/api/agente/Licenca/ReconhecimentoFacial",
		"LicenciadoEmailParaNotificacoesDaEmpresa": "/api/agente/Licenciado/EmailParaNotificacoesDaEmpresa",
		"LicenciadoLogo":                           "/api/agente/Licenciado/Logo",
		"LicenciadoDefinicoes":                     "/api/agente/Licenciado/Definicoes",
		"Mensagem":                                 "/api/agente/Mensagem",
		"Numerador":                                "/api/agente/Numerador",
		"Pesquisa":                                 "/api/agente/Pesquisa",
		"PesquisaRespostas":                        "/api/agente/Pesquisa/Respostas",
		"PesquisaRespostasPorPagina":               "/api/agente/Pesquisa/RespostasPorPagina",
		"PesquisaResposta":                         "/api/agente/Pesquisa/Resposta",
		"PontoBatidas":                             "/api/agente/Ponto/Batidas",
		"PontoProcessamentoRedefinicao":            "/api/agente/Ponto/Processamento/Redefinicao",
		"PontoAssinaturaRetroativaSolicitacao":     "/api/agente/Ponto/AssinaturaRetroativa/Solicitacao",
		"PontoObterTimeZone":                       "/api/agente/Ponto/ObterTimeZone",
		"Termo":                                    "/api/agente/Termo",
		"VersaoApp":                                "/api/agente/VersaoApp",
	}

	endpointsApp := map[string]string{
		"Appferias": "/api/app/Ferias",
	}

	return &API{
		Client:          api,
		EndpointsAgente: endpointsAgente,
		EndpointsApp:    endpointsApp,
	}

}

func ReturnTokenId() (string, error) {
	clientID := "1d339letjh0ndq27pm3e9gvfko"
	username := "60515860409"
	password := "12345678"

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
		"USERNAME": aws.String(username),
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
