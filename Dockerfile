FROM golang:1.17.1-alpine

WORKDIR /app/cachapi

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -tags musl -o ./out/cachapi ./cmd/cachapi

CMD ["./out/cachapi"]