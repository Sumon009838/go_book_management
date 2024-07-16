FROM golang:latest as builder


LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"


WORKDIR /app


COPY go.mod go.sum ./


RUN go mod download


COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o main .



FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

ENTRYPOINT ["./main"]