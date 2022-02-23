ARG  BUILDER_IMAGE=golang:buster
ARG  DISTROLESS_IMAGE=gcr.io/distroless/base

FROM ${BUILDER_IMAGE} as builder
WORKDIR $GOPATH/src
COPY go.mod .
RUN go mod download
RUN go mod verify

COPY . .
RUN go build -o /shortener shortener/cmd/main.go
COPY config/shortener.local.yaml /shortener.local.yaml
FROM ${DISTROLESS_IMAGE}
COPY --from=builder /shortener /shortener
COPY --from=builder /shortener.local.yaml /shortener.local.yaml
ENTRYPOINT ["/shortener", "-config", "shortener.local.yaml"]
