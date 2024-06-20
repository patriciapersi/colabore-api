package testUtil

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
)

func SetupEnv(envs map[string]string) {
	for key, value := range envs {
		os.Setenv(key, value)
	}
}

func LoadEnv() error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("erro ao obter o diretório de trabalho atual: %v", err)
	}

	// Vai subindo nos diretórios até encontrar o .env
	for {
		envPath := filepath.Join(cwd, ".env")
		if _, err := os.Stat(envPath); err == nil {
			// Carrega as variáveis de ambiente do arquivo .env
			err := godotenv.Load(envPath)
			if err != nil {
				return fmt.Errorf("erro ao carregar o arquivo .env: %v", err)
			}
			return nil
		}

		// Se não encontrou o arquivo .env, tenta o diretório pai
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			return fmt.Errorf("arquivo .env não encontrado")
		}
		cwd = parentDir
	}
}

func GetFirstDayOfMonth() string {
	// Obter a data atual
	now := time.Now()

	// Construir a data do primeiro dia do mês atual
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Formatando a data no formato desejado (DD/MM/YYYY)
	formattedDate := firstOfMonth.Format("02/01/2006")

	return formattedDate
}

func GenerateCPF() string {
	var digits []int
	for i := 0; i < 9; i++ {
		digits = append(digits, rand.Intn(10))
	}
	sum := 0
	for i := 0; i < 9; i++ {
		sum += digits[i] * (10 - i)
	}
	remainder := sum % 11
	var digit1 int
	if remainder < 2 {
		digit1 = 0
	} else {
		digit1 = 11 - remainder
	}
	digits = append(digits, digit1)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += digits[i] * (11 - i)
	}
	remainder = sum % 11
	var digit2 int
	if remainder < 2 {
		digit2 = 0
	} else {
		digit2 = 11 - remainder
	}
	digits = append(digits, digit2)

	cpf := fmt.Sprintf("%d%d%d%d%d%d%d%d%d%d%d", digits[0], digits[1], digits[2], digits[3], digits[4], digits[5], digits[6], digits[7], digits[8], digits[9], digits[10])

	return cpf
}
