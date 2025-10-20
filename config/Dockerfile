FROM golang:1.24.0

WORKDIR /app

COPY go.mod go.sum /app

RUN go mod tidy

COPY . .

ENV CGO_ENABLED=1

EXPOSE 8080

RUN go build -o ./tracker ./cmd

CMD "./tracker"