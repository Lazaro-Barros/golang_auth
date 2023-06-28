# Estágio de compilação
FROM golang:1.20-alpine AS build

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos do código-fonte para o contêiner
COPY . .

# Instala as dependências do projeto
RUN go mod download

# # Executando testes e validando coverage
# RUN chmod +x scripts/init_test_database.sh
# RUN sh scripts/init_test_database.sh

# Compila o código do Go
RUN go build cmd/main.go

# # Estágio de produção
# FROM alpine:latest

# # Define o diretório de trabalho
# WORKDIR /app

# # Copia o executável do estágio de compilação
# COPY --from=build /app/main .

# Define o comando de inicialização
CMD ["./main"]