FROM golang:alpine AS builder
WORKDIR /go/src/github.com/escalopa/migrate-go
COPY ./cmd ./cmd
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o /go/bin/migrate-go ./cmd/main.go

FROM alpine:latest AS production

RUN apk add --no-cache tzdata

COPY --from=builder /go/bin/migrate-go /go/bin/migrate-go
ENTRYPOINT ["/go/bin/migrate-go"]
