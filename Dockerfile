FROM golang:1.21.0 as builder
WORKDIR /work
COPY . .
RUN go mod tidy && go build -o getpodip src/main.go

FROM debian:12-slim as prod
COPY --from=builder /work/getpodip /work/

WORKDIR /work

CMD ["./getpodip"]

