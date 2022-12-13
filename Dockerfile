FROM golang:1.19-alpine as builder

WORKDIR /code

ENV CGO_ENABLED 0
ENV GOPATH /go
ENV GOCACHE /go-build

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod/cache \
    go mod download && go mod verify

COPY . .

RUN --mount=type=cache,target=/go/pkg/mod/cache \
    --mount=type=cache,target=/go-build \
    go build -o bin/gateway main.go

FROM scratch
COPY --from=builder /code/bin/gateway /usr/local/bin/gateway
CMD ["/usr/local/bin/gateway"]
