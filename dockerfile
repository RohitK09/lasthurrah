FROM golang:latest

WORKDIR /lasthurrah

COPY go.mod go.sum ./

ENV GOPROXY direct

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]