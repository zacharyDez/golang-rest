FROM golang:1.15

WORKDIR /go
COPY src src

CMD ["go test /src"]