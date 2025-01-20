# Etapa 1: Build da aplicação
FROM golang:1.20 AS builder

# Definir diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar os arquivos do projeto para dentro do contêiner
COPY go.mod go.sum ./
RUN go mod download

# Copiar o restante dos arquivos da aplicação
COPY . .

# Compilar a aplicação Go
RUN go build -o app .

# Etapa 2: Imagem final para execução
FROM alpine:latest

# Adicionar certificado SSL para comunicação HTTPS (caso necessário)
RUN apk --no-cache add ca-certificates

# Definir diretório de trabalho dentro do contêiner
WORKDIR /root/

# Copiar o binário da aplicação da imagem builder para a imagem final
COPY --from=builder /app/app .

# Expor a porta em que a aplicação será executada
EXPOSE 8080

# Comando de entrada para execução da aplicação
CMD ["./app"]
