# Comandos para compilação da aplicação Go
.PHONY: build
build:
	@echo "🏗️  Compilando a aplicação..."
	GOOS=linux GOARCH=amd64 go build -o 
	@echo "✅ Aplicação compilada com sucesso!"

# Comando para rodar a aplicação localmente
.PHONY: run
run:
	@echo "🚀 Executando a aplicação..."
	go run .

# Comando para rodar via Docker Compose (banco de dados + aplicação)
.PHONY: up
up:
	@echo "🚀 Subindo serviços com Docker Compose..."
	docker-compose up -d
	@echo "✅ Serviços estão rodando!"

# Comando para parar os serviços Docker
.PHONY: down
down:
	@echo "🛑 Parando serviços..."
	docker-compose down
	@echo "✅ Serviços parados!"

# Construção da imagem Docker
.PHONY: docker-build
docker-build:
	@echo "🐳 Construindo imagem Docker..."
	docker build .
	@echo "✅ Imagem Docker criada com sucesso!"

# Rodar todos os testes da aplicação
.PHONY: test
test:
	@echo "🧪 Rodando todos os testes..."
	go test ./... -v
	@echo "✅ Todos os testes foram executados!"

# Rodar testes específicos (usar com make test-specific TEST=./package/name)
.PHONY: test-specific
test-specific:
	@echo "🧪 Rodando testes específicos em $(TEST)..."
	go test $(TEST) -v
	@echo "✅ Testes específicos foram executados!"
