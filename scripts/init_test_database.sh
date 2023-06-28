#!/bin/bash

# Inicia o banco de dados de teste
docker-compose -f docker-compose.yaml up -d

# Executa os testes
go test -coverprofile=coverage.out ./...

# # Extrai as métricas de cobertura
# coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Para o banco de dados de teste
docker-compose -f docker-compose.yaml down

# # Verifica se a cobertura é menor que 60%
# if [ $(echo "$coverage < 60" | bc -l) -eq 1 ]; then
#   echo "FAILED"
#   echo "A cobertura de teste é menor que 60%: $coverage%"
#   exit 1
# fi

# echo "A cobertura de teste é de $coverage%"


