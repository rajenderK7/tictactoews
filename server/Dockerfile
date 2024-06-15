# Build stage
FROM golang as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/tictactoews/main.go


# Runtime
FROM alpine:latest

WORKDIR /tictactoews

COPY --from=builder /app/bin/app .

EXPOSE 4000

CMD ["./app"]
