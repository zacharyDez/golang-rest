FROM golang:1.15

WORKDIR /go
COPY src src
RUN go build src/main.go

CMD ["./main"]