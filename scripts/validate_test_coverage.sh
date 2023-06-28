#!/bin/bash

# Executa os testes e gera o perfil de cobertura
go test -coverprofile=coverage.out .././...

# Extrai as métricas de cobertura
coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')

# Verifica se a cobertura é menor que 50%
if [ $(echo "$coverage < 50" | bc -l) -eq 1 ]; then
  echo "A cobertura de teste é menor que 50%: $coverage%"
  exit 1
fi

echo "A cobertura de teste é de $coverage%"