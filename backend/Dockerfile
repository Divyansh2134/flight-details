FROM golang:1.22-bullseye AS goBuild

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM debian:bullseye-slim

WORKDIR /root/

COPY --from=goBuild /app/main .

EXPOSE 8000

CMD ["./main"]
