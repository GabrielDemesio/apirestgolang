# Usar a imagem base do Go
FROM golang:1.22 AS builder

# Definir o diretório de trabalho
WORKDIR /api-go-rest

# Copiar go.mod e go.sum para o contêiner
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o restante do código
COPY . .

# Compilar o aplicativo
RUN go build -o main .

# Usar uma imagem menor para executar o aplicativo
FROM alpine:latest

# Copiar o binário do builder
COPY --from=builder /api-go-rest/main .

# Expor a porta que o aplicativo usará
EXPOSE 8000

# Comando para executar o aplicativo
CMD ["./main"]
