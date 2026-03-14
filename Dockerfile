# syntax=docker/dockerfile:1

# ---- build stage ----
FROM golang:1.26-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/action ./cmd/action

# ---- final stage ----
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

COPY --from=builder /out/action /usr/local/bin/action

ENTRYPOINT ["/usr/local/bin/action"]
