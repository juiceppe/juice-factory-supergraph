FROM golang:1.20 as builder

WORKDIR /app

COPY ./subgraphs/ingredients-subgraph/ ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM scratch

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 3007

ENTRYPOINT [ "/app/main" ]