FROM golang:1.23.3

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
