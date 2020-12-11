# syntax=docker/dockerfile:experimental
FROM golang:1.15.2-buster AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN --mount=type=secret,id=gitconfig,target=/etc/gitconfig \
    go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gateway ./cmd


FROM gcr.io/distroless/base-debian10:nonroot AS runtime

WORKDIR /

COPY --from=builder /app/bin/gateway ./
USER nonroot:nonroot

ENTRYPOINT ["/gateway"]
