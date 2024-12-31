FROM golang:1.22.0-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o scheduler ./cmd/api/main.go

EXPOSE 8080

CMD ["./scheduler"]
