FROM golang:1.26

WORKDIR /go/src/app

COPY . . 

EXPOSE 5000

RUN go build -o main cmd/main.go

CMD [ "./main" ]