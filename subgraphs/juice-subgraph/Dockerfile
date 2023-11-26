# Use an official Golang runtime as a parent image
FROM golang:1.20 as builder

WORKDIR /app

COPY ./subgraphs/juice-subgraph/ ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3009

ENTRYPOINT [ "/app/main" ]