FROM golang:1.23.3

WORKDIR /go/src/app

COPY . .

RUN go mod tidy
RUN go build -o /go/bin/app .

EXPOSE 8080

CMD ["/go/bin/app"]