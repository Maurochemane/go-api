#Dockerfile padrão


# FROM golang:1.25

# set working directory
# WORKDIR /go/src/app

# Copy the source code
# COPY . . 

#EXPOSE the port
# EXPOSE 8000

# Build the Go app
# RUN go build -o main cmd/main.go

# Run the executable
# CMD ["./main"]

#------------------------------------------------------------------------------

    #--------------------------------------------
#Aplicar o dockerfile optimizado 


#Step 1: Baixar e compilar o binário
FROM golang:1.25 as stage1


WORKDIR /app

# Copia o go.mod e faz download das dependencia
COPY go.mod go.sum ./
RUN go mod download

# Copia o codigo da aplicação e compila o binário
COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o executable ./cmd

####stage 2 

# Copiar o binário do stage 1 para a imagem final.
FROM  scratch

# Copia apenas o binário.
COPY --from=stage1 /app/executable /
EXPOSE 8000

# Define o ponto de entrada para o container como /executavel.
# O binário será executado quando o container for iniciado.
ENTRYPOINT [ "/executable" ]

