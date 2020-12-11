# syntax=docker/dockerfile:experimental
FROM alpine:3.12.0 AS base

ARG GRPC_HEALTH_PROBE_VERSION=0.3.2

WORKDIR /workspace

RUN apk add --update curl
RUN curl -fL \
    https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/v${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 \
    -o grpc_health_probe && chmod +x grpc_health_probe


FROM golang:1.15.2-buster AS build-base

WORKDIR /workspace

ENV GOPRIVATE="github.com/doogeun"
ENV GONOSUMDB="github.com/doogeun"
COPY go.mod go.sum ./
RUN --mount=type=secret,id=gitconfig,target=/etc/gitconfig \
    go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN go build -o ./out/api ./cmd/main.go


FROM gcr.io/distroless/base-debian10:nonroot AS runtime

WORKDIR /doogeun/bin

COPY --from=base /workspace/grpc_health_probe /doogeun/bin/grpc_health_probe
COPY --from=build-base /workspace/out/api /doogeun/bin/api

CMD ["/doogeun/bin/api"]
