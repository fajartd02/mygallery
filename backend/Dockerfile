FROM golang:1.18-alpine AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY . .
RUN go mod tidy
RUN GOOS=linux go build -o main -ldflags=-X=main.version=${VERSION} .

FROM alpine
WORKDIR /go/bin
COPY --from=builder /go/src/app/.env /go/bin
COPY --from=builder /go/src/app/main /go/bin/main
ENV PATH="/go/bin:${PATH}"
CMD ["/go/bin/main"]