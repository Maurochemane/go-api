

# instalar bibliotecas
go mod tidy

# Lib Http
go get github.com/gin-gonic/gin 

# Lib Postgres
_ "github.com/lib/pq"

# Testando a lib

package main 

import "github.com/gin-gonic/gin"

func main(){

	server := gin.Default()
	server.GET("/ping", func (ctx *gin.Contex){
		ctx.JSON(200, gin.H){
			"message":"yes here we go again"
		}
	})

	server.Run(":5000")
}

# Subindo o container postgres
--docker compose up -d go_db

# Mockando um dado

` func (p *productController) GetProducts(ctx *gin.Contex){

	products := []model.product{
		{
			ID: 1,
			Name: "Nanami",
			Price: 200,
		}
	}

	ctx.JSON(http.StatusOk, products)
}`

# Build da imagem

docker build -t go-api .

# Features 

dockerfile com certificados ssl

# --- Stage 1: Build ---
FROM golang:1.20-alpine AS stage1

# Instala certificados CA para permitir conexões HTTPS externas
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Cache de dependências (otimiza o tempo de build)
COPY go.mod go.sum ./
RUN go mod download

# Copia o código e compila
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /executable

# --- Stage 2: Final (Imagem de Produção) ---
FROM scratch

# Copia os certificados SSL/TLS do stage 1
COPY --from=stage1 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copia apenas o binário compilado
COPY --from=stage1 /executable /executable

# Porta que a aplicação utiliza
EXPOSE 8000

# Execução do binário
ENTRYPOINT [ "/executable" ]