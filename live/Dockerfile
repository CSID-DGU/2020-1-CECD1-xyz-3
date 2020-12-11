FROM golang:1.15.6-buster as builder

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd ./cmd
COPY ./pkg ./pkg
RUN go build -o live ./cmd/main.go


FROM gcr.io/distroless/base as runtime

COPY --from=builder /workspace/live /
CMD ["/live"]
