FROM golang:1.12.5 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY main.go main.go
COPY pkg/ pkg/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -a -o viking main.go

FROM scratch
WORKDIR /
COPY --from=builder /workspace/viking .
COPY openapi.json .

CMD ["/viking"]
