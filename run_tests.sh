#!/bin/bash

echo "Rodando testes..."
go clean -cache
go test ./tests/...

if [ $? -ne 0 ]; then
    echo "Erro: Os testes falharam. Cancelando o push."
    exit 1
else
    echo "Testes passaram. Continuando com o git push."
fi
