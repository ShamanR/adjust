FROM golang:1.17.8

WORKDIR /app

COPY go.mod ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/calcmd5 ./...

CMD ["/app/bin/myhttp", "http://google.com"]
