FROM golang:1.15

WORKDIR /go/src/hello

COPY ./main.go .
RUN go get -d -v ./...
RUN go install -v ./...
CMD ["hello"]