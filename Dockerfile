FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -v -o certexpired cmd/certexpired/main.go

FROM bash:alpine3.17

COPY --from=builder /app/certexpired /certexpired

CMD ["/certexpired"]